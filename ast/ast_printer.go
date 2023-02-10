package ast

import (
	"fmt"
	"io"
	"strings"
)

func PrintAST(w io.Writer, root Node) {
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

var _ AstVisitor = (*astPrinter)(nil)

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

	if len(n.Consts()) != 0 {
		p.printf("(consts")
		p.indent++

		for _, decl := range n.Consts() {
			p.printf("(%v [%v]", decl.Token().Text, decl.Type())
			p.indent++

			decl.Expr().Visit(p)

			p.indent--
			p.printf(")")
		}

		p.indent--
		p.printf(")")
	}

	if len(n.Vars()) != 0 {
		p.printf("(vars")
		p.indent++

		for _, decl := range n.Vars() {
			p.printf("(%v [%v])", decl.Token().Text, decl.Type())
		}

		p.indent--
		p.printf(")")
	}

	n.Block().Visit(p)

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitInvalidStmt(n *InvalidStmt) {
	p.printf("(invalid %v)", n.Token().Text)
}

func (p *astPrinter) VisitStmtSequence(n *StmtSequence) {
	p.printf("(stmts")
	p.indent++

	for _, stmt := range n.Stmts() {
		stmt.Visit(p)
	}

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitAssignStmt(n *AssignStmt) {
	p.printf("(assign %v", n.Token().Text)
	p.indent++

	n.Expr().Visit(p)

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitIfStmt(n *IfStmt) {
	p.printf("(if")
	p.indent++

	n.Expr().Visit(p)
	n.TrueBlock().Visit(p)
	n.FalseBlock().Visit(p)

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitRepeatStmt(n *RepeatStmt) {
	p.printf("(repeat")
	p.indent++

	cond := n.Condition()
	cond.Stmt().Visit(p)
	cond.Expr().Visit(p)

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitWhileStmt(n *WhileStmt) {
	p.printf("(while")
	p.indent++

	for _, c := range n.Conditions() {
		p.printf("(cond")
		p.indent++

		c.Expr().Visit(p)
		c.Stmt().Visit(p)

		p.indent--
		p.printf(")")
	}

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitForStmt(n *ForStmt) {
	p.printf("(for %v", n.Iter().Text)
	p.indent++

	n.From().Visit(p)
	n.To().Visit(p)
	n.By().Visit(p)
	n.Stmt().Visit(p)

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitExprStmt(n *ExprStmt) {
	p.printf("(expr2stmt")
	p.indent++

	n.Expr().Visit(p)

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitInvalidExpr(n *InvalidExpr) {
	p.printf("(invalid %v)", n.Token().Text)
}

func (p *astPrinter) VisitCallExpr(n *CallExpr) {
	p.printf("(call %q [%v]", n.Token().Text, n.typ)
	p.indent++

	for _, arg := range n.Args() {
		arg.Visit(p)
	}

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitBinaryExpr(n *BinaryExpr) {
	p.printf("(%v [%v]", n.Token().Type, n.typ)
	p.indent++

	for _, arg := range n.Args() {
		arg.Visit(p)
	}

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitDesignatorExpr(n *DesignatorExpr) {
	p.printf("(%v [%v] %q)", n.kind, n.typ, n.Token().Text)
}

func (p *astPrinter) VisitNumberExpr(n *NumberExpr) {
	switch n.Type() {
	case TypeInt64:
		p.printf("(number [%v] %d)", n.typ, n.Token().Int)
	case TypeFloat64:
		p.printf("(number [%v] %f)", n.typ, n.Token().Real)
	}
}

func (p *astPrinter) VisitStringExpr(n *StringExpr) {
	p.printf("(string %q)", n.Token().Text)
}

func (p *astPrinter) VisitCharExpr(n *CharExpr) {
	p.printf("(char %q)", n.Token().Text)
}

func (p *astPrinter) VisitBooleanExpr(n *BooleanExpr) {
	if n.Token().Bool {
		p.printf("#true")
	} else {
		p.printf("#false")
	}
}

func (p *astPrinter) VisitSetExpr(n *SetExpr) {
	// condense ranges in the list of bits.
	var parts []string

	start := -1

	bits := n.Bits()

	for i := 0; i < len(bits); i++ {
		if i < len(bits)-1 && bits[i]+1 == bits[i+1] {
			if start < 0 {
				start = i
			}

			// doing a range
			continue
		}

		if start >= 0 {
			// finished range
			parts = append(parts, fmt.Sprintf("%d..%d", bits[start], bits[i]))
			start = -1
		} else {
			parts = append(parts, fmt.Sprintf("%d", bits[i]))
		}
	}

	p.printf("(set (%s))", strings.Join(parts, ", "))
}

func (p *astPrinter) VisitNotExpr(n *NotExpr) {
	p.printf("(not [%v]", n.Type())
	p.indent++

	n.Expr().Visit(p)

	p.indent--
	p.printf(")")
}
