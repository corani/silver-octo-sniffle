package main

import (
	"fmt"
	"io"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

var zero = constant.NewInt(types.I64, 0)

func generateIR(writer io.Writer, root Node) error {
	g := &generator{
		currentModule: ir.NewModule(),
		currentBlock:  nil,
		currentValue:  nil,
		stdlib:        make(map[string]*ir.Func),
		strings:       make(map[string]*ir.Global),
		consts:        make(map[string]Value),
		vars:          make(map[string]*ir.Global),
	}

	module, err := g.Generate(root)
	if err != nil {
		return err
	}

	fmt.Fprintln(writer, module)

	return nil
}

type generator struct {
	currentModule *ir.Module
	currentFunc   *ir.Func
	currentBlock  *ir.Block
	currentValue  value.Value
	stdlib        map[string]*ir.Func
	strings       map[string]*ir.Global
	consts        map[string]Value
	vars          map[string]*ir.Global
	errors        []error
}

var _ Visitor = (*generator)(nil)

func (g *generator) Generate(root Node) (*ir.Module, error) {
	g.generateStdlib()

	root.Visit(g)

	return g.currentModule, g.Error()
}

// TODO(daniel): improve error reporting. Maybe with a callback?
func (g *generator) Error() error {
	if len(g.errors) == 0 {
		return nil
	}

	txt := "generate errors"

	for _, v := range g.errors {
		txt = txt + ": " + v.Error()
	}

	return fmt.Errorf(txt)
}

func (g *generator) VisitModule(n *Module) {
	for _, decl := range n.consts {
		g.consts[decl.token.Text] = decl.value
	}

	for _, decl := range n.vars {
		switch decl.typ {
		case TypeInt64:
			g.vars[decl.token.Text] = g.currentModule.NewGlobalDef("", constant.NewInt(types.I64, 0))
		case TypeFloat64:
			g.vars[decl.token.Text] = g.currentModule.NewGlobalDef("", constant.NewFloat(types.Double, 0))
		case TypeBoolean:
			g.vars[decl.token.Text] = g.currentModule.NewGlobalDef("", constant.NewInt(types.I1, 0))
		}
	}

	// TODO(daniel): is it correct that only the "main" module has statements, so that we can
	// generate an entry-point from them?
	if n.stmts != nil {
		g.currentFunc = g.currentModule.NewFunc("main", types.I64)
		g.currentBlock = g.currentFunc.NewBlock("entry")

		n.stmts.Visit(g)

		g.currentBlock.NewRet(zero)
	}
}

func (g *generator) VisitStmtSequence(n *StmtSequence) {
	for _, stmt := range n.stmts {
		stmt.Visit(g)
	}
}

func (g *generator) VisitAssignStmt(n *AssignStmt) {
	expr := g.visitAndReturnValue(n.expr)

	g.currentBlock.NewStore(expr, g.vars[n.token.Text])
}

func (g *generator) VisitIfStmt(n *IfStmt) {
	condition := g.visitAndReturnValue(n.expr)

	trueBlk := g.currentFunc.NewBlock("")
	falseBlk := g.currentFunc.NewBlock("")
	endBlk := g.currentFunc.NewBlock("")

	g.currentBlock.NewCondBr(condition, trueBlk, falseBlk)

	g.currentBlock = trueBlk
	n.trueBlock.Visit(g)
	g.currentBlock.NewBr(endBlk)

	g.currentBlock = falseBlk
	n.falseBlock.Visit(g)
	g.currentBlock.NewBr(endBlk)

	g.currentBlock = endBlk
}

func (g *generator) VisitRepeatStmt(n *RepeatStmt) {
	startBlk := g.currentFunc.NewBlock("")
	doneBlk := g.currentFunc.NewBlock("")

	g.currentBlock.NewBr(startBlk)
	g.currentBlock = startBlk

	n.cond.stmt.Visit(g)

	cond := g.visitAndReturnValue(n.cond.expr)

	g.currentBlock.NewCondBr(cond, doneBlk, startBlk)
	g.currentBlock = doneBlk
}

func (g *generator) VisitWhileStmt(n *WhileStmt) {
	var startBlk, condBlk, bodyBlk, doneBlk *ir.Block

	condBlk = g.currentFunc.NewBlock("")
	startBlk = condBlk

	g.currentBlock.NewBr(condBlk)

	for _, cond := range n.conds {
		bodyBlk = g.currentFunc.NewBlock("")
		doneBlk = g.currentFunc.NewBlock("")

		g.currentBlock = condBlk
		value := g.visitAndReturnValue(cond.expr)
		g.currentBlock.NewCondBr(value, bodyBlk, doneBlk)

		g.currentBlock = bodyBlk
		cond.stmt.Visit(g)
		g.currentBlock.NewBr(startBlk)

		condBlk = doneBlk
	}

	g.currentBlock = doneBlk
}

func (g *generator) VisitForStmt(n *ForStmt) {
	bodyBlk := g.currentFunc.NewBlock("")
	doneBlk := g.currentFunc.NewBlock("")

	fromv := g.visitAndReturnValue(n.from)
	incrv := g.visitAndReturnValue(n.by)

	g.currentBlock.NewStore(fromv, g.vars[n.iter.Text])
	g.currentBlock.NewBr(bodyBlk)

	g.currentBlock = bodyBlk
	n.stmt.Visit(g)

	oldv := g.currentBlock.NewLoad(g.vars[n.iter.Text].ContentType, g.vars[n.iter.Text])
	newv := g.currentBlock.NewAdd(oldv, incrv)
	g.currentBlock.NewStore(newv, g.vars[n.iter.Text])
	tov := g.visitAndReturnValue(n.to)

	var cond value.Value

	if n.by.ConstValue().Int() > 0 {
		cond = g.currentBlock.NewICmp(enum.IPredSLE, newv, tov)
	} else {
		cond = g.currentBlock.NewICmp(enum.IPredSGE, newv, tov)
	}

	g.currentBlock.NewCondBr(cond, bodyBlk, doneBlk)

	g.currentBlock = doneBlk
}

func (g *generator) VisitExprStmt(n *ExprStmt) {
	// NOTE(daniel): ignore the result.
	n.expr.Visit(g)
}

func (g *generator) VisitCallExpr(n *CallExpr) {
	switch n.Token().Text {
	case "print":
		g.callPrint(n.args)
	case "INC":
		g.callIncDec(n.args[0], 1)
	case "DEC":
		g.callIncDec(n.args[0], -1)
	case "FLT":
		g.callFLT(n.args[0])
	case "FLOOR":
		g.callFLOOR(n.args[0])
	default:
		g.errors = append(g.errors, fmt.Errorf("don't know how to call %q",
			n.Token().Text))
	}
}

func (g *generator) VisitBinaryExpr(n *BinaryExpr) {
	left := g.visitAndReturnValue(n.args[0])
	right := g.visitAndReturnValue(n.args[1])

	if !left.Type().Equal(right.Type()) {
		g.errors = append(g.errors, fmt.Errorf("types are different in %s",
			n.token.Type))
	}

	switch n.token.Type {
	case TokenMinus:
		if left.Type().Equal(types.I64) {
			// INTEGER subtraction
			g.currentValue = g.currentBlock.NewSub(left, right)
		} else {
			// REAL subtraction
			g.currentValue = g.currentBlock.NewFSub(left, right)
		}
	case TokenPlus:
		if left.Type().Equal(types.I64) {
			// INTEGER addition
			g.currentValue = g.currentBlock.NewAdd(left, right)
		} else {
			// REAL addition
			g.currentValue = g.currentBlock.NewFAdd(left, right)
		}
	case TokenAsterisk:
		if left.Type().Equal(types.I64) {
			// INTEGER multiplication
			g.currentValue = g.currentBlock.NewMul(left, right)
		} else {
			// REAL multiplication
			g.currentValue = g.currentBlock.NewFMul(left, right)
		}
	case TokenSlash:
		// REAL division
		g.currentValue = g.currentBlock.NewFDiv(left, right)
	case TokenDIV:
		// INTEGER division
		g.currentValue = g.currentBlock.NewSDiv(left, right)
	case TokenMOD:
		// INTEGER modulus
		g.currentValue = g.currentBlock.NewSRem(left, right)
	case TokenAmpersand:
		// logical AND
		g.currentValue = g.currentBlock.NewAnd(left, right)
	case TokenOR:
		// logical OR
		g.currentValue = g.currentBlock.NewOr(left, right)
	case TokenEQ:
		if left.Type().Equal(types.I64) {
			g.currentValue = g.currentBlock.NewICmp(enum.IPredEQ, left, right)
		} else {
			// TODO(daniel): do we need "ordered" or "unordered" float compares?
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredUEQ, left, right)
		}
	case TokenNE:
		if left.Type().Equal(types.I64) {
			g.currentValue = g.currentBlock.NewICmp(enum.IPredNE, left, right)
		} else {
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredUNE, left, right)
		}
	case TokenLT:
		if left.Type().Equal(types.I64) {
			// TODO(daniel): "signed" or "unsigned" integer compares?
			g.currentValue = g.currentBlock.NewICmp(enum.IPredSLT, left, right)
		} else {
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredULT, left, right)
		}
	case TokenLE:
		if left.Type().Equal(types.I64) {
			g.currentValue = g.currentBlock.NewICmp(enum.IPredSLE, left, right)
		} else {
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredULE, left, right)
		}
	case TokenGE:
		if left.Type().Equal(types.I64) {
			g.currentValue = g.currentBlock.NewICmp(enum.IPredSGE, left, right)
		} else {
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredUGE, left, right)
		}
	case TokenGT:
		if left.Type().Equal(types.I64) {
			g.currentValue = g.currentBlock.NewICmp(enum.IPredSGT, left, right)
		} else {
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredUGT, left, right)
		}
	default:
		g.errors = append(g.errors, fmt.Errorf("unsupported token type: %+v", n.token))
	}
}

func (g *generator) VisitDesignatorExpr(n *DesignatorExpr) {
	if v, ok := g.consts[n.token.Text]; ok {
		switch v.Type() {
		case TypeInt64:
			g.currentValue = constant.NewInt(types.I64, int64(v.Int()))
		case TypeFloat64:
			g.currentValue = constant.NewFloat(types.Double, v.Real())
		case TypeBoolean:
			g.currentValue = constant.NewBool(v.Bool)
		}
	} else if v, ok := g.vars[n.token.Text]; ok {
		g.currentValue = g.currentBlock.NewLoad(v.ContentType, v)
	}
}

func (g *generator) VisitNumberExpr(n *NumberExpr) {
	switch n.Type() {
	case TypeInt64:
		g.currentValue = constant.NewInt(types.I64, int64(n.token.Int))
	case TypeFloat64:
		g.currentValue = constant.NewFloat(types.Double, n.token.Real)
	}
}

func (g *generator) VisitStringExpr(n *StringExpr) {
	str := g.internString(n.token.Text + "\000")

	g.currentValue = g.currentBlock.NewGetElementPtr(str.ContentType, str, zero, zero)
}

func (g *generator) VisitBooleanExpr(n *BooleanExpr) {
	if n.token.Bool {
		g.currentValue = constant.True
	} else {
		g.currentValue = constant.False
	}
}

func (g *generator) VisitNotExpr(n *NotExpr) {
	val := g.visitAndReturnValue(n.expr)

	// TODO(daniel): is this how you do it?
	g.currentValue = g.currentBlock.NewICmp(enum.IPredEQ, val, constant.NewInt(types.I1, 0))
}

func (g *generator) visitAndReturnValue(n Node) value.Value {
	n.Visit(g)

	return g.currentValue
}

func (g *generator) callIncDec(arg Expr, offset int64) {
	v := g.visitAndReturnValue(arg)
	g.currentValue = g.currentBlock.NewAdd(v, constant.NewInt(types.I64, offset))
	g.currentBlock.NewStore(g.currentValue, g.vars[arg.Token().Text])
}

func (g *generator) callFLT(arg Expr) {
	number := g.visitAndReturnValue(arg)
	g.currentValue = g.currentBlock.NewSIToFP(number, types.Double)
}

func (g *generator) callFLOOR(arg Expr) {
	number := g.visitAndReturnValue(arg)
	g.currentValue = g.currentBlock.NewFPToSI(number, types.I64)
}

func (g *generator) callPrint(args []Expr) {
	// TODO(daniel): make print a (polymorphic?) function in the "stdlib" and make an actual
	// function call here.
	for _, arg := range args {
		// NOTE(daniel): we could probably do something smarter here, but for now this works.
		switch arg.Type() {
		case TypeInt64:
			g.printInteger(arg)
		case TypeFloat64:
			g.printReal(arg)
		case TypeString:
			g.printString(arg)
		case TypeBoolean:
			g.printBoolean(arg)
		default:
			g.errors = append(g.errors, fmt.Errorf("don't know how to print type: %s", arg.Type()))
		}
	}
}

func (g *generator) printInteger(arg Expr) {
	number := g.visitAndReturnValue(arg)

	format := g.internString("%d\n\000")
	formatptr := g.currentBlock.NewGetElementPtr(format.ContentType, format, zero, zero)

	g.currentValue = g.currentBlock.NewCall(g.stdlib["printf"], formatptr, number)
}

func (g *generator) printReal(arg Expr) {
	number := g.visitAndReturnValue(arg)

	format := g.internString("%f\n\000")
	formatptr := g.currentBlock.NewGetElementPtr(format.ContentType, format, zero, zero)

	g.currentValue = g.currentBlock.NewCall(g.stdlib["printf"], formatptr, number)
}

func (g *generator) printString(arg Expr) {
	str := g.visitAndReturnValue(arg)

	g.currentValue = g.currentBlock.NewCall(g.stdlib["puts"], str)
}

func (g *generator) printBoolean(arg Expr) {
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

func (g *generator) generateStdlib() {
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

func (g *generator) internString(s string) *ir.Global {
	if v, ok := g.strings[s]; ok {
		return v
	}

	def := g.currentModule.NewGlobalDef("", constant.NewCharArrayFromString(s))
	g.strings[s] = def

	return def
}
