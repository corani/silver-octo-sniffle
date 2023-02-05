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
	p.printf("(module %q", n.name)
	p.indent++

	if len(n.consts) != 0 {
		p.printf("(consts")
		p.indent++

		for _, decl := range n.consts {
			p.printf("(%v [%v]", decl.token.Text, decl.typ)
			p.indent++

			decl.expr.Visit(p)

			p.indent--
			p.printf(")")
		}

		p.indent--
		p.printf(")")
	}

	if len(n.vars) != 0 {
		p.printf("(vars")
		p.indent++

		for _, decl := range n.vars {
			p.printf("(%v [%v])", decl.token.Text, decl.typ)
		}

		p.indent--
		p.printf(")")
	}

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

func (p *astPrinter) VisitAssignStmt(n *AssignStmt) {
	p.printf("(assign %v", n.token.Text)
	p.indent++

	n.expr.Visit(p)

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitIfStmt(n *IfStmt) {
	p.printf("(if")
	p.indent++

	n.expr.Visit(p)
	n.trueBlock.Visit(p)
	n.falseBlock.Visit(p)

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
	p.printf("(call %q [%v]", n.Token().Text, n.typ)
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

func (p *astPrinter) VisitDesignatorExpr(n *DesignatorExpr) {
	p.printf("(%v [%v] %q)", n.kind, n.typ, n.token.Text)
}

func (p *astPrinter) VisitNumberExpr(n *NumberExpr) {
	switch n.Type() {
	case TypeInt64:
		p.printf("(number [%v] %d)", n.typ, n.token.Int)
	case TypeFloat64:
		p.printf("(number [%v] %f)", n.typ, n.token.Real)
	default:
		p.printf("(number [%v] %q)", n.typ, n.token.Text)
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
