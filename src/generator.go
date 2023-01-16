package main

import (
	"fmt"
	"io"
)

func generateIR(writer io.Writer, root Node) {
	g := &generator{
		out:   writer,
		ssaid: 1,
		conid: 0,
	}

	g.Generate(root)
}

type generator struct {
	out   io.Writer
	ssaid int
	conid int
}

var _ Visitor = (*generator)(nil)

func (g *generator) Generate(root Node) {
	fmt.Fprintln(g.out, `target datalayout = "e-m:e-p270:32:32-p271:32:32-p272:64:64-i64:64-f80:128-n8:16:32:64-S128"`)
	fmt.Fprintln(g.out, `target triple = "x86_64-pc-linux-gnu"`)
	format := "%d\000" // this nonsense is needed to keep govet quiet
	fmt.Fprintln(g.out, `@format = internal constant [3 x i8] c"`+format+`"`)
	fmt.Fprintln(g.out, `define dso_local i32 @main() {`)

	root.Visit(g)

	fmt.Fprintln(g.out, `  ret i32 0`)
	fmt.Fprintln(g.out, `}`)
	fmt.Fprintln(g.out, `declare dso_local i32 @puts(i8*)`)
	fmt.Fprintln(g.out, `declare dso_local i32 @rand()`)
	fmt.Fprintln(g.out, `declare dso_local i32 @sprintf(i8*, i8*, ...)`)
}

func (g *generator) VisitNegateExpr(n *NegateExpr) {
	n.expr.Visit(g)
	// TODO(daniel): generate negate last register.
}

func (g *generator) VisitNumberExpr(n *NumberExpr) {
	val := g.assign("alloca i32")
	fmt.Fprintf(g.out, "store i32 %v, i32* %v\n", n.token.Number, val)
	g.assign("load i32, i32* %v", val)
}

func (g *generator) VisitBinaryExpr(n *BinaryExpr) {
	n.args[0].Visit(g)
	left := fmt.Sprintf("%%%d", g.ssaid-1)
	n.args[1].Visit(g)
	right := fmt.Sprintf("%%%d", g.ssaid-1)
	switch n.token.Type {
	case TokenMinus:
		g.assign("sub i32 %v, %v", left, right)
	case TokenPlus:
		g.assign("add i32 %v, %v", left, right)
	}
}

func (g *generator) VisitPrintStmt(n *PrintStmt) {
	for _, arg := range n.args {
		arg.Visit(g)
		number := fmt.Sprintf("%%%d", g.ssaid-1)

		buf := g.assign("alloca [318 x i8]")
		bufptr := g.assign("getelementptr inbounds [318 x i8], [318 x i8]* %v, i64 0, i64 0", buf)
		formatptr := g.assign(`getelementptr inbounds [3 x i8], [3 x i8]* @format, i64 0, i64 0`)
		g.assign(`call i32 (i8*, i8*, ...) @sprintf(i8* %v, i8* %v, i32 %v)`, bufptr, formatptr, number)
		g.assign(`call i32 @puts(i8* %v)`, bufptr)
	}
}

func (g *generator) assign(format string, args ...any) string {
	id := g.ssaid
	g.ssaid++

	fmt.Fprintf(g.out, "%%%d = %v\n", id, fmt.Sprintf(format, args...))

	return fmt.Sprintf("%%%d", id)
}
