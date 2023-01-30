package main

import "fmt"

func typeCheck(root Node) {
	checker := &typeChecker{}

	checker.Check(root)
}

type typeChecker struct{}

var _ Visitor = (*typeChecker)(nil)

func (c *typeChecker) Check(root Node) {
	root.Visit(c)
}

func (c *typeChecker) VisitModule(m *Module) {
	for _, v := range m.stmts {
		v.Visit(c)
	}
}

func (c *typeChecker) VisitExprStmt(s *ExprStmt) {
	s.expr.Visit(c)
}

func (c *typeChecker) VisitCallExpr(e *CallExpr) {
	for _, v := range e.args {
		v.Visit(c)
	}

	// TODO(daniel): return type of calls.
	switch e.token.Text {
	case "print":
		e.typ = TypeVoid
	default:
		panic(fmt.Sprintf("don't know how to call %q", e.token.Text))
	}
}

func (c *typeChecker) VisitBinaryExpr(e *BinaryExpr) {
	for _, v := range e.args {
		v.Visit(c)
	}

	switch {
	case e.token.Type == TokenSlash:
		e.typ = TypeFloat64
	case e.args[0].Type() != e.args[1].Type():
		e.typ = TypeFloat64
	default:
		e.typ = e.args[0].Type()
	}
}

func (c *typeChecker) VisitNumberExpr(e *NumberExpr) {
	e.typ = TypeInt64
}

func (c *typeChecker) VisitStringExpr(e *StringExpr) {
	// nop
}
