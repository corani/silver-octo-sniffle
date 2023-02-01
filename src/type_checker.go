package main

import "fmt"

func typeCheck(root Node) error {
	checker := &typeChecker{}

	return checker.Check(root)
}

type typeChecker struct {
	errors []error
}

var _ Visitor = (*typeChecker)(nil)

func (c *typeChecker) Check(root Node) error {
	root.Visit(c)

	return c.Error()
}

// TODO(daniel): improve error reporting. Maybe with a callback?
func (c *typeChecker) Error() error {
	if len(c.errors) == 0 {
		return nil
	}

	txt := "type check errors"

	for _, v := range c.errors {
		txt = txt + ": " + v.Error()
	}

	return fmt.Errorf(txt)
}

func (c *typeChecker) VisitModule(m *Module) {
	m.stmts.Visit(c)
}

func (c *typeChecker) VisitStmtSequence(s *StmtSequence) {
	for _, v := range s.stmts {
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
		c.errors = append(c.errors, fmt.Errorf("don't know how to call %q", e.token.Text))
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
	case e.token.Type == TokenAmpersand || e.token.Type == TokenOR:
		if e.args[0].Type() != TypeBoolean || e.args[1].Type() != TypeBoolean {
			c.errors = append(c.errors, fmt.Errorf("can only %s boolean values", e.token.Type))
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
		c.errors = append(c.errors, fmt.Errorf("`~` not supported for type %v", e.expr.Type()))
	}
}
