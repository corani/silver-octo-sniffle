package main

import (
	"fmt"
	"io"
	"strings"
)

func printAST(w io.Writer, root Node) {
	printer := &astPrinter{
		indent: 0,
		out:    w,
	}

	printer.Print(root)
}

type astPrinter struct {
	indent int
	out    io.Writer
}

var _ Visitor = (*astPrinter)(nil)

func (p *astPrinter) Print(root Node) {
	fmt.Fprintln(p.out, "AST:")
	root.Visit(p)
}

func (p *astPrinter) printf(format string, args ...any) {
	format = strings.Repeat("  ", p.indent) + format + "\n"

	fmt.Fprintf(p.out, format, args...)
}

func (p *astPrinter) VisitModule(n *Module) {
	p.printf("(module")
	p.indent++

	for _, stmt := range n.stmts {
		stmt.Visit(p)
	}

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitExprStmt(n *ExprStmt) {
	p.printf("(expr2stmt")
	p.indent++

	n.expr.Visit(p)

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitCallExpr(n *CallExpr) {
	p.printf("(%v [%v]", n.token.Text, n.typ)
	p.indent++

	for _, arg := range n.args {
		arg.Visit(p)
	}

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitBinaryExpr(n *BinaryExpr) {
	p.printf("(%v [%v]", n.token.Type, n.typ)
	p.indent++

	for _, arg := range n.args {
		arg.Visit(p)
	}

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitNumberExpr(n *NumberExpr) {
	p.printf("(number [%v] %d)", n.typ, n.token.Number)
}
