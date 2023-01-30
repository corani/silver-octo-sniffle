package main

import (
	"fmt"
	"io"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
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
	main := g.currentModule.NewFunc("main", types.I32)
	g.currentBlock = main.NewBlock("entry")

	for _, stmt := range n.stmts {
		stmt.Visit(g)
	}

	g.currentBlock.NewRet(zero)
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

	// TODO(daniel): is automatic conversion a thing?
	// TODO(daniel): this assumes only `I32` and `Double` types.
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
	case TokenStar:
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
	case TokenDiv:
		// INTEGER division
		g.currentValue = g.currentBlock.NewSDiv(g.anyToInteger(left), g.anyToInteger(right))
	case TokenMod:
		// INTEGER modulus
		g.currentValue = g.currentBlock.NewSRem(g.anyToInteger(left), g.anyToInteger(right))
	case TokenAmpersand:
		// logical AND
	case TokenOr:
		// logical OR
	default:
		panic(fmt.Sprintf("unsupported token type: %+v", n.token))
	}
}

func (g *generator) VisitNumberExpr(n *NumberExpr) {
	g.currentValue = constant.NewInt(types.I32, int64(n.token.Number))
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
			number := g.visitAndReturnValue(arg)

			format := g.internString("%d\n\000")
			formatptr := g.currentBlock.NewGetElementPtr(format.ContentType, format, zero, zero)

			g.currentValue = g.currentBlock.NewCall(g.stdlib["printf"], formatptr, number)
		case TypeFloat64:
			number := g.visitAndReturnValue(arg)

			format := g.internString("%f\n\000")
			formatptr := g.currentBlock.NewGetElementPtr(format.ContentType, format, zero, zero)

			g.currentValue = g.currentBlock.NewCall(g.stdlib["printf"], formatptr, number)
		case TypeString:
			str := g.visitAndReturnValue(arg)

			g.currentValue = g.currentBlock.NewCall(g.stdlib["puts"], str)
		default:
			panic(fmt.Sprintf("don't know how to print type: %s", arg.Type()))
		}
	}
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
