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

func generateIR(writer io.Writer, root Node) {
	g := &generator{
		currentModule: ir.NewModule(),
		currentBlock:  nil,
		currentValue:  nil,
		stdlib:        make(map[string]*ir.Func),
		strings:       make(map[string]*ir.Global),
		buf:           nil,
	}

	g.Generate(root)

	fmt.Fprintln(writer, g.currentModule)
}

type generator struct {
	currentModule *ir.Module
	currentBlock  *ir.Block
	currentValue  value.Value
	stdlib        map[string]*ir.Func
	strings       map[string]*ir.Global
	buf           *ir.InstAlloca
}

var _ Visitor = (*generator)(nil)

func (g *generator) Generate(root Node) {
	g.currentModule.TargetTriple = g.getTargetTriple()
	g.generateStdlib()

	root.Visit(g)
}

func (g *generator) VisitModule(n *Module) {
	main := g.currentModule.NewFunc("main", types.I32)
	g.stdlib["main"] = main

	g.currentBlock = main.NewBlock("entry")

	for _, stmt := range n.stmts {
		stmt.Visit(g)
	}

	g.currentBlock.NewRet(constant.NewInt(types.I32, 0))
}

func (g *generator) VisitPrintStmt(n *PrintStmt) {
	for _, arg := range n.args {
		arg.Visit(g)
		number := g.currentValue

		if g.buf == nil {
			g.buf = g.currentBlock.NewAlloca(types.NewArray(318, types.I8))
		}

		bufptr := g.currentBlock.NewGetElementPtr(g.buf.ElemType, g.buf, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))

		format := g.intern("%d\000")
		formatptr := g.currentBlock.NewGetElementPtr(format.ContentType, format, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))

		g.currentBlock.NewCall(g.stdlib["sprintf"], bufptr, formatptr, number)
		g.currentBlock.NewCall(g.stdlib["puts"], bufptr)
	}
}

func (g *generator) VisitBinaryExpr(n *BinaryExpr) {
	n.args[0].Visit(g)
	left := g.currentValue
	n.args[1].Visit(g)
	right := g.currentValue

	switch n.token.Type {
	case TokenMinus:
		g.currentValue = g.currentBlock.NewSub(left, right)
	case TokenPlus:
		g.currentValue = g.currentBlock.NewAdd(left, right)
	case TokenStar:
		g.currentValue = g.currentBlock.NewMul(left, right)
	case TokenSlash:
		g.currentValue = g.currentBlock.NewSDiv(left, right)
	}
}

func (g *generator) VisitNumberExpr(n *NumberExpr) {
	g.currentValue = constant.NewInt(types.I32, int64(n.token.Number))
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
}

func (g *generator) intern(s string) *ir.Global {
	if v, ok := g.strings[s]; ok {
		return v
	}

	def := g.currentModule.NewGlobalDef("", constant.NewCharArrayFromString(s))
	g.strings[s] = def

	return def
}

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
