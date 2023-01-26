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

func (p *astPrinter) VisitModule(n *Module) {
	p.printf("(module")
	p.indent++

	for _, stmt := range n.stmts {
		stmt.Visit(p)
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
