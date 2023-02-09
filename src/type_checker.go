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
		case "CHAR":
			decl.typ = TypeChar
		case "SET":
			decl.typ = TypeSet
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

func (c *typeChecker) VisitForStmt(s *ForStmt) {
	iter, ok := c.vars[s.iter.Text]
	if !ok {
		c.errors = append(c.errors, fmt.Errorf("FOR iterator must be a variable"))
	}

	// TODO(daniel): is this true?
	if iter.typ != TypeInt64 {
		c.errors = append(c.errors, fmt.Errorf("FOR iterator must be an integer, got %v",
			iter.typ))
	}

	s.from.Visit(c)
	s.to.Visit(c)
	s.by.Visit(c)

	if s.from.Type() != iter.typ {
		c.errors = append(c.errors, fmt.Errorf("FOR iterator FROM must be an integer, got %v",
			s.from.Type()))
	}

	if s.to.Type() != iter.typ {
		c.errors = append(c.errors, fmt.Errorf("FOR iterator TO must be an integer, got %v",
			s.to.Type()))
	}

	if s.by.Type() != iter.typ {
		c.errors = append(c.errors, fmt.Errorf("FOR iterator BY must be an integer, got %v",
			s.by.Type()))
	}

	if s.by.ConstValue() == nil {
		c.errors = append(c.errors, fmt.Errorf("FOR iterator BY must be a constant expression"))
	} else if s.by.ConstValue().Int() == 0 {
		c.errors = append(c.errors, fmt.Errorf("FOR iterator BY must not be zero"))
	}

	if s.from.ConstValue() != nil && s.to.ConstValue() != nil && s.by.ConstValue() != nil {
		from := s.from.ConstValue().Int()
		to := s.to.ConstValue().Int()
		by := s.by.ConstValue().Int()

		if from > to && by > 0 {
			c.errors = append(c.errors, fmt.Errorf("FOR iterator BY must be negative if FROM > BY"))
		} else if from < to && by < 0 {
			c.errors = append(c.errors, fmt.Errorf("FOR iterator BY must be positive if FROM < BY"))
		}
	}

	s.stmt.Visit(c)
}

func (c *typeChecker) VisitExprStmt(s *ExprStmt) {
	s.expr.Visit(c)
}

func (c *typeChecker) VisitCallExpr(e *CallExpr) {
	for _, v := range e.args {
		v.Visit(c)
	}

	if v, ok := builtin[e.Token().Text]; ok {
		t, err := v.Validate(e.args)
		if err != nil {
			c.errors = append(c.errors, err)

			return
		}

		cv, err := v.ConstValue(e.args)
		if err != nil {
			c.errors = append(c.errors, err)

			return
		}

		e.typ = t
		e.constValue = cv
	} else {
		c.errors = append(c.errors, fmt.Errorf("builtin function %q not found", e.Token().Text))
	}
}

func (c *typeChecker) VisitBinaryExpr(e *BinaryExpr) {
	for _, v := range e.args {
		v.Visit(c)
	}

	switch {
	case len(e.args) != 2:
		c.errors = append(c.errors, fmt.Errorf("binary expression with %d arguments", len(e.args)))
	case e.token.Type == TokenIN:
		e.typ = TypeBoolean
	case e.args[0].Type() != e.args[1].Type():
		e.typ = TypeVoid

		c.errors = append(c.errors, fmt.Errorf("different types for binary operator %s",
			e.token.Type))
	case e.token.Type == TokenSlash:
		if e.args[0].Type() == TypeSet && e.args[1].Type() == TypeSet {
			e.typ = TypeSet
		} else {
			e.typ = TypeFloat64

			if e.args[0].Type() != e.Type() {
				c.errors = append(c.errors, fmt.Errorf("unexpected type for binary operator %s",
					e.token.Type))
			}
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

func (c *typeChecker) VisitCharExpr(e *CharExpr) {
	e.constValue = evaluateCharExpr(e.token)
}

func (c *typeChecker) VisitBooleanExpr(e *BooleanExpr) {
	if v, err := evaluateBooleanExpr(e.token); err != nil {
		c.errors = append(c.errors, err)
	} else {
		e.constValue = v
	}
}

func (c *typeChecker) VisitSetExpr(e *SetExpr) {
	e.constValue = evaluateSetExpr(e)
}

func (c *typeChecker) VisitNotExpr(e *NotExpr) {
	e.expr.Visit(c)

	if v, err := evaluateNotExpr(e.expr); err != nil {
		c.errors = append(c.errors, err)
	} else {
		e.constValue = v
	}
}
