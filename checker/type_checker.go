package checker

import (
	"errors"

	"github.com/corani/silver-octo-sniffle/ast"
	"github.com/corani/silver-octo-sniffle/reporter"
	"github.com/corani/silver-octo-sniffle/token"
)

var ErrChecking = errors.New("type check error")

func TypeCheck(out *reporter.Reporter, root ast.Node) {
	checker := &typeChecker{
		out: out,
		scope: &scope{
			parent:  nil,
			symbols: ast.GetBuiltinSymbols(),
		},
		modules: map[string]*module{
			"": {
				builtins: ast.GetBuiltinFunctions(),
			},
			"C": {
				builtins: ast.GetCFunctions(),
			},
			"Texts": {
				builtins: ast.GetTextsFunctions(),
			},
		},
		currentValue: nil,
	}

	checker.Check(root)
}

type module struct {
	builtins map[string]ast.Function
}

type scope struct {
	parent  *scope
	symbols map[string]ast.Symbol
}

type typeChecker struct {
	out          *reporter.Reporter
	scope        *scope
	modules      map[string]*module
	currentValue *ast.Value
}

var (
	_ ast.AstVisitor     = (*typeChecker)(nil)
	_ ast.BuiltinVisitor = (*typeChecker)(nil)
)

func (c *typeChecker) Check(root ast.Node) {
	root.Visit(c)
}

func (c *typeChecker) enterScope() {
	c.scope = &scope{
		parent:  c.scope,
		symbols: make(map[string]ast.Symbol),
	}
}

func (c *typeChecker) leaveScope() {
	c.scope = c.scope.parent
}

func (c *typeChecker) findSymbol(name string) (ast.Symbol, bool) {
	scope := c.scope

	for scope != nil {
		if v, ok := scope.symbols[name]; ok {
			return v, true
		}

		scope = scope.parent
	}

	return nil, false
}

func (c *typeChecker) defineSymbol(name string, sym ast.Symbol) {
	c.scope.symbols[name] = sym
}

func (c *typeChecker) findFunction(q *ast.QualIdent) (ast.Function, bool) {
	mod := c.modules[""]

	if q.Qualifier() != nil {
		name := q.Qualifier().Text
		if v, ok := c.modules[name]; ok {
			mod = v
		} else {
			c.out.Errorf(q.Token(), "module %v not found", name)

			return nil, false
		}
	}

	name := q.Ident().Text

	if v, ok := mod.builtins[name]; ok {
		return v, true
	}

	c.out.Errorf(q.Token(), "identifier %v not found", name)

	return nil, false
}

// ----- statements -------------------------------------------------------------------------------

func (c *typeChecker) VisitInvalidStmt(s *ast.InvalidStmt) {
	// nothing
}

func (c *typeChecker) VisitModule(m *ast.Module) {
	for _, decl := range m.Decls() {
		decl.Visit(c)
	}

	m.Block().Visit(c)
}

func (c *typeChecker) VisitStmtSequence(s *ast.StmtSequence) {
	for _, v := range s.Stmts() {
		v.Visit(c)
	}
}

func (c *typeChecker) VisitAssignStmt(s *ast.AssignStmt) {
	lhs := s.Designator()
	rhs := s.Expr()

	lhs.Visit(c)
	rhs.Visit(c)

	if lhs.Kind() == ast.KindVar {
		if rhs.Type() != lhs.Type() {
			c.out.Errorf(s.Token(), "can't assign type %q to variable %q (which is of type %q)",
				rhs.Type(), lhs.Token().Text, lhs.Type())
		}
	} else {
		c.out.Errorf(s.Token(), "undefined identifier %q", s.Token().Text)
		c.out.Infof(s.Token(), "expected a variable, got a %s", lhs.Kind())
	}
}

func (c *typeChecker) VisitIfStmt(s *ast.IfStmt) {
	s.Expr().Visit(c)
	if s.Expr().Type() != ast.TypeBoolean {
		c.out.Errorf(s.Token(), "condition for IF must be boolean, got %v",
			s.Expr().Type())
	}

	s.TrueBlock().Visit(c)
	s.FalseBlock().Visit(c)
}

func (c *typeChecker) VisitRepeatStmt(s *ast.RepeatStmt) {
	cond := s.Condition()

	cond.Expr().Visit(c)
	cond.Stmt().Visit(c)

	if cond.Expr().Type() != ast.TypeBoolean {
		c.out.Errorf(s.Token(), "condition for REPEAT must be boolean, got %v",
			cond.Expr().Type())
	}
}

func (c *typeChecker) VisitWhileStmt(s *ast.WhileStmt) {
	for _, cond := range s.Conditions() {
		cond.Expr().Visit(c)
		cond.Stmt().Visit(c)

		if cond.Expr().Type() != ast.TypeBoolean {
			c.out.Errorf(s.Token(), "condition for WHILE must be boolean, got %v",
				cond.Expr().Type())
		}
	}
}

func (c *typeChecker) VisitForStmt(s *ast.ForStmt) {
	sym, ok := c.findSymbol(s.Iter().Text)
	if !ok {
		c.out.Errorf(s.Iter(), "undefined identifier %q", s.Iter().Text)
	}

	iter, ok := sym.(*ast.VarDecl)
	if !ok {
		c.out.Errorf(s.Token(), "FOR iterator must be a variable, got %s", sym.Kind())
	} else {
		// TODO(daniel): is this true?
		if iter.Type() != ast.TypeInt64 {
			c.out.Errorf(iter.Token(), "FOR iterator must be an integer, got %v",
				iter.Type())
		}
	}

	from := s.From()
	to := s.To()
	by := s.By()

	from.Visit(c)
	to.Visit(c)
	by.Visit(c)

	if iter != nil {
		if from.Type() != iter.Type() {
			c.out.Errorf(from.Token(), "FOR iterator FROM must be an integer, got %v",
				from.Type())
		}

		if to.Type() != iter.Type() {
			c.out.Errorf(to.Token(), "FOR iterator TO must be an integer, got %v",
				to.Type())
		}

		if by.Type() != iter.Type() {
			c.out.Errorf(by.Token(), "FOR iterator BY must be an integer, got %v",
				by.Type())
		}
	}

	if by.ConstValue() == nil {
		c.out.Errorf(by.Token(), "FOR iterator BY must be a constant expression")
	} else if by.ConstValue().Int() == 0 {
		c.out.Errorf(by.Token(), "FOR iterator BY must not be zero")
	}

	if from.ConstValue() != nil && to.ConstValue() != nil && by.ConstValue() != nil {
		fromVal := from.ConstValue().Int()
		toVal := to.ConstValue().Int()
		byVal := by.ConstValue().Int()

		if fromVal > toVal && byVal > 0 {
			c.out.Errorf(by.Token(), "FOR iterator BY must be negative if FROM > BY")
		} else if fromVal < toVal && byVal < 0 {
			c.out.Errorf(by.Token(), "FOR iterator BY must be positive if FROM < BY")
		}
	}

	s.Stmt().Visit(c)
}

func (c *typeChecker) VisitReturnStmt(s *ast.ReturnStmt) {
	s.Expr().Visit(c)
}

func (c *typeChecker) VisitExprStmt(s *ast.ExprStmt) {
	s.Expr().Visit(c)
}

// ----- expressions ------------------------------------------------------------------------------

func (c *typeChecker) VisitInvalidExpr(s *ast.InvalidExpr) {
	// nothing
}

func (c *typeChecker) VisitCallExpr(e *ast.CallExpr) {
	lhs := e.Designator()

	lhs.Visit(c)

	for _, v := range e.Args() {
		v.Visit(c)
	}

	if v, ok := c.findFunction(lhs.QualIdent()); ok {
		t, err := v.Validate(e.Args())
		if err != nil {
			c.out.Errorf(e.Token(), "%v", err)

			return
		}

		v.Visit(c, e.Args())

		e.Update(t, c.currentValue, v)
	} else {
		c.out.Errorf(lhs.Token(), "builtin function %q not found",
			e.Token().Text)
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
		c.out.Errorf(e.Token(), "binary expression with %d arguments",
			len(e.Args()))
	case e.Token().Type == token.TokenIN:
		outType = ast.TypeBoolean
	case e.Lhs().Type() != e.Rhs().Type():
		outType = ast.TypeVoid

		c.out.Errorf(e.Token(), "different types for binary operator %s",
			e.Token().Type)
	case e.Token().Type == token.TokenSlash:
		if e.Lhs().Type() == ast.TypeSet && e.Rhs().Type() == ast.TypeSet {
			outType = ast.TypeSet
		} else {
			outType = ast.TypeFloat64

			if e.Lhs().Type() != outType {
				c.out.Errorf(e.Token(), "unexpected type for binary operator %s",
					e.Token().Type)
			}
		}
	case e.Token().IsRelation():
		outType = ast.TypeBoolean
	case e.Token().Type == token.TokenAmpersand || e.Token().Type == token.TokenOR:
		outType = ast.TypeBoolean

		if e.Lhs().Type() != outType {
			c.out.Errorf(e.Token(), "unexpected type for binary operator %s",
				e.Token().Type)
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
	// TODO(daniel): support qualidents.
	if sym, ok := c.findSymbol(e.Token().Text); ok {
		switch t := sym.(type) {
		case *ast.VarDecl:
			decl := t.TypeDecl()

			// TODO(daniel): process selectors.
			for _, selector := range e.Selectors() {
				switch selector.(type) {
				case *ast.DerefSelector:
					decl = decl.(*ast.TypePointerDecl).To()
				}
			}

			e.Update(decl.Type(), ast.KindVar, nil)
		case *ast.ConstDecl:
			e.Update(t.Type(), ast.KindConst, t.Expr().ConstValue())
			// TODO(daniel): process selectors.
		default:
			c.out.Errorf(e.Token(), "unsupported kind %q for identifier %q",
				sym.Kind(), e.Token().Text)
		}
	} else if fn, ok := c.findFunction(e.QualIdent()); ok {
		e.Update(fn.Type(), ast.KindProc, nil)
	} else {
		c.out.Errorf(e.Token(), "undefined identifier %q",
			e.Token().Text)
	}
}

func (c *typeChecker) VisitNotExpr(e *ast.NotExpr) {
	e.Expr().Visit(c)

	if v, err := evaluateNotExpr(e.Expr()); err != nil {
		c.out.Errorf(e.Token(), "%v", err)
	} else {
		e.Update(v)
	}
}

// ----- literals ---------------------------------------------------------------------------------

func (c *typeChecker) VisitNumberLit(e *ast.NumberLit) {
	if v, err := evaluateNumberLit(e); err != nil {
		c.out.Errorf(e.Token(), "%v", err)
	} else {
		e.Update(v.Type(), v)
	}
}

func (c *typeChecker) VisitStringLit(e *ast.StringLit) {
	e.Update(evaluateStringLit(e))
}

func (c *typeChecker) VisitCharLit(e *ast.CharLit) {
	e.Update(evaluateCharLit(e))
}

func (c *typeChecker) VisitBooleanLit(e *ast.BooleanLit) {
	if v, err := evaluateBooleanLit(e); err != nil {
		c.out.Errorf(e.Token(), "%v", err)
	} else {
		e.Update(v)
	}
}

func (c *typeChecker) VisitSetLit(e *ast.SetLit) {
	e.Update(evaluateSetLit(e))
}

// ----- declarations -----------------------------------------------------------------------------

func (c *typeChecker) VisitConstDecl(decl *ast.ConstDecl) {
	decl.Expr().Visit(c)

	if v := decl.Expr().ConstValue(); v != nil {
		decl.Update(decl.Expr().Type(), v)
	} else {
		c.out.Errorf(decl.Token(), "initializer for %q is not constant", decl.Token().Text)
	}

	c.defineSymbol(decl.Token().Text, decl)
}

func (c *typeChecker) VisitTypeBaseDecl(decl *ast.TypeBaseDecl) {
	if decl.TypeRef() != nil {
		typeRef := decl.TypeRef()
		typeRef.Visit(c)

		if sym, ok := c.findSymbol(typeRef.Token().Text); ok {
			decl.Update(typeRef.Token(), sym.Type())
		} else {
			decl.Update(typeRef.Token(), ast.TypeVoid)
			c.out.Errorf(typeRef.Token(), "unknown type %q", typeRef.Token().Text)
		}
	}
}

func (c *typeChecker) VisitTypePointerDecl(decl *ast.TypePointerDecl) {
	if decl.TypeRef() != nil {
		// NOTE(daniel): resolve named type
		decl.TypeRef().Visit(c)
		decl.Update(decl.TypeRef().TypeDecl())
	} else {
		// NOTE(daniel): anonymous type.
		decl.To().Visit(c)
	}

	if decl.Token().Type == token.TokenIdent {
		c.defineSymbol(decl.Token().Text, decl)
	}
}

func (c *typeChecker) VisitVarDecl(decl *ast.VarDecl) {
	if decl.TypeRef() != nil {
		// NOTE(daniel): resolve named type.
		typeRef := decl.TypeRef()
		typeRef.Visit(c)

		if sym, ok := c.findSymbol(typeRef.Token().Text); ok {
			decl.Update(sym.(ast.TypeDecl), sym.Type())
		} else {
			c.out.Errorf(decl.TypeRef().Token(), "unknown type %q",
				decl.TypeRef().Token().Text)

			return
		}
	} else {
		// NOTE(daniel): anonymous type.
		decl.TypeDecl().Visit(c)
		decl.Update(decl.TypeDecl(), decl.TypeDecl().Type())
	}

	if decl.TypeDecl().Type() != ast.TypeVoid {
		c.defineSymbol(decl.Token().Text, decl)
	} else {
		c.out.Errorf(decl.Token(), "unknown type %q for variable %q",
			decl.TypeDecl().Token().Text, decl.Token().Text)
	}
}

func (c *typeChecker) VisitProcDecl(decl *ast.ProcDecl) {
	c.enterScope()
	decl.Stmts().Visit(c)
	c.leaveScope()
}

func (c *typeChecker) VisitTypeRef(ref *ast.TypeRef) {
	// NOTE(daniel): solidify type reference.
	if ref.TypeDecl() == nil {
		if sym, ok := c.findSymbol(ref.Token().Text); ok {
			ref.Update(sym.(ast.TypeDecl))
		}
	}
}

// ----- builtin functions ------------------------------------------------------------------------

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

func (c *typeChecker) VisitBuiltinLSL(f *ast.BuiltinLSL, args []ast.Expr) {
	c.out.Errorf(args[0].Token(), "BuiltinLSL is not implemented")
	c.currentValue = nil
}

func (c *typeChecker) VisitBuiltinASR(f *ast.BuiltinASR, args []ast.Expr) {
	c.out.Errorf(args[0].Token(), "BuiltinASR is not implemented")
	c.currentValue = nil
}

func (c *typeChecker) VisitBuiltinROR(f *ast.BuiltinROR, args []ast.Expr) {
	c.out.Errorf(args[0].Token(), "BuiltinROR is not implemented")
	c.currentValue = nil
}

func (c *typeChecker) VisitBuiltinLEN(f *ast.BuiltinLEN, args []ast.Expr) {
	c.out.Errorf(args[0].Token(), "BuiltinLEN is not implemented")
	c.currentValue = nil
}

func (c *typeChecker) VisitBuiltinORD(f *ast.BuiltinORD, args []ast.Expr) {
	c.currentValue = nil

	if v := args[0].ConstValue(); v != nil {
		switch args[0].Type() {
		case ast.TypeBoolean, ast.TypeChar, ast.TypeSet:
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

func (c *typeChecker) VisitBuiltinDELETE(*ast.BuiltinDELETE, []ast.Expr) {
	c.currentValue = nil
}

func (c *typeChecker) VisitBuiltinASSERT(*ast.BuiltinASSERT, []ast.Expr) {
	c.currentValue = nil
}

func (c *typeChecker) VisitCPrint(*ast.CPrint, []ast.Expr) {
	c.currentValue = nil
}

func (c *typeChecker) VisitTextsWriteInt(*ast.TextsWriteInt, []ast.Expr) {
	c.currentValue = nil
}

func (c *typeChecker) VisitTextsWriteReal(*ast.TextsWriteReal, []ast.Expr) {
	c.currentValue = nil
}

func (c *typeChecker) VisitTextsWriteString(*ast.TextsWriteString, []ast.Expr) {
	c.currentValue = nil
}

func (c *typeChecker) VisitTextsWriteChar(*ast.TextsWriteChar, []ast.Expr) {
	c.currentValue = nil
}

func (c *typeChecker) VisitTextsWriteLn(*ast.TextsWriteLn, []ast.Expr) {
	c.currentValue = nil
}
