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
	fmt.Fprintln(p.out, "## AST")
	fmt.Fprintln(p.out, "```scheme")
	root.Visit(p)
	fmt.Fprintln(p.out, "```")
}

func (p *astPrinter) printf(format string, args ...any) {
	format = strings.Repeat("  ", p.indent) + format + "\n"

	fmt.Fprintf(p.out, format, args...)
}

func (p *astPrinter) VisitModule(n *Module) {
	p.printf("(module")
	p.indent++

	n.stmts.Visit(p)

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitStmtSequence(n *StmtSequence) {
	p.printf("(stmts")
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
	switch n.Type() {
	case TypeInt64:
		p.printf("(number [%v] %d)", n.typ, n.token.Int)
	case TypeFloat64:
		p.printf("(number [%v] %f)", n.typ, n.token.Real)
	}
}

func (p *astPrinter) VisitStringExpr(n *StringExpr) {
	p.printf("(string %q)", n.token.Text)
}

func (p *astPrinter) VisitBooleanExpr(n *BooleanExpr) {
	if n.token.Bool {
		p.printf("#true")
	} else {
		p.printf("#false")
	}
}

func (p *astPrinter) VisitNotExpr(n *NotExpr) {
	p.printf("(not [%v]", n.Type())
	p.indent++

	n.expr.Visit(p)

	p.indent--
	p.printf(")")
}
