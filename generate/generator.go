package generate

import (
	"fmt"
	"io"

	"github.com/corani/silver-octo-sniffle/ast"
	"github.com/corani/silver-octo-sniffle/lex"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

var zero = constant.NewInt(types.I64, 0)

func GenerateIR(writer io.Writer, root ast.Node) error {
	g := &Generator{
		currentModule: ir.NewModule(),
		currentBlock:  nil,
		currentValue:  nil,
		stdlib:        make(map[string]*ir.Func),
		strings:       make(map[string]*ir.Global),
		consts:        make(map[string]*ast.Value),
		vars:          make(map[string]*ir.Global),
	}

	module, err := g.Generate(root)
	if err != nil {
		return err
	}

	fmt.Fprintln(writer, module)

	return nil
}

type Generator struct {
	currentModule *ir.Module
	currentFunc   *ir.Func
	currentBlock  *ir.Block
	currentValue  value.Value
	stdlib        map[string]*ir.Func
	strings       map[string]*ir.Global
	consts        map[string]*ast.Value
	vars          map[string]*ir.Global
	errors        []error
}

var (
	_ ast.AstVisitor     = (*Generator)(nil)
	_ ast.BuiltinVisitor = (*Generator)(nil)
)

func (g *Generator) Generate(root ast.Node) (*ir.Module, error) {
	g.generateStdlib()

	root.Visit(g)

	return g.currentModule, g.Error()
}

// TODO(daniel): improve error reporting. Maybe with a callback?
func (g *Generator) Error() error {
	if len(g.errors) == 0 {
		return nil
	}

	txt := "generate errors"

	for _, v := range g.errors {
		txt = txt + ": " + v.Error()
	}

	return fmt.Errorf(txt)
}

func (g *Generator) VisitModule(n *ast.Module) {
	for _, decl := range n.Consts() {
		g.consts[decl.Token().Text] = decl.Value()
	}

	for _, decl := range n.Vars() {
		switch decl.Type() {
		case ast.TypeInt64, ast.TypeSet:
			g.vars[decl.Token().Text] = g.currentModule.NewGlobalDef("", constant.NewInt(types.I64, 0))
		case ast.TypeFloat64:
			g.vars[decl.Token().Text] = g.currentModule.NewGlobalDef("", constant.NewFloat(types.Double, 0))
		case ast.TypeBoolean:
			g.vars[decl.Token().Text] = g.currentModule.NewGlobalDef("", constant.NewInt(types.I1, 0))
		case ast.TypeChar:
			g.vars[decl.Token().Text] = g.currentModule.NewGlobalDef("", constant.NewInt(types.I8, 0))
		default:
			panic(fmt.Sprintf("don't know how to declare global variable of type %s", decl.Type()))
		}
	}

	// TODO(daniel): is it correct that only the "main" module has statements, so that we can
	// generate an entry-point from them?
	if n.Block() != nil {
		g.currentFunc = g.currentModule.NewFunc("main", types.I64)
		g.currentBlock = g.currentFunc.NewBlock("entry")

		n.Block().Visit(g)

		g.currentBlock.NewRet(zero)
	}
}

func (g *Generator) VisitStmtSequence(n *ast.StmtSequence) {
	for _, stmt := range n.Stmts() {
		stmt.Visit(g)
	}
}

func (g *Generator) VisitAssignStmt(n *ast.AssignStmt) {
	expr := g.visitAndReturnValue(n.Expr())

	g.currentBlock.NewStore(expr, g.vars[n.Token().Text])
}

func (g *Generator) VisitIfStmt(n *ast.IfStmt) {
	condition := g.visitAndReturnValue(n.Expr())

	trueBlk := g.currentFunc.NewBlock("")
	falseBlk := g.currentFunc.NewBlock("")
	endBlk := g.currentFunc.NewBlock("")

	g.currentBlock.NewCondBr(condition, trueBlk, falseBlk)

	g.currentBlock = trueBlk
	n.TrueBlock().Visit(g)
	g.currentBlock.NewBr(endBlk)

	g.currentBlock = falseBlk
	n.FalseBlock().Visit(g)
	g.currentBlock.NewBr(endBlk)

	g.currentBlock = endBlk
}

func (g *Generator) VisitRepeatStmt(n *ast.RepeatStmt) {
	startBlk := g.currentFunc.NewBlock("")
	doneBlk := g.currentFunc.NewBlock("")

	g.currentBlock.NewBr(startBlk)
	g.currentBlock = startBlk

	cond := n.Condition()

	cond.Stmt().Visit(g)

	v := g.visitAndReturnValue(cond.Expr())

	g.currentBlock.NewCondBr(v, doneBlk, startBlk)
	g.currentBlock = doneBlk
}

func (g *Generator) VisitWhileStmt(n *ast.WhileStmt) {
	var startBlk, condBlk, bodyBlk, doneBlk *ir.Block

	condBlk = g.currentFunc.NewBlock("")
	startBlk = condBlk

	g.currentBlock.NewBr(condBlk)

	for _, cond := range n.Conditions() {
		bodyBlk = g.currentFunc.NewBlock("")
		doneBlk = g.currentFunc.NewBlock("")

		g.currentBlock = condBlk
		value := g.visitAndReturnValue(cond.Expr())
		g.currentBlock.NewCondBr(value, bodyBlk, doneBlk)

		g.currentBlock = bodyBlk
		cond.Stmt().Visit(g)
		g.currentBlock.NewBr(startBlk)

		condBlk = doneBlk
	}

	g.currentBlock = doneBlk
}

func (g *Generator) VisitForStmt(n *ast.ForStmt) {
	bodyBlk := g.currentFunc.NewBlock("")
	doneBlk := g.currentFunc.NewBlock("")

	fromv := g.visitAndReturnValue(n.From())
	incrv := g.visitAndReturnValue(n.By())

	g.currentBlock.NewStore(fromv, g.vars[n.Iter().Text])
	g.currentBlock.NewBr(bodyBlk)

	g.currentBlock = bodyBlk
	n.Stmt().Visit(g)

	oldv := g.currentBlock.NewLoad(g.vars[n.Iter().Text].ContentType, g.vars[n.Iter().Text])
	newv := g.currentBlock.NewAdd(oldv, incrv)
	g.currentBlock.NewStore(newv, g.vars[n.Iter().Text])
	tov := g.visitAndReturnValue(n.To())

	var cond value.Value

	if n.By().ConstValue().Int() > 0 {
		cond = g.currentBlock.NewICmp(enum.IPredSLE, newv, tov)
	} else {
		cond = g.currentBlock.NewICmp(enum.IPredSGE, newv, tov)
	}

	g.currentBlock.NewCondBr(cond, bodyBlk, doneBlk)

	g.currentBlock = doneBlk
}

func (g *Generator) VisitExprStmt(n *ast.ExprStmt) {
	// NOTE(daniel): ignore the result.
	n.Expr().Visit(g)
}

func (g *Generator) VisitCallExpr(n *ast.CallExpr) {
	name := n.Token().Text

	if fn := n.Builtin(); fn != nil {
		fn.Visit(g, n.Args())
	} else {
		g.errors = append(g.errors, fmt.Errorf("don't know how to call %q", name))
	}
}

func (g *Generator) VisitBinaryExpr(n *ast.BinaryExpr) {
	left := g.visitAndReturnValue(n.Lhs())
	right := g.visitAndReturnValue(n.Rhs())

	if !left.Type().Equal(right.Type()) {
		g.errors = append(g.errors, fmt.Errorf("types are different in %s",
			n.Token().Type))
	}

	// TODO(daniel): check if set operations are correct!
	switch n.Token().Type {
	case lex.TokenMinus:
		if left.Type().Equal(types.I64) {
			if n.Lhs().Type() == ast.TypeSet {
				// SET difference
				g.currentValue = g.currentBlock.NewXor(right, constant.NewInt(types.I64, -1))
				g.currentValue = g.currentBlock.NewAnd(left, g.currentValue)
			} else {
				// INTEGER subtraction
				g.currentValue = g.currentBlock.NewSub(left, right)
			}
		} else {
			// REAL subtraction
			g.currentValue = g.currentBlock.NewFSub(left, right)
		}
	case lex.TokenPlus:
		if left.Type().Equal(types.I64) {
			if n.Lhs().Type() == ast.TypeSet {
				// SET union
				g.currentValue = g.currentBlock.NewOr(left, right)
			} else {
				// INTEGER addition
				g.currentValue = g.currentBlock.NewAdd(left, right)
			}
		} else {
			// REAL addition
			g.currentValue = g.currentBlock.NewFAdd(left, right)
		}
	case lex.TokenAsterisk:
		if left.Type().Equal(types.I64) {
			if n.Lhs().Type() == ast.TypeSet {
				// SET intersection
				g.currentValue = g.currentBlock.NewAnd(left, right)
			} else {
				// INTEGER multiplication
				g.currentValue = g.currentBlock.NewMul(left, right)
			}
		} else {
			// REAL multiplication
			g.currentValue = g.currentBlock.NewFMul(left, right)
		}
	case lex.TokenSlash:
		if n.Lhs().Type() == ast.TypeSet {
			// SET symmetric difference
			g.currentValue = g.currentBlock.NewXor(left, right)
		} else {
			// REAL division
			g.currentValue = g.currentBlock.NewFDiv(left, right)
		}
	case lex.TokenDIV:
		// INTEGER division
		g.currentValue = g.currentBlock.NewSDiv(left, right)
	case lex.TokenMOD:
		// INTEGER modulus
		g.currentValue = g.currentBlock.NewSRem(left, right)
	case lex.TokenAmpersand:
		// logical AND
		g.currentValue = g.currentBlock.NewAnd(left, right)
	case lex.TokenOR:
		// logical OR
		g.currentValue = g.currentBlock.NewOr(left, right)
	case lex.TokenEQ:
		if _, ok := left.Type().(*types.IntType); ok {
			g.currentValue = g.currentBlock.NewICmp(enum.IPredEQ, left, right)
		} else {
			// TODO(daniel): do we need "ordered" or "unordered" float compares?
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredUEQ, left, right)
		}
	case lex.TokenNE:
		if _, ok := left.Type().(*types.IntType); ok {
			g.currentValue = g.currentBlock.NewICmp(enum.IPredNE, left, right)
		} else {
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredUNE, left, right)
		}
	case lex.TokenLT:
		if _, ok := left.Type().(*types.IntType); ok {
			// TODO(daniel): "signed" or "unsigned" integer compares?
			g.currentValue = g.currentBlock.NewICmp(enum.IPredSLT, left, right)
		} else {
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredULT, left, right)
		}
	case lex.TokenLE:
		if _, ok := left.Type().(*types.IntType); ok {
			g.currentValue = g.currentBlock.NewICmp(enum.IPredSLE, left, right)
		} else {
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredULE, left, right)
		}
	case lex.TokenGE:
		if _, ok := left.Type().(*types.IntType); ok {
			g.currentValue = g.currentBlock.NewICmp(enum.IPredSGE, left, right)
		} else {
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredUGE, left, right)
		}
	case lex.TokenGT:
		if _, ok := left.Type().(*types.IntType); ok {
			g.currentValue = g.currentBlock.NewICmp(enum.IPredSGT, left, right)
		} else {
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredUGT, left, right)
		}
	case lex.TokenIN:
		g.currentValue = g.currentBlock.NewShl(constant.NewInt(types.I64, 1), left)
		g.currentValue = g.currentBlock.NewAnd(right, g.currentValue)
		g.currentValue = g.currentBlock.NewICmp(enum.IPredNE, g.currentValue, zero)
	default:
		g.errors = append(g.errors, fmt.Errorf("unsupported token type: %+v", n.Token()))
	}
}

func (g *Generator) VisitDesignatorExpr(n *ast.DesignatorExpr) {
	if v, ok := g.consts[n.Token().Text]; ok {
		switch v.Type() {
		case ast.TypeInt64, ast.TypeSet:
			g.currentValue = constant.NewInt(types.I64, int64(v.Int()))
		case ast.TypeFloat64:
			g.currentValue = constant.NewFloat(types.Double, v.Real())
		case ast.TypeBoolean:
			g.currentValue = constant.NewBool(v.Bool())
		}
	} else if v, ok := g.vars[n.Token().Text]; ok {
		g.currentValue = g.currentBlock.NewLoad(v.ContentType, v)
	}
}

func (g *Generator) VisitNumberExpr(n *ast.NumberExpr) {
	switch n.Type() {
	case ast.TypeInt64:
		g.currentValue = constant.NewInt(types.I64, int64(n.Token().Int))
	case ast.TypeFloat64:
		g.currentValue = constant.NewFloat(types.Double, n.Token().Real)
	}
}

func (g *Generator) VisitStringExpr(n *ast.StringExpr) {
	str := g.internString(n.Token().Text + "\000")

	g.currentValue = g.currentBlock.NewGetElementPtr(str.ContentType, str, zero, zero)
}

func (g *Generator) VisitCharExpr(n *ast.CharExpr) {
	g.currentValue = constant.NewInt(types.I8, int64(n.ConstValue().Int()))
}

func (g *Generator) VisitBooleanExpr(n *ast.BooleanExpr) {
	if n.Token().Bool {
		g.currentValue = constant.True
	} else {
		g.currentValue = constant.False
	}
}

func (g *Generator) VisitSetExpr(n *ast.SetExpr) {
	g.currentValue = constant.NewInt(types.I64, int64(n.ConstValue().Int()))
}

func (g *Generator) VisitNotExpr(n *ast.NotExpr) {
	val := g.visitAndReturnValue(n.Expr())

	// TODO(daniel): is this how you do it?
	g.currentValue = g.currentBlock.NewICmp(enum.IPredEQ, val, constant.NewInt(types.I1, 0))
}

func (g *Generator) visitAndReturnValue(n ast.Node) value.Value {
	n.Visit(g)

	return g.currentValue
}

func (g *Generator) generateStdlib() {
	g.stdlib["puts"] = g.currentModule.NewFunc("puts", types.I64,
		ir.NewParam("str", types.I8Ptr))
	g.stdlib["rand"] = g.currentModule.NewFunc("rand", types.I64)

	sprintf := g.currentModule.NewFunc("sprintf", types.I64,
		ir.NewParam("buf", types.I8Ptr),
		ir.NewParam("format", types.I8Ptr))
	sprintf.Sig.Variadic = true
	g.stdlib["sprintf"] = sprintf

	printf := g.currentModule.NewFunc("printf", types.I64,
		ir.NewParam("format", types.I8Ptr))
	printf.Sig.Variadic = true
	g.stdlib["printf"] = printf
}

func (g *Generator) internString(s string) *ir.Global {
	if v, ok := g.strings[s]; ok {
		return v
	}

	def := g.currentModule.NewGlobalDef("", constant.NewCharArrayFromString(s))
	g.strings[s] = def

	return def
}

// ----- builtin functions ----------------------------------------------------

func (g *Generator) VisitBuiltinABS(f *ast.BuiltinABS, args []ast.Expr) {
}

func (g *Generator) VisitBuiltinODD(f *ast.BuiltinODD, args []ast.Expr) {
}

func (g *Generator) VisitBuiltinLSL(f *ast.BuiltinLSL, args []ast.Expr) {
}

func (g *Generator) VisitBuiltinASR(f *ast.BuiltinASR, args []ast.Expr) {
}

func (g *Generator) VisitBuiltinROR(f *ast.BuiltinROR, args []ast.Expr) {
}

func (g *Generator) VisitBuiltinLEN(f *ast.BuiltinLEN, args []ast.Expr) {
}

func (g *Generator) VisitBuiltinORD(f *ast.BuiltinORD, args []ast.Expr) {
	switch args[0].Type() {
	case ast.TypeBoolean:
		v := g.visitAndReturnValue(args[0])
		g.currentValue = g.currentBlock.NewZExt(v, types.I64)
	case ast.TypeChar:
		v := g.visitAndReturnValue(args[0])
		g.currentValue = g.currentBlock.NewZExt(v, types.I64)
	}
}

func (g *Generator) VisitBuiltinCHR(f *ast.BuiltinCHR, args []ast.Expr) {
	v := g.visitAndReturnValue(args[0])
	g.currentValue = g.currentBlock.NewTrunc(v, types.I8)
}

func (g *Generator) VisitBuiltinFLOOR(f *ast.BuiltinFLOOR, args []ast.Expr) {
	v := g.visitAndReturnValue(args[0])
	g.currentValue = g.currentBlock.NewFPToSI(v, types.I64)
}

func (g *Generator) VisitBuiltinFLT(f *ast.BuiltinFLT, args []ast.Expr) {
	v := g.visitAndReturnValue(args[0])
	g.currentValue = g.currentBlock.NewSIToFP(v, types.Double)
}

func (g *Generator) VisitBuiltinINC(f *ast.BuiltinINC, args []ast.Expr) {
	offset := 1

	if len(args) == 2 {
		offset = args[1].ConstValue().Int()
	}

	if offset > 0 {
		v := g.visitAndReturnValue(args[0])

		g.currentValue = g.currentBlock.NewAdd(v, constant.NewInt(types.I64, int64(offset)))
		g.currentBlock.NewStore(g.currentValue, g.vars[args[0].Token().Text])
	}
}

func (g *Generator) VisitBuiltinDEC(f *ast.BuiltinDEC, args []ast.Expr) {
	offset := 1

	if len(args) == 2 {
		offset = args[1].ConstValue().Int()
	}

	if offset > 0 {
		v := g.visitAndReturnValue(args[0])

		g.currentValue = g.currentBlock.NewSub(v, constant.NewInt(types.I64, int64(offset)))
		g.currentBlock.NewStore(g.currentValue, g.vars[args[0].Token().Text])
	}
}

func (g *Generator) VisitBuiltinINCL(f *ast.BuiltinINCL, args []ast.Expr) {
	v := g.visitAndReturnValue(args[0])
	b := g.visitAndReturnValue(args[1])
	g.currentValue = g.currentBlock.NewShl(constant.NewInt(types.I64, 1), b)
	g.currentValue = g.currentBlock.NewOr(v, g.currentValue)
	g.currentBlock.NewStore(g.currentValue, g.vars[args[0].Token().Text])
}

func (g *Generator) VisitBuiltinEXCL(f *ast.BuiltinEXCL, args []ast.Expr) {
	v := g.visitAndReturnValue(args[0])
	b := g.visitAndReturnValue(args[1])
	g.currentValue = g.currentBlock.NewShl(constant.NewInt(types.I64, 1), b)
	g.currentValue = g.currentBlock.NewXor(g.currentValue, constant.NewInt(types.I64, -1))
	g.currentValue = g.currentBlock.NewAnd(v, g.currentValue)
	g.currentBlock.NewStore(g.currentValue, g.vars[args[0].Token().Text])
}

func (g *Generator) VisitBuiltinPACK(f *ast.BuiltinPACK, args []ast.Expr) {
}

func (g *Generator) VisitBuiltinUNPK(f *ast.BuiltinUNPK, args []ast.Expr) {
}

func (g *Generator) VisitBuiltinNEW(f *ast.BuiltinNEW, args []ast.Expr) {
}

func (g *Generator) VisitBuiltinASSERT(f *ast.BuiltinASSERT, args []ast.Expr) {
}

func (g *Generator) VisitBuiltinPrint(f *ast.BuiltinPrint, args []ast.Expr) {
	arg := args[0]

	// NOTE(daniel): we could probably do something smarter here, but for now this works.
	switch arg.Type() {
	case ast.TypeInt64:
		g.printInteger(arg)
	case ast.TypeFloat64:
		g.printReal(arg)
	case ast.TypeString:
		g.printString(arg)
	case ast.TypeBoolean:
		g.printBoolean(arg)
	case ast.TypeChar:
		g.printChar(arg)
	case ast.TypeSet:
		g.printSet(arg)
	default:
		g.errors = append(g.errors, fmt.Errorf("don't know how to print type: %s",
			arg.Type()))
	}
}

func (g *Generator) printInteger(arg ast.Expr) {
	number := g.visitAndReturnValue(arg)

	format := g.internString("%d\n\000")
	formatptr := g.currentBlock.NewGetElementPtr(format.ContentType, format, zero, zero)

	g.currentValue = g.currentBlock.NewCall(g.stdlib["printf"], formatptr, number)
}

func (g *Generator) printReal(arg ast.Expr) {
	number := g.visitAndReturnValue(arg)

	format := g.internString("%f\n\000")
	formatptr := g.currentBlock.NewGetElementPtr(format.ContentType, format, zero, zero)

	g.currentValue = g.currentBlock.NewCall(g.stdlib["printf"], formatptr, number)
}

func (g *Generator) printString(arg ast.Expr) {
	str := g.visitAndReturnValue(arg)

	g.currentValue = g.currentBlock.NewCall(g.stdlib["puts"], str)
}

func (g *Generator) printBoolean(arg ast.Expr) {
	TRUE := g.internString("TRUE\000")
	FALSE := g.internString("FALSE\000")

	boolean := g.visitAndReturnValue(arg)

	trueBlk := g.currentFunc.NewBlock("")
	falseBlk := g.currentFunc.NewBlock("")
	endBlk := g.currentFunc.NewBlock("")

	g.currentBlock.NewCondBr(boolean, trueBlk, falseBlk)

	g.currentBlock = trueBlk
	truePtr := g.currentBlock.NewGetElementPtr(TRUE.ContentType, TRUE, zero, zero)
	g.currentBlock.NewBr(endBlk)

	g.currentBlock = falseBlk
	falsePtr := g.currentBlock.NewGetElementPtr(FALSE.ContentType, FALSE, zero, zero)
	g.currentBlock.NewBr(endBlk)

	g.currentBlock = endBlk
	strptr := g.currentBlock.NewPhi(ir.NewIncoming(truePtr, trueBlk), ir.NewIncoming(falsePtr, falseBlk))
	g.currentValue = g.currentBlock.NewCall(g.stdlib["puts"], strptr)
}

func (g *Generator) printChar(arg ast.Expr) {
	v := g.visitAndReturnValue(arg)

	format := g.internString("%c\n\000")
	formatptr := g.currentBlock.NewGetElementPtr(format.ContentType, format, zero, zero)

	g.currentValue = g.currentBlock.NewCall(g.stdlib["printf"], formatptr, v)
}

func (g *Generator) printSet(arg ast.Expr) {
	// TODO(daniel): can we get a better print for sets?
	g.printInteger(arg)
}