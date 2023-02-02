package main

import "fmt"

func typeCheck(root Node) error {
	checker := &typeChecker{
		vars:   make(map[string]Type),
		errors: nil,
	}

	return checker.Check(root)
}

type typeChecker struct {
	vars   map[string]Type
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
	if m.vars != nil {
		for k, v := range m.vars {
			switch v.Text {
			case "INTEGER":
				c.vars[k.Text] = TypeInt64
			case "REAL":
				c.vars[k.Text] = TypeFloat64
			case "BOOLEAN":
				c.vars[k.Text] = TypeBoolean
			default:
				// TODO(daniel): support "CHAR", "SET", "ARRAY", "RECORD", "POINTER", ...
				c.errors = append(c.errors, fmt.Errorf("unknown type %q for variable %q", v.Text, k.Text))
			}
		}
	}

	m.stmts.Visit(c)
}

func (c *typeChecker) VisitStmtSequence(s *StmtSequence) {
	for _, v := range s.stmts {
		v.Visit(c)
	}
}

func (c *typeChecker) VisitAssignStmt(s *AssignStmt) {
	s.expr.Visit(c)

	if t, ok := c.vars[s.token.Text]; ok {
		if s.expr.Type() != t {
			c.errors = append(c.errors, fmt.Errorf("can't assign type %q to variable %q (which is of type %q)",
				s.expr.Type(), s.token.Text, t))
		}
	} else {
		c.errors = append(c.errors, fmt.Errorf("undefined identifier %q", s.token.Text))
	}
}

func (c *typeChecker) VisitIfStmt(s *IfStmt) {
	s.expr.Visit(c)
	if s.expr.Type() != TypeBoolean {
		c.errors = append(c.errors, fmt.Errorf("condition for IF must be boolean, got %v", s.expr.Type()))
	}

	s.trueBlock.Visit(c)
	s.falseBlock.Visit(c)
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

func (c *typeChecker) VisitDesignatorExpr(e *DesignatorExpr) {
	if t, ok := c.vars[e.token.Text]; ok {
		e.typ = t
		e.kind = KindVar
	} else {
		c.errors = append(c.errors, fmt.Errorf("undefined identifier %q", e.token.Text))
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
