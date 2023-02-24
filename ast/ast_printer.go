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

func (p *astPrinter) visitDecls(decls []Decl) {
	var (
		consts []*ConstDecl
		types  []TypeDecl
		vars   []*VarDecl
		procs  []*ProcDecl
	)

	for _, decl := range decls {
		switch d := decl.(type) {
		case *ConstDecl:
			consts = append(consts, d)
		case *VarDecl:
			vars = append(vars, d)
		case *ProcDecl:
			procs = append(procs, d)
		case TypeDecl:
			types = append(types, d)
		}
	}

	if len(consts) != 0 {
		p.printf("(consts")
		p.indent++

		for _, decl := range consts {
			decl.Visit(p)
		}

		p.indent--
		p.printf(")")
	}

	if len(types) != 0 {
		p.printf("(types")
		p.indent++

		for _, decl := range types {
			p.printf("(%v", decl.Token().Text)
			p.indent++
			decl.Visit(p)
			p.indent--
			p.printf(")")
		}

		p.indent--
		p.printf(")")
	}

	if len(vars) != 0 {
		p.printf("(vars")
		p.indent++

		for _, decl := range vars {
			decl.Visit(p)
		}

		p.indent--
		p.printf(")")
	}

	if len(procs) != 0 {
		p.printf("(procs")
		p.indent++

		for _, decl := range procs {
			decl.Visit(p)
		}

		p.indent--
		p.printf(")")
	}
}

func (p *astPrinter) VisitModule(n *Module) {
	p.printf("(module %q", n.name)
	p.indent++

	p.visitDecls(n.Decls())

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
	p.printf("(assign")
	p.indent++

	n.Designator().Visit(p)
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

func (p *astPrinter) VisitReturnStmt(n *ReturnStmt) {
	p.printf("(return")
	p.indent++

	n.Expr().Visit(p)

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
	p.printf("(call")
	p.indent++

	n.Designator().Visit(p)

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
	q := n.QualIdent()

	name := q.Ident().Text
	if q.Qualifier() != nil {
		name = q.Qualifier().Text + "." + name
	}

	// TODO(daniel): do this recursively?
	for _, s := range n.Selectors() {
		switch sel := s.(type) {
		case *DerefSelector:
			name = fmt.Sprintf("%s^", name)
		case *DotSelector:
			name = fmt.Sprintf("%s.%s", name, sel.Ident().Text)
		case *BracketSelector:
			// TODO(daniel): print expressions.
			name = fmt.Sprintf("%s[...]", name)
		default:
		}
	}

	p.printf("(%v [%v] %q)", n.kind, n.typ, name)
}

func (p *astPrinter) VisitNotExpr(n *NotExpr) {
	p.printf("(not [%v]", n.Type())
	p.indent++

	n.Expr().Visit(p)

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitNumberLit(n *NumberLit) {
	switch n.Type() {
	case TypeInt64:
		p.printf("(number [%v] %d)", n.typ, n.Token().Int)
	case TypeFloat64:
		p.printf("(number [%v] %f)", n.typ, n.Token().Real)
	}
}

func (p *astPrinter) VisitStringLit(n *StringLit) {
	p.printf("(string %q)", n.Token().Text)
}

func (p *astPrinter) VisitCharLit(n *CharLit) {
	p.printf("(char %q)", n.Token().Text)
}

func (p *astPrinter) VisitBooleanLit(n *BooleanLit) {
	if n.Token().Bool {
		p.printf("#true")
	} else {
		p.printf("#false")
	}
}

func (p *astPrinter) VisitSetLit(n *SetLit) {
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

func (p *astPrinter) VisitConstDecl(decl *ConstDecl) {
	p.printf("(%v [%v]", decl.Token().Text, decl.Type())
	p.indent++

	decl.Expr().Visit(p)

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitTypeBaseDecl(decl *TypeBaseDecl) {
	p.printf("(%v [%v])", decl.Token().Text, decl.Type())
}

func (p *astPrinter) VisitTypePointerDecl(decl *TypePointerDecl) {
	p.printf("(pointer")
	p.indent++

	decl.To().Visit(p)

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitVarDecl(decl *VarDecl) {
	p.printf("(%v ", decl.Token().Text)
	p.indent++

	if decl.TypeDecl() != nil {
		decl.TypeDecl().Visit(p)
	} else if decl.TypeRef() != nil {
		decl.TypeRef().Visit(p)
	}

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitProcDecl(decl *ProcDecl) {
	p.printf("(%v [", decl.Token().Text)
	p.indent += 2

	if decl.ReturnType() != nil {
		decl.ReturnType().Visit(p)
	} else {
		p.printf("VOID")
	}

	p.indent--
	p.printf("]")

	p.visitDecls(decl.Decls())
	decl.Stmts().Visit(p)

	p.indent--
	p.printf(")")
}

func (p *astPrinter) VisitTypeRef(ref *TypeRef) {
	p.printf("(typeref %v)", ref.Token().Text)
}
