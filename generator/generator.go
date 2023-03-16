package generator

import (
	"errors"
	"fmt"
	"io"

	"github.com/corani/silver-octo-sniffle/ast"
	"github.com/corani/silver-octo-sniffle/reporter"
	"github.com/corani/silver-octo-sniffle/token"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

var ErrGenerating = errors.New("generator error")

var zero = constant.NewInt(types.I64, 0)

func GenerateIR(out *reporter.Reporter, writer io.Writer, root ast.Node) {
	g := &Generator{
		out:           out,
		currentModule: ir.NewModule(),
		currentBlock:  nil,
		currentValue:  nil,
		currentNode:   nil,
		stdlib:        make(map[string]*ir.Func),
		strings:       make(map[string]*ir.Global),
		consts:        make(map[string]*ast.Value),
		vars:          make(map[string]*ir.Global),
	}

	defer func() {
		if r := recover(); r != nil {
			token := g.currentNode.Token()

			out.Errorf(token, "panic during code generation:")
			out.Errorf(token, "%v", r)
		}
	}()

	module := g.Generate(root)

	fmt.Fprintln(writer, module)
}

type Generator struct {
	out           *reporter.Reporter
	currentModule *ir.Module
	currentFunc   *ir.Func
	currentBlock  *ir.Block
	currentValue  value.Value
	currentNode   ast.Node
	stdlib        map[string]*ir.Func
	strings       map[string]*ir.Global
	consts        map[string]*ast.Value
	vars          map[string]*ir.Global
}

var (
	_ ast.AstVisitor     = (*Generator)(nil)
	_ ast.BuiltinVisitor = (*Generator)(nil)
)

func (g *Generator) Generate(n ast.Node) *ir.Module {
	g.currentNode = n

	g.generateStdlib()

	n.Visit(g)

	return g.currentModule
}

func (g *Generator) VisitModule(n *ast.Module) {
	leaveNode := g.enterNode(n)
	defer leaveNode()

	for _, decl := range n.Decls() {
		decl.Visit(g)
	}

	// TODO(daniel): is it correct that only the "main" module has statements, so that we can
	// generate an entry-point from them?
	if n.Block() != nil {
		g.generateMain(n.Block())
	}
}

func (g *Generator) generateOberonMain(n ast.Stmt) *ir.Func {
	leaveNode := g.enterNode(n)
	defer leaveNode()

	g.currentFunc = g.currentModule.NewFunc("oberonMain", types.Void)
	g.currentBlock = g.currentFunc.NewBlock("entry")

	n.Visit(g)

	g.currentBlock.NewRet(nil)

	return g.currentFunc
}

func (g *Generator) generateMainPreamble(main *ir.Func) {
	g.vars["__argc"] = g.currentModule.NewGlobalDef("__argc",
		constant.NewInt(types.I64, 0))
	g.vars["__argv"] = g.currentModule.NewGlobalDef("__argv",
		constant.NewIntToPtr(constant.NewInt(types.I8, 0), types.NewPointer(types.I8Ptr)))

	// Capture command-line arguments
	g.currentBlock.NewStore(main.Params[0], g.vars["__argc"])
	g.currentBlock.NewStore(main.Params[1], g.vars["__argv"])

	// TODO(daniel): capture environment variables?
}

func (g *Generator) generateMain(n ast.Stmt) {
	oberonMain := g.generateOberonMain(n)

	g.currentFunc = g.currentModule.NewFunc("main", types.I64,
		ir.NewParam("argc", types.I64),
		ir.NewParam("argv", types.NewPointer(types.I8Ptr)))
	g.currentBlock = g.currentFunc.NewBlock("entry")

	g.generateMainPreamble(g.currentFunc)

	g.currentBlock.NewCall(oberonMain)
	g.currentBlock.NewRet(zero)
}

// ----- statements --------------------------------------------------------------------------------

func (g *Generator) VisitInvalidStmt(n *ast.InvalidStmt) {
	leaveNode := g.enterNode(n)
	defer leaveNode()
}

func (g *Generator) VisitStmtSequence(n *ast.StmtSequence) {
	leaveNode := g.enterNode(n)
	defer leaveNode()

	for _, stmt := range n.Stmts() {
		stmt.Visit(g)
	}
}

func (g *Generator) VisitAssignStmt(n *ast.AssignStmt) {
	leaveNode := g.enterNode(n)
	defer leaveNode()

	des := g.visitAndReturnValue(n.Designator())
	expr := g.visitAndReturnValue(n.Expr())

	g.currentBlock.NewStore(expr, des)
}

func (g *Generator) VisitIfStmt(n *ast.IfStmt) {
	leaveNode := g.enterNode(n)
	defer leaveNode()

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
	leaveNode := g.enterNode(n)
	defer leaveNode()

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
	leaveNode := g.enterNode(n)
	defer leaveNode()

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
	leaveNode := g.enterNode(n)
	defer leaveNode()

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

func (g *Generator) VisitReturnStmt(n *ast.ReturnStmt) {
	leaveNode := g.enterNode(n)
	defer leaveNode()

	n.Expr().Visit(g)
}

func (g *Generator) VisitExprStmt(n *ast.ExprStmt) {
	leaveNode := g.enterNode(n)
	defer leaveNode()

	// NOTE(daniel): ignore the result.
	n.Expr().Visit(g)
}

// ----- expressions -------------------------------------------------------------------------------

func (g *Generator) VisitInvalidExpr(n *ast.InvalidExpr) {
	leaveNode := g.enterNode(n)
	defer leaveNode()
}

func (g *Generator) VisitCallExpr(n *ast.CallExpr) {
	leaveNode := g.enterNode(n)
	defer leaveNode()

	name := n.Token().Text

	if fn := n.Builtin(); fn != nil {
		fn.Visit(g, n.Args())
	} else {
		g.out.Errorf(n.Token(), "don't know how to call %q", name)
	}
}

func (g *Generator) VisitBinaryExpr(n *ast.BinaryExpr) {
	leaveNode := g.enterNode(n)
	defer leaveNode()

	left := g.visitAndReturnValue(n.Lhs())
	right := g.visitAndReturnValue(n.Rhs())

	if !left.Type().Equal(right.Type()) {
		g.out.Errorf(n.Token(), "types are different in %s",
			n.Token().Type)
	}

	// TODO(daniel): check if set operations are correct!
	switch n.Token().Type {
	case token.TokenMinus:
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
	case token.TokenPlus:
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
	case token.TokenAsterisk:
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
	case token.TokenSlash:
		if n.Lhs().Type() == ast.TypeSet {
			// SET symmetric difference
			g.currentValue = g.currentBlock.NewXor(left, right)
		} else {
			// REAL division
			g.currentValue = g.currentBlock.NewFDiv(left, right)
		}
	case token.TokenDIV:
		// INTEGER division
		g.currentValue = g.currentBlock.NewSDiv(left, right)
	case token.TokenMOD:
		// INTEGER modulus
		g.currentValue = g.currentBlock.NewSRem(left, right)
	case token.TokenAmpersand:
		// logical AND
		g.currentValue = g.currentBlock.NewAnd(left, right)
	case token.TokenOR:
		// logical OR
		g.currentValue = g.currentBlock.NewOr(left, right)
	case token.TokenEQ:
		if _, ok := left.Type().(*types.IntType); ok {
			g.currentValue = g.currentBlock.NewICmp(enum.IPredEQ, left, right)
		} else {
			// TODO(daniel): do we need "ordered" or "unordered" float compares?
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredUEQ, left, right)
		}
	case token.TokenNE:
		if _, ok := left.Type().(*types.IntType); ok {
			g.currentValue = g.currentBlock.NewICmp(enum.IPredNE, left, right)
		} else {
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredUNE, left, right)
		}
	case token.TokenLT:
		if _, ok := left.Type().(*types.IntType); ok {
			// TODO(daniel): "signed" or "unsigned" integer compares?
			g.currentValue = g.currentBlock.NewICmp(enum.IPredSLT, left, right)
		} else {
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredULT, left, right)
		}
	case token.TokenLE:
		if _, ok := left.Type().(*types.IntType); ok {
			g.currentValue = g.currentBlock.NewICmp(enum.IPredSLE, left, right)
		} else {
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredULE, left, right)
		}
	case token.TokenGE:
		if _, ok := left.Type().(*types.IntType); ok {
			g.currentValue = g.currentBlock.NewICmp(enum.IPredSGE, left, right)
		} else {
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredUGE, left, right)
		}
	case token.TokenGT:
		if _, ok := left.Type().(*types.IntType); ok {
			g.currentValue = g.currentBlock.NewICmp(enum.IPredSGT, left, right)
		} else {
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredUGT, left, right)
		}
	case token.TokenIN:
		g.currentValue = g.currentBlock.NewShl(constant.NewInt(types.I64, 1), left)
		g.currentValue = g.currentBlock.NewAnd(right, g.currentValue)
		g.currentValue = g.currentBlock.NewICmp(enum.IPredNE, g.currentValue, zero)
	default:
		g.out.Errorf(n.Token(), "unsupported token type: %+v", n.Token().Type)
	}
}

func (g *Generator) VisitDesignatorExpr(n *ast.DesignatorExpr) {
	leaveNode := g.enterNode(n)
	defer leaveNode()

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
		if n.IsAssignment() {
			// NOTE(daniel):
			// for `x^ := 33` we want to end up with the following:
			// ```llvm
			// %0 = load i64*, i64** @x  	; <==
			// store i64 33, i64* %0
			// ```
			//
			// for `x := 33` we want to end up with the following:
			// ```llvm
			// store i64 33, i64* @x
			// ```

			// TODO(daniel): process selectors
			if len(n.Selectors()) > 0 {
				g.currentValue = g.currentBlock.NewLoad(v.ContentType, v)
			} else {
				g.currentValue = v
			}
		} else {
			// NOTE(daniel):
			// for `print(x^)` we want to end up with the following:
			// ```llvm
			// %0 = load i64*, i64** @x
			// %1 = load i64, i64* %0
			// %2 = call i64 (i8*, ...) @printf(..., i64 %1)
			// ```
			//
			// for `print(x)` we want to end up with the following:
			// ```llvm
			// %0 = load i64, i64* @x		; <==
			// %1 = call i64 (i8*, ...) @printf(..., i64 %0)
			// ```
			g.currentValue = g.currentBlock.NewLoad(v.ContentType, v)

			// TODO(daniel): process selectors
			if len(n.Selectors()) > 0 {
				if pt, ok := g.currentValue.Type().(*types.PointerType); ok {
					g.currentValue = g.currentBlock.NewLoad(pt.ElemType, g.currentValue)
				}
			}
		}
	}
}

func (g *Generator) VisitNotExpr(n *ast.NotExpr) {
	leaveNode := g.enterNode(n)
	defer leaveNode()

	val := g.visitAndReturnValue(n.Expr())

	// TODO(daniel): is this how you do it?
	g.currentValue = g.currentBlock.NewICmp(enum.IPredEQ, val, constant.NewInt(types.I1, 0))
}

// ----- literals ----------------------------------------------------------------------------------

func (g *Generator) VisitNumberLit(n *ast.NumberLit) {
	leaveNode := g.enterNode(n)
	defer leaveNode()

	switch n.Type() {
	case ast.TypeInt64:
		g.currentValue = constant.NewInt(types.I64, int64(n.Token().Int))
	case ast.TypeFloat64:
		g.currentValue = constant.NewFloat(types.Double, n.Token().Real)
	}
}

func (g *Generator) VisitStringLit(n *ast.StringLit) {
	leaveNode := g.enterNode(n)
	defer leaveNode()

	str := g.internString(n.Token().Text + "\000")

	g.currentValue = g.currentBlock.NewGetElementPtr(str.ContentType, str, zero, zero)
}

func (g *Generator) VisitCharLit(n *ast.CharLit) {
	leaveNode := g.enterNode(n)
	defer leaveNode()

	g.currentValue = constant.NewInt(types.I8, int64(n.ConstValue().Int()))
}

func (g *Generator) VisitBooleanLit(n *ast.BooleanLit) {
	leaveNode := g.enterNode(n)
	defer leaveNode()

	if n.Token().Bool {
		g.currentValue = constant.True
	} else {
		g.currentValue = constant.False
	}
}

func (g *Generator) VisitSetLit(n *ast.SetLit) {
	leaveNode := g.enterNode(n)
	defer leaveNode()

	g.currentValue = constant.NewInt(types.I64, int64(n.ConstValue().Int()))
}

// ----- declarations ------------------------------------------------------------------------------

func (g *Generator) VisitConstDecl(decl *ast.ConstDecl) {
	leaveNode := g.enterNode(decl)
	defer leaveNode()

	g.consts[decl.Token().Text] = decl.Value()
}

func (g *Generator) VisitTypeBaseDecl(decl *ast.TypeBaseDecl) {
	leaveNode := g.enterNode(decl)
	defer leaveNode()
}

func (g *Generator) VisitTypePointerDecl(decl *ast.TypePointerDecl) {
	leaveNode := g.enterNode(decl)
	defer leaveNode()
}

func (g *Generator) VisitVarDecl(decl *ast.VarDecl) {
	leaveNode := g.enterNode(decl)
	defer leaveNode()

	switch decl.Type() {
	case ast.TypeInt64, ast.TypeSet:
		g.vars[decl.Token().Text] = g.currentModule.NewGlobalDef("", constant.NewInt(types.I64, 0))
	case ast.TypeFloat64:
		g.vars[decl.Token().Text] = g.currentModule.NewGlobalDef("", constant.NewFloat(types.Double, 0))
	case ast.TypeBoolean:
		g.vars[decl.Token().Text] = g.currentModule.NewGlobalDef("", constant.NewInt(types.I1, 0))
	case ast.TypeChar:
		g.vars[decl.Token().Text] = g.currentModule.NewGlobalDef("", constant.NewInt(types.I8, 0))
	case ast.TypePointer:
		g.vars[decl.Token().Text] = g.currentModule.NewGlobalDef("",
			constant.NewIntToPtr(constant.NewInt(types.I64, 0), types.I64Ptr))
	default:
		g.out.Errorf(decl.Token(), "don't know how to declare global variable of type %s", decl.Type())
	}
}

func (g *Generator) VisitProcDecl(decl *ast.ProcDecl) {
	leaveNode := g.enterNode(decl)
	defer leaveNode()
}

func (g *Generator) VisitTypeRef(ref *ast.TypeRef) {
	leaveNode := g.enterNode(ref)
	defer leaveNode()
}

func (g *Generator) visitAndReturnValue(n ast.Node) value.Value {
	leaveNode := g.enterNode(n)
	defer leaveNode()

	n.Visit(g)

	return g.currentValue
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
	case ast.TypeSet:
		g.visitAndReturnValue(args[0])
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
	// TODO(daniel): hardcoded malloc size
	g.currentValue = g.currentBlock.NewCall(g.stdlib["malloc"], constant.NewInt(types.I64, 8))
	g.currentValue = g.currentBlock.NewBitCast(g.currentValue, types.I64Ptr)
	g.currentBlock.NewStore(g.currentValue, g.vars[args[0].Token().Text])
}

func (g *Generator) VisitBuiltinDELETE(f *ast.BuiltinDELETE, args []ast.Expr) {
	if v, ok := g.vars[args[0].Token().Text]; ok {
		g.currentValue = g.currentBlock.NewLoad(v.ContentType, v)
		g.currentValue = g.currentBlock.NewBitCast(g.currentValue, types.I8Ptr)
		g.currentValue = g.currentBlock.NewCall(g.stdlib["free"], g.currentValue)
	}
}

func (g *Generator) VisitBuiltinASSERT(f *ast.BuiltinASSERT, args []ast.Expr) {
}

func (g *Generator) VisitCPrint(f *ast.CPrint, args []ast.Expr) {
	arg := args[0]

	// NOTE(daniel): we could probably do something smarter here, but for now this works.
	switch arg.Type() {
	case ast.TypeInt64:
		g.printInteger(arg)
		g.println()
	case ast.TypeFloat64:
		g.printReal(arg)
		g.println()
	case ast.TypeString:
		g.printString(arg)
		g.println()
	case ast.TypeChar:
		g.printChar(arg)
		g.println()
	default:
		g.out.Errorf(arg.Token(), "don't know how to print type: %s",
			arg.Type())
	}
}

func (g *Generator) VisitTextsWriteInt(f *ast.TextsWriteInt, args []ast.Expr) {
	g.printInteger(args[0])
}

func (g *Generator) VisitTextsWriteReal(f *ast.TextsWriteReal, args []ast.Expr) {
	g.printReal(args[0])
}

func (g *Generator) VisitTextsWriteString(f *ast.TextsWriteString, args []ast.Expr) {
	g.printString(args[0])
}

func (g *Generator) VisitTextsWriteChar(f *ast.TextsWriteChar, args []ast.Expr) {
	g.printChar(args[0])
}

func (g *Generator) VisitTextsWriteLn(f *ast.TextsWriteLn, args []ast.Expr) {
	g.println()
}

func (g *Generator) printInteger(arg ast.Expr) {
	leaveNode := g.enterNode(arg)
	defer leaveNode()

	number := g.visitAndReturnValue(arg)

	format := g.internString("%d\000")
	formatptr := g.currentBlock.NewGetElementPtr(format.ContentType, format, zero, zero)

	g.currentValue = g.currentBlock.NewCall(g.stdlib["printf"], formatptr, number)
}

func (g *Generator) printReal(arg ast.Expr) {
	leaveNode := g.enterNode(arg)
	defer leaveNode()

	number := g.visitAndReturnValue(arg)

	format := g.internString("%f\000")
	formatptr := g.currentBlock.NewGetElementPtr(format.ContentType, format, zero, zero)

	g.currentValue = g.currentBlock.NewCall(g.stdlib["printf"], formatptr, number)
}

func (g *Generator) printString(arg ast.Expr) {
	leaveNode := g.enterNode(arg)
	defer leaveNode()

	str := g.visitAndReturnValue(arg)

	g.currentValue = g.currentBlock.NewCall(g.stdlib["printf"], str)
}

func (g *Generator) printChar(arg ast.Expr) {
	leaveNode := g.enterNode(arg)
	defer leaveNode()

	v := g.visitAndReturnValue(arg)

	format := g.internString("%c\000")
	formatptr := g.currentBlock.NewGetElementPtr(format.ContentType, format, zero, zero)

	g.currentValue = g.currentBlock.NewCall(g.stdlib["printf"], formatptr, v)
}

func (g *Generator) println() {
	format := g.internString("\000")
	formatptr := g.currentBlock.NewGetElementPtr(format.ContentType, format, zero, zero)

	g.currentValue = g.currentBlock.NewCall(g.stdlib["puts"], formatptr)
}

// ----- helper functions --------------------------------------------------------------------------

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

	malloc := g.currentModule.NewFunc("malloc", types.I8Ptr,
		ir.NewParam("size", types.I64))
	g.stdlib["malloc"] = malloc

	free := g.currentModule.NewFunc("free", types.I8Ptr,
		ir.NewParam("ptr", types.I8Ptr))
	g.stdlib["free"] = free
}

func (g *Generator) internString(s string) *ir.Global {
	if v, ok := g.strings[s]; ok {
		return v
	}

	def := g.currentModule.NewGlobalDef("", constant.NewCharArrayFromString(s))
	g.strings[s] = def

	return def
}

func (g *Generator) enterNode(n ast.Node) func() {
	old := g.currentNode

	g.currentNode = n

	return func() {
		// NOTE(daniel): deferred functions are called while unwinding a panic. We don't want
		// to reset the current node as we'd end up reporting the panic on the root node.
		if r := recover(); r != nil {
			panic(r)
		}

		g.currentNode = old
	}
}
