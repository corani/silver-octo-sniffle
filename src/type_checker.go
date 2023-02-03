package main

import "fmt"

func typeCheck(root Node) error {
	checker := &typeChecker{
		consts: make(map[string]ConstDecl),
		vars:   make(map[string]VarDecl),
		errors: nil,
	}

	return checker.Check(root)
}

type typeChecker struct {
	consts map[string]ConstDecl
	vars   map[string]VarDecl
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
	for i, decl := range m.consts {
		decl.expr.Visit(c)

		decl.typ = decl.expr.Type()
		if v := decl.expr.ConstValue(); v != nil {
			decl.value = *v
		} else {
			c.errors = append(c.errors, fmt.Errorf("initializer for %q is not constant", decl.token.Text))
		}

		m.consts[i] = decl
		c.consts[decl.token.Text] = decl
	}

	for i, decl := range m.vars {
		switch decl.typToken.Text {
		case "INTEGER":
			decl.typ = TypeInt64
		case "REAL":
			decl.typ = TypeFloat64
		case "BOOLEAN":
			decl.typ = TypeBoolean
		default:
			// TODO(daniel): support "CHAR", "SET", "ARRAY", "RECORD", "POINTER", ...
			c.errors = append(c.errors, fmt.Errorf("unknown type %q for variable %q",
				decl.typToken.Text, decl.token.Text))
		}

		m.vars[i] = decl
		c.vars[decl.token.Text] = decl
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
		if s.expr.Type() != t.typ {
			c.errors = append(c.errors, fmt.Errorf("can't assign type %q to variable %q (which is of type %q)",
				s.expr.Type(), s.token.Text, t.typ))
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

	if len(e.args) != 2 {
		c.errors = append(c.errors, fmt.Errorf("binary expression with %d arguments", len(e.args)))
	}

	switch {
	case e.token.Type == TokenSlash:
		e.typ = TypeFloat64

		lhs := e.args[0].ConstValue()
		rhs := e.args[1].ConstValue()

		if lhs != nil && rhs != nil {
			e.constValue = &Value{
				typ:  e.Type(),
				real: lhs.Real() / rhs.Real(),
			}
		}
	case e.token.isRelation():
		e.typ = TypeBoolean

		// TODO(daniel): const evaluation
	case e.token.Type == TokenAmpersand || e.token.Type == TokenOR:
		if e.args[0].Type() != TypeBoolean || e.args[1].Type() != TypeBoolean {
			c.errors = append(c.errors, fmt.Errorf("can only %s boolean values", e.token.Type))
		}

		e.typ = TypeBoolean

		// TODO(daniel): const evaluation
	case e.args[0].Type() != e.args[1].Type():
		e.typ = TypeFloat64

		// TODO(daniel): const evaluation
	default:
		e.typ = e.args[0].Type()

		// TODO(daniel): const evaluation
	}
}

func (c *typeChecker) VisitDesignatorExpr(e *DesignatorExpr) {
	if t, ok := c.consts[e.token.Text]; ok {
		e.typ = t.typ
		e.kind = KindConst
		e.constValue = t.expr.ConstValue()
	} else if t, ok := c.vars[e.token.Text]; ok {
		e.typ = t.typ
		e.kind = KindVar
	} else {
		c.errors = append(c.errors, fmt.Errorf("undefined identifier %q", e.token.Text))
	}
}

func (c *typeChecker) VisitNumberExpr(e *NumberExpr) {
	switch e.token.Type {
	case TokenInteger:
		e.typ = TypeInt64
		e.constValue = &Value{
			typ:     e.Type(),
			integer: e.token.Int,
		}
	case TokenReal:
		e.typ = TypeFloat64
		e.constValue = &Value{
			typ:  e.Type(),
			real: e.token.Real,
		}
	}
}

func (c *typeChecker) VisitStringExpr(e *StringExpr) {
	e.constValue = &Value{
		typ:    e.Type(),
		String: e.token.Text,
	}
}

func (c *typeChecker) VisitBooleanExpr(e *BooleanExpr) {
	var v bool

	switch e.token.Text {
	case "TRUE":
		v = true
	case "FALSE":
		v = false
	default:
		c.errors = append(c.errors, fmt.Errorf("invalid boolean literal: %q", e.token.Text))
	}

	e.constValue = &Value{
		typ:  e.Type(),
		Bool: v,
	}
}

func (c *typeChecker) VisitNotExpr(e *NotExpr) {
	e.expr.Visit(c)

	if e.expr.Type() != TypeBoolean {
		c.errors = append(c.errors, fmt.Errorf("`~` not supported for type %v", e.expr.Type()))
	}

	if c := e.expr.ConstValue(); c != nil {
		e.constValue = &Value{
			typ:  e.Type(),
			Bool: !c.Bool,
		}
	}
}
