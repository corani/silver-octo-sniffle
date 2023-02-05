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

func (c *typeChecker) VisitRepeatStmt(s *RepeatStmt) {
	s.cond.expr.Visit(c)
	s.cond.stmt.Visit(c)

	if s.cond.expr.Type() != TypeBoolean {
		c.errors = append(c.errors, fmt.Errorf("condition for REPEAT must be boolean, got %v",
			s.cond.expr.Type()))
	}
}

func (c *typeChecker) VisitWhileStmt(s *WhileStmt) {
	for _, cond := range s.conds {
		cond.expr.Visit(c)
		cond.stmt.Visit(c)

		if cond.expr.Type() != TypeBoolean {
			c.errors = append(c.errors, fmt.Errorf("condition for WHILE must be boolean, got %v",
				cond.expr.Type()))
		}
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
	switch e.Token().Text {
	case "print":
		e.typ = TypeVoid
	case "INC", "DEC":
		if c.isLen(e.args, 1) && c.isVar(e.args[0]) && c.isType(e.args[0], TypeInt64) {
			e.typ = TypeVoid
		}
	case "FLT":
		if c.isLen(e.args, 1) && c.isType(e.args[0], TypeInt64) {
			e.typ = TypeFloat64

			if v := e.args[0].ConstValue(); v != nil {
				e.constValue = &Value{
					typ:  e.typ,
					real: v.Real(),
				}
			}
		}
	case "FLOOR":
		if c.isLen(e.args, 1) && c.isType(e.args[0], TypeFloat64) {
			e.typ = TypeInt64

			if v := e.args[0].ConstValue(); v != nil {
				e.constValue = &Value{
					typ:     e.typ,
					integer: v.Int(),
				}
			}
		}
	default:
		c.errors = append(c.errors, fmt.Errorf("don't know how to call %q",
			e.Token().Text))
	}
}

func (c *typeChecker) isLen(e []Expr, exp int) bool {
	if len(e) != exp {
		c.errors = append(c.errors, fmt.Errorf("expected %d arguments, got %d", exp, len(e)))

		return false
	}

	return true
}

func (c *typeChecker) isVar(e Expr) bool {
	if e.Token().Type != TokenIdent {
		c.errors = append(c.errors, fmt.Errorf("expected identifier, got %s", e.Token().Type))

		return false
	} else if _, ok := c.vars[e.Token().Text]; !ok {
		c.errors = append(c.errors, fmt.Errorf("%q must be a variable", e.Token().Text))

		return false
	}

	return true
}

func (c *typeChecker) isType(e Expr, exp Type) bool {
	if e.Type() != exp {
		c.errors = append(c.errors, fmt.Errorf("expected type %s, got %s", exp, e.Type()))

		return false
	}

	return true
}

func (c *typeChecker) VisitBinaryExpr(e *BinaryExpr) {
	for _, v := range e.args {
		v.Visit(c)
	}

	if len(e.args) != 2 {
		c.errors = append(c.errors, fmt.Errorf("binary expression with %d arguments", len(e.args)))
	}

	switch {
	case e.args[0].Type() != e.args[1].Type():
		e.typ = TypeVoid

		c.errors = append(c.errors, fmt.Errorf("different types for binary operator %s",
			e.token.Type))
	case e.token.Type == TokenSlash:
		e.typ = TypeFloat64

		if e.args[0].Type() != e.Type() {
			c.errors = append(c.errors, fmt.Errorf("unexpected type for binary operator %s",
				e.token.Type))
		}
	case e.token.isRelation():
		e.typ = TypeBoolean
	case e.token.Type == TokenAmpersand || e.token.Type == TokenOR:
		e.typ = TypeBoolean

		if e.args[0].Type() != e.Type() {
			c.errors = append(c.errors, fmt.Errorf("unexpected type for binary operator %s",
				e.token.Type))
		}
	default:
		e.typ = e.args[0].Type()
	}

	lhs := e.args[0].ConstValue()
	rhs := e.args[1].ConstValue()

	if lhs != nil && rhs != nil {
		e.constValue = evaluateBinaryExpr(e.token.Type, e.Type(), lhs, rhs)
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
	if v, err := evaluateNumberExpr(e.token); err != nil {
		c.errors = append(c.errors, err)
	} else {
		e.constValue = v
		e.typ = v.Type()
	}
}

func (c *typeChecker) VisitStringExpr(e *StringExpr) {
	e.constValue = evaluateStringExpr(e.token)
}

func (c *typeChecker) VisitBooleanExpr(e *BooleanExpr) {
	if v, err := evaluateBooleanExpr(e.token); err != nil {
		c.errors = append(c.errors, err)
	} else {
		e.constValue = v
	}
}

func (c *typeChecker) VisitNotExpr(e *NotExpr) {
	e.expr.Visit(c)

	if v, err := evaluateNotExpr(e.expr); err != nil {
		c.errors = append(c.errors, err)
	} else {
		e.constValue = v
	}
}
