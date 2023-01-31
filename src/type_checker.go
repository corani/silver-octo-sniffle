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
	case e.token.isRelation():
		e.typ = TypeBoolean
	case e.token.Type == TokenAnd || e.token.Type == TokenOr:
		if e.args[0].Type() != TypeBoolean || e.args[1].Type() != TypeBoolean {
			panic(fmt.Sprintf("can only %s boolean values", e.token.Type))
		}

		e.typ = TypeBoolean
	case e.args[0].Type() != e.args[1].Type():
		e.typ = TypeFloat64
	default:
		e.typ = e.args[0].Type()
	}
}

func (c *typeChecker) VisitNumberExpr(e *NumberExpr) {
	switch e.token.Type {
	case TokenInteger:
		e.typ = TypeInt64
	case TokenReal:
		e.typ = TypeFloat64
	}
}

func (c *typeChecker) VisitStringExpr(e *StringExpr) {
	// nop
}

func (c *typeChecker) VisitBooleanExpr(e *BooleanExpr) {
	// nop
}

func (c *typeChecker) VisitNotExpr(e *NotExpr) {
	e.expr.Visit(c)

	if e.expr.Type() != TypeBoolean {
		panic(fmt.Sprintf("`~` not supported for type %v", e.expr.Type()))
	}
}
