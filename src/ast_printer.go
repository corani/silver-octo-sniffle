package main

import (
	"fmt"
	"strings"
)

func printAST(root Node) {
	printer := &astPrinter{
		indent: 0,
	}

	printer.Print(root)
}

type astPrinter struct {
	indent int
}

var _ Visitor = (*astPrinter)(nil)

func (p *astPrinter) Print(root Node) {
	fmt.Println("AST:")
	root.Visit(p)
}

func (p *astPrinter) printf(format string, args ...any) {
	format = strings.Repeat("  ", p.indent) + format + "\n"

	fmt.Printf(format, args...)
}

func (p *astPrinter) VisitNegateExpr(n *NegateExpr) {
	p.printf("(negate")
	p.indent++

	n.expr.Visit(p)

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitNumberExpr(n *NumberExpr) {
	p.printf("(number %d)", n.token.Number)
}

func (p *astPrinter) VisitBinaryExpr(n *BinaryExpr) {
	p.printf("(%v", n.token.Type)
	p.indent++

	for _, arg := range n.args {
		arg.Visit(p)
	}

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitPrintStmt(n *PrintStmt) {
	p.printf("(print")
	p.indent++

	for _, arg := range n.args {
		arg.Visit(p)
	}

	p.indent--
	p.printf(")")
}
