package check

import (
	"fmt"

	"github.com/corani/silver-octo-sniffle/ast"
	"github.com/corani/silver-octo-sniffle/lex"
)

func TypeCheck(root ast.Node) error {
	checker := &typeChecker{
		consts:       make(map[string]*ast.ConstDecl),
		vars:         make(map[string]*ast.VarDecl),
		builtins:     ast.GetBuiltinFunctions(),
		errors:       nil,
		currentValue: nil,
	}

	return checker.Check(root)
}

type typeChecker struct {
	consts       map[string]*ast.ConstDecl
	vars         map[string]*ast.VarDecl
	builtins     map[string]ast.Function
	errors       []error
	currentValue *ast.Value
}

var (
	_ ast.AstVisitor     = (*typeChecker)(nil)
	_ ast.BuiltinVisitor = (*typeChecker)(nil)
)

func (c *typeChecker) Check(root ast.Node) error {
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

func (c *typeChecker) VisitModule(m *ast.Module) {
	for _, decl := range m.Consts() {
		decl.Expr().Visit(c)

		if v := decl.Expr().ConstValue(); v != nil {
			decl.Update(decl.Expr().Type(), v)
		} else {
			c.errors = append(c.errors, fmt.Errorf("initializer for %q is not constant", decl.Token().Text))
		}

		c.consts[decl.Token().Text] = decl
	}

	for _, decl := range m.Vars() {
		switch decl.TypeToken().Text {
		case "INTEGER":
			decl.Update(ast.TypeInt64)
		case "REAL":
			decl.Update(ast.TypeFloat64)
		case "BOOLEAN":
			decl.Update(ast.TypeBoolean)
		case "CHAR":
			decl.Update(ast.TypeChar)
		case "SET":
			decl.Update(ast.TypeSet)
		default:
			// TODO(daniel): support "ARRAY", "RECORD", "POINTER", and custom types.
			c.errors = append(c.errors, fmt.Errorf("unknown type %q for variable %q",
				decl.TypeToken().Text, decl.Token().Text))
		}

		c.vars[decl.Token().Text] = decl
	}

	m.Block().Visit(c)
}

func (c *typeChecker) VisitStmtSequence(s *ast.StmtSequence) {
	for _, v := range s.Stmts() {
		v.Visit(c)
	}
}

func (c *typeChecker) VisitAssignStmt(s *ast.AssignStmt) {
	lhs := s.Token()
	rhs := s.Expr()

	rhs.Visit(c)

	if t, ok := c.vars[lhs.Text]; ok {
		if rhs.Type() != t.Type() {
			c.errors = append(c.errors, fmt.Errorf("can't assign type %q to variable %q (which is of type %q)",
				rhs.Type(), lhs.Text, t.Type()))
		}
	} else {
		c.errors = append(c.errors, fmt.Errorf("undefined identifier %q", s.Token().Text))
	}
}

func (c *typeChecker) VisitIfStmt(s *ast.IfStmt) {
	s.Expr().Visit(c)
	if s.Expr().Type() != ast.TypeBoolean {
		c.errors = append(c.errors, fmt.Errorf("condition for IF must be boolean, got %v", s.Expr().Type()))
	}

	s.TrueBlock().Visit(c)
	s.FalseBlock().Visit(c)
}

func (c *typeChecker) VisitRepeatStmt(s *ast.RepeatStmt) {
	cond := s.Condition()

	cond.Expr().Visit(c)
	cond.Stmt().Visit(c)

	if cond.Expr().Type() != ast.TypeBoolean {
		c.errors = append(c.errors, fmt.Errorf("condition for REPEAT must be boolean, got %v",
			cond.Expr().Type()))
	}
}

func (c *typeChecker) VisitWhileStmt(s *ast.WhileStmt) {
	for _, cond := range s.Conditions() {
		cond.Expr().Visit(c)
		cond.Stmt().Visit(c)

		if cond.Expr().Type() != ast.TypeBoolean {
			c.errors = append(c.errors, fmt.Errorf("condition for WHILE must be boolean, got %v",
				cond.Expr().Type()))
		}
	}
}

func (c *typeChecker) VisitForStmt(s *ast.ForStmt) {
	iter, ok := c.vars[s.Iter().Text]
	if !ok {
		c.errors = append(c.errors, fmt.Errorf("FOR iterator must be a variable"))
	}

	// TODO(daniel): is this true?
	if iter.Type() != ast.TypeInt64 {
		c.errors = append(c.errors, fmt.Errorf("FOR iterator must be an integer, got %v",
			iter.Type()))
	}

	from := s.From()
	to := s.To()
	by := s.By()

	from.Visit(c)
	to.Visit(c)
	by.Visit(c)

	if from.Type() != iter.Type() {
		c.errors = append(c.errors, fmt.Errorf("FOR iterator FROM must be an integer, got %v",
			from.Type()))
	}

	if to.Type() != iter.Type() {
		c.errors = append(c.errors, fmt.Errorf("FOR iterator TO must be an integer, got %v",
			to.Type()))
	}

	if by.Type() != iter.Type() {
		c.errors = append(c.errors, fmt.Errorf("FOR iterator BY must be an integer, got %v",
			by.Type()))
	}

	if by.ConstValue() == nil {
		c.errors = append(c.errors, fmt.Errorf("FOR iterator BY must be a constant expression"))
	} else if by.ConstValue().Int() == 0 {
		c.errors = append(c.errors, fmt.Errorf("FOR iterator BY must not be zero"))
	}

	if from.ConstValue() != nil && to.ConstValue() != nil && by.ConstValue() != nil {
		fromVal := from.ConstValue().Int()
		toVal := to.ConstValue().Int()
		byVal := by.ConstValue().Int()

		if fromVal > toVal && byVal > 0 {
			c.errors = append(c.errors, fmt.Errorf("FOR iterator BY must be negative if FROM > BY"))
		} else if fromVal < toVal && byVal < 0 {
			c.errors = append(c.errors, fmt.Errorf("FOR iterator BY must be positive if FROM < BY"))
		}
	}

	s.Stmt().Visit(c)
}

func (c *typeChecker) VisitExprStmt(s *ast.ExprStmt) {
	s.Expr().Visit(c)
}

func (c *typeChecker) VisitCallExpr(e *ast.CallExpr) {
	for _, v := range e.Args() {
		v.Visit(c)
	}

	if v, ok := c.builtins[e.Token().Text]; ok {
		t, err := v.Validate(e.Args())
		if err != nil {
			c.errors = append(c.errors, err)

			return
		}

		v.Visit(c, e.Args())

		e.Update(t, c.currentValue, v)
	} else {
		c.errors = append(c.errors, fmt.Errorf("builtin function %q not found", e.Token().Text))
	}
}

func (c *typeChecker) VisitBinaryExpr(e *ast.BinaryExpr) {
	for _, v := range e.Args() {
		v.Visit(c)
	}

	var (
		outType    ast.Type
		constValue *ast.Value
	)

	switch {
	case len(e.Args()) != 2:
		c.errors = append(c.errors, fmt.Errorf("binary expression with %d arguments", len(e.Args())))
	case e.Token().Type == lex.TokenIN:
		outType = ast.TypeBoolean
	case e.Lhs().Type() != e.Rhs().Type():
		outType = ast.TypeVoid

		c.errors = append(c.errors, fmt.Errorf("different types for binary operator %s",
			e.Token().Type))
	case e.Token().Type == lex.TokenSlash:
		if e.Lhs().Type() == ast.TypeSet && e.Rhs().Type() == ast.TypeSet {
			outType = ast.TypeSet
		} else {
			outType = ast.TypeFloat64

			if e.Lhs().Type() != outType {
				c.errors = append(c.errors, fmt.Errorf("unexpected type for binary operator %s",
					e.Token().Type))
			}
		}
	case e.Token().IsRelation():
		outType = ast.TypeBoolean
	case e.Token().Type == lex.TokenAmpersand || e.Token().Type == lex.TokenOR:
		outType = ast.TypeBoolean

		if e.Lhs().Type() != outType {
			c.errors = append(c.errors, fmt.Errorf("unexpected type for binary operator %s",
				e.Token().Type))
		}
	default:
		outType = e.Lhs().Type()
	}

	lhs := e.Lhs().ConstValue()
	rhs := e.Rhs().ConstValue()

	if lhs != nil && rhs != nil {
		constValue = evaluateBinaryExpr(e.Token().Type, outType, lhs, rhs)
	}

	e.Update(outType, constValue)
}

func (c *typeChecker) VisitDesignatorExpr(e *ast.DesignatorExpr) {
	if t, ok := c.consts[e.Token().Text]; ok {
		e.Update(t.Type(), ast.KindConst, t.Expr().ConstValue())
	} else if t, ok := c.vars[e.Token().Text]; ok {
		e.Update(t.Type(), ast.KindVar, nil)
	} else {
		c.errors = append(c.errors, fmt.Errorf("undefined identifier %q", e.Token().Text))
	}
}

func (c *typeChecker) VisitNumberExpr(e *ast.NumberExpr) {
	if v, err := evaluateNumberExpr(e.Token()); err != nil {
		c.errors = append(c.errors, err)
	} else {
		e.Update(v.Type(), v)
	}
}

func (c *typeChecker) VisitStringExpr(e *ast.StringExpr) {
	e.Update(evaluateStringExpr(e.Token()))
}

func (c *typeChecker) VisitCharExpr(e *ast.CharExpr) {
	e.Update(evaluateCharExpr(e.Token()))
}

func (c *typeChecker) VisitBooleanExpr(e *ast.BooleanExpr) {
	if v, err := evaluateBooleanExpr(e.Token()); err != nil {
		c.errors = append(c.errors, err)
	} else {
		e.Update(v)
	}
}

func (c *typeChecker) VisitSetExpr(e *ast.SetExpr) {
	e.Update(evaluateSetExpr(e))
}

func (c *typeChecker) VisitNotExpr(e *ast.NotExpr) {
	e.Expr().Visit(c)

	if v, err := evaluateNotExpr(e.Expr()); err != nil {
		c.errors = append(c.errors, err)
	} else {
		e.Update(v)
	}
}

func (c *typeChecker) VisitBuiltinPrint(*ast.BuiltinPrint, []ast.Expr) {
	c.currentValue = nil
}

func (c *typeChecker) VisitBuiltinABS(f *ast.BuiltinABS, args []ast.Expr) {
	c.currentValue = nil

	if v := args[0].ConstValue(); v != nil {
		switch args[0].Type() {
		case ast.TypeInt64:
			val := v.Int()
			if val < 0 {
				val = -val
			}

			c.currentValue = ast.NewInt(val)
		case ast.TypeFloat64:
			val := v.Real()
			if val < 0 {
				val = -val
			}

			c.currentValue = ast.NewReal(val)
		}
	}
}

func (c *typeChecker) VisitBuiltinODD(f *ast.BuiltinODD, args []ast.Expr) {
	c.currentValue = nil

	if v := args[0].ConstValue(); v != nil {
		c.currentValue = ast.NewBoolean(v.Int()%2 == 1)
	}
}

func (c *typeChecker) VisitBuiltinLSL(*ast.BuiltinLSL, []ast.Expr) {
	c.errors = append(c.errors, fmt.Errorf("BuiltinLSL is not implemented"))
	c.currentValue = nil
}

func (c *typeChecker) VisitBuiltinASR(*ast.BuiltinASR, []ast.Expr) {
	c.errors = append(c.errors, fmt.Errorf("BuiltinASR is not implemented"))
	c.currentValue = nil
}

func (c *typeChecker) VisitBuiltinROR(*ast.BuiltinROR, []ast.Expr) {
	c.errors = append(c.errors, fmt.Errorf("BuiltinROR is not implemented"))
	c.currentValue = nil
}

func (c *typeChecker) VisitBuiltinLEN(*ast.BuiltinLEN, []ast.Expr) {
	c.errors = append(c.errors, fmt.Errorf("BuiltinLEN is not implemented"))
	c.currentValue = nil
}

func (c *typeChecker) VisitBuiltinORD(f *ast.BuiltinORD, args []ast.Expr) {
	c.currentValue = nil

	if v := args[0].ConstValue(); v != nil {
		switch args[0].Type() {
		case ast.TypeBoolean, ast.TypeChar:
			c.currentValue = ast.NewInt(v.Int())
		}
	}
}

func (c *typeChecker) VisitBuiltinCHR(f *ast.BuiltinCHR, args []ast.Expr) {
	c.currentValue = nil

	if v := args[0].ConstValue(); v != nil {
		c.currentValue = ast.NewChar(v.Char())
	}
}

func (c *typeChecker) VisitBuiltinFLOOR(f *ast.BuiltinFLOOR, args []ast.Expr) {
	c.currentValue = nil

	if v := args[0].ConstValue(); v != nil {
		c.currentValue = ast.NewInt(v.Int())
	}
}

func (c *typeChecker) VisitBuiltinFLT(f *ast.BuiltinFLT, args []ast.Expr) {
	c.currentValue = nil

	if v := args[0].ConstValue(); v != nil {
		c.currentValue = ast.NewReal(v.Real())
	}
}

func (c *typeChecker) VisitBuiltinINC(*ast.BuiltinINC, []ast.Expr) {
	c.currentValue = nil
}

func (c *typeChecker) VisitBuiltinDEC(*ast.BuiltinDEC, []ast.Expr) {
	c.currentValue = nil
}

func (c *typeChecker) VisitBuiltinINCL(*ast.BuiltinINCL, []ast.Expr) {
	c.currentValue = nil
}

func (c *typeChecker) VisitBuiltinEXCL(*ast.BuiltinEXCL, []ast.Expr) {
	c.currentValue = nil
}

func (c *typeChecker) VisitBuiltinPACK(*ast.BuiltinPACK, []ast.Expr) {
	c.currentValue = nil
}

func (c *typeChecker) VisitBuiltinUNPK(*ast.BuiltinUNPK, []ast.Expr) {
	c.currentValue = nil
}

func (c *typeChecker) VisitBuiltinNEW(*ast.BuiltinNEW, []ast.Expr) {
	c.currentValue = nil
}

func (c *typeChecker) VisitBuiltinASSERT(*ast.BuiltinASSERT, []ast.Expr) {
	c.currentValue = nil
}