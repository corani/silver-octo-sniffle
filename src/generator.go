package main

import (
	"fmt"
	"io"
	"runtime"

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
	g.currentModule.TargetTriple = g.getTargetTriple()
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

func (g *generator) VisitPrintStmt(n *PrintStmt) {
	for _, arg := range n.args {
		number := g.visitAndReturnValue(arg)

		format := g.internString("%d\n\000")
		formatptr := g.currentBlock.NewGetElementPtr(format.ContentType, format, zero, zero)

		g.currentBlock.NewCall(g.stdlib["printf"], formatptr, number)
	}
}

func (g *generator) VisitBinaryExpr(n *BinaryExpr) {
	if len(n.args) != 2 {
		panic(fmt.Sprintf("expected two arguments, got %d", len(n.args)))
	}

	left := g.visitAndReturnValue(n.args[0])
	right := g.visitAndReturnValue(n.args[1])

	switch n.token.Type {
	case TokenMinus:
		g.currentValue = g.currentBlock.NewSub(left, right)
	case TokenPlus:
		g.currentValue = g.currentBlock.NewAdd(left, right)
	case TokenStar:
		g.currentValue = g.currentBlock.NewMul(left, right)
	case TokenSlash:
		g.currentValue = g.currentBlock.NewSDiv(left, right)
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

// incomplete and probably wrong!
func (g *generator) getTargetTriple() string {
	var os string

	switch runtime.GOOS {
	case "linux":
		os = "linux"
	default:
		panic(fmt.Sprintf("unknown OS: %v", runtime.GOOS))
	}

	switch runtime.GOARCH {
	case "amd64":
		return fmt.Sprintf("x86_64-pc-%v-gnu", os)
	case "arm64":
		return fmt.Sprintf("aarch64-unknown-%v-gnu", os)
	default:
		panic(fmt.Sprintf("unknown architecture: %v", runtime.GOARCH))
	}
}
