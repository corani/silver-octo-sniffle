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

var zero = constant.NewInt(types.I32, 0)

func generateIR(writer io.Writer, root Node) {
	g := &generator{
		currentModule: ir.NewModule(),
		currentBlock:  nil,
		currentValue:  nil,
		stdlib:        make(map[string]*ir.Func),
		strings:       make(map[string]*ir.Global),
	}

	module := g.Generate(root)

	fmt.Fprintln(writer, module)
}

type generator struct {
	currentModule *ir.Module
	currentFunc   *ir.Func
	currentBlock  *ir.Block
	currentValue  value.Value
	stdlib        map[string]*ir.Func
	strings       map[string]*ir.Global
}

var _ Visitor = (*generator)(nil)

func (g *generator) Generate(root Node) *ir.Module {
	g.generateStdlib()

	root.Visit(g)

	return g.currentModule
}

func (g *generator) VisitModule(n *Module) {
	g.currentFunc = g.currentModule.NewFunc("main", types.I32)
	g.currentBlock = g.currentFunc.NewBlock("entry")

	n.stmts.Visit(g)

	g.currentBlock.NewRet(zero)
}

func (g *generator) VisitStmtSequence(n *StmtSequence) {
	for _, stmt := range n.stmts {
		stmt.Visit(g)
	}
}

func (g *generator) VisitExprStmt(n *ExprStmt) {
	// NOTE(daniel): ignore the result.
	n.expr.Visit(g)
}

func (g *generator) VisitCallExpr(n *CallExpr) {
	switch n.token.Text {
	case "print":
		g.callPrint(n.args)
	default:
		panic(fmt.Sprintf("don't know how to call %q", n.token.Text))
	}
}

func (g *generator) VisitBinaryExpr(n *BinaryExpr) {
	left := g.visitAndReturnValue(n.args[0])
	right := g.visitAndReturnValue(n.args[1])

	// TODO(daniel): is automatic conversion a thing? Alternatively we could generate AST nodes
	// during type checking.
	switch n.token.Type {
	case TokenMinus:
		if left.Type().Equal(types.I32) && right.Type().Equal(types.I32) {
			// INTEGER subtraction
			g.currentValue = g.currentBlock.NewSub(left, right)
		} else {
			// REAL subtraction
			g.currentValue = g.currentBlock.NewFSub(g.anyToReal(left), g.anyToReal(right))
		}
	case TokenPlus:
		if left.Type().Equal(types.I32) && right.Type().Equal(types.I32) {
			// INTEGER addition
			g.currentValue = g.currentBlock.NewAdd(left, right)
		} else {
			// REAL addition
			g.currentValue = g.currentBlock.NewFAdd(g.anyToReal(left), g.anyToReal(right))
		}
	case TokenAsterisk:
		if left.Type().Equal(types.I32) && right.Type().Equal(types.I32) {
			// INTEGER multiplication
			g.currentValue = g.currentBlock.NewMul(left, right)
		} else {
			// REAL multiplication
			g.currentValue = g.currentBlock.NewFMul(g.anyToReal(left), g.anyToReal(right))
		}
	case TokenSlash:
		// REAL division
		g.currentValue = g.currentBlock.NewFDiv(g.anyToReal(left), g.anyToReal(right))
	case TokenDIV:
		// INTEGER division
		g.currentValue = g.currentBlock.NewSDiv(g.anyToInteger(left), g.anyToInteger(right))
	case TokenMOD:
		// INTEGER modulus
		g.currentValue = g.currentBlock.NewSRem(g.anyToInteger(left), g.anyToInteger(right))
	case TokenAmpersand:
		// logical AND
		g.currentValue = g.currentBlock.NewAnd(left, right)
	case TokenOR:
		// logical OR
		g.currentValue = g.currentBlock.NewOr(left, right)
	case TokenEQ:
		if left.Type().Equal(types.I32) && right.Type().Equal(types.I32) {
			g.currentValue = g.currentBlock.NewICmp(enum.IPredEQ, left, right)
		} else {
			// TODO(daniel): do we need "ordered" or "unordered" float compares?
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredUEQ, g.anyToReal(left), g.anyToReal(right))
		}
	case TokenNE:
		if left.Type().Equal(types.I32) && right.Type().Equal(types.I32) {
			g.currentValue = g.currentBlock.NewICmp(enum.IPredNE, left, right)
		} else {
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredUNE, g.anyToReal(left), g.anyToReal(right))
		}
	case TokenLT:
		if left.Type().Equal(types.I32) && right.Type().Equal(types.I32) {
			// TODO(daniel): "signed" or "unsigned" integer compares?
			g.currentValue = g.currentBlock.NewICmp(enum.IPredSLT, left, right)
		} else {
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredULT, g.anyToReal(left), g.anyToReal(right))
		}
	case TokenLE:
		if left.Type().Equal(types.I32) && right.Type().Equal(types.I32) {
			g.currentValue = g.currentBlock.NewICmp(enum.IPredSLE, left, right)
		} else {
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredULE, g.anyToReal(left), g.anyToReal(right))
		}
	case TokenGE:
		if left.Type().Equal(types.I32) && right.Type().Equal(types.I32) {
			g.currentValue = g.currentBlock.NewICmp(enum.IPredSGE, left, right)
		} else {
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredUGE, g.anyToReal(left), g.anyToReal(right))
		}
	case TokenGT:
		if left.Type().Equal(types.I32) && right.Type().Equal(types.I32) {
			g.currentValue = g.currentBlock.NewICmp(enum.IPredSGT, left, right)
		} else {
			g.currentValue = g.currentBlock.NewFCmp(enum.FPredUGT, g.anyToReal(left), g.anyToReal(right))
		}
	default:
		panic(fmt.Sprintf("unsupported token type: %+v", n.token))
	}
}

func (g *generator) VisitNumberExpr(n *NumberExpr) {
	switch n.Type() {
	case TypeInt64:
		g.currentValue = constant.NewInt(types.I32, int64(n.token.Int))
	case TypeFloat64:
		g.currentValue = constant.NewFloat(types.Double, n.token.Real)
	}
}

func (g *generator) VisitStringExpr(n *StringExpr) {
	str := g.internString(n.token.Text)

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

func (g *generator) anyToReal(v value.Value) value.Value {
	if !v.Type().Equal(types.Double) {
		v = g.currentBlock.NewSIToFP(v, types.Double)
	}

	return v
}

func (g *generator) anyToInteger(v value.Value) value.Value {
	if !v.Type().Equal(types.I32) {
		v = g.currentBlock.NewFPToSI(v, types.I32)
	}

	return v
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
			panic(fmt.Sprintf("don't know how to print type: %s", arg.Type()))
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
	g.stdlib["puts"] = g.currentModule.NewFunc("puts", types.I32,
		ir.NewParam("str", types.I8Ptr))
	g.stdlib["rand"] = g.currentModule.NewFunc("rand", types.I32)

	sprintf := g.currentModule.NewFunc("sprintf", types.I32,
		ir.NewParam("buf", types.I8Ptr),
		ir.NewParam("format", types.I8Ptr))
	sprintf.Sig.Variadic = true
	g.stdlib["sprintf"] = sprintf

	printf := g.currentModule.NewFunc("printf", types.I32,
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
