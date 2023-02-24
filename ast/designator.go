package ast

import "github.com/corani/silver-octo-sniffle/token"

// ----- DesignatorExpr -------------------------------------------------------

type QualIdent struct {
	qual  *token.Token
	ident token.Token
}

func NewQualIdent(t1 token.Token, t2 ...token.Token) *QualIdent {
	switch len(t2) {
	case 0:
		return &QualIdent{
			qual:  nil,
			ident: t1,
		}
	case 1:
		return &QualIdent{
			qual:  &t1,
			ident: t2[0],
		}
	default:
		panic("too many arguments")
	}
}

func (n *QualIdent) Qualifier() *token.Token {
	return n.qual
}

func (n *QualIdent) Ident() token.Token {
	return n.ident
}

func (n *QualIdent) Token() token.Token {
	if n.qual != nil {
		return *n.qual
	}

	return n.ident
}

// ----- DesignatorExpr -------------------------------------------------------

type DesignatorExpr struct {
	qualIdent  *QualIdent
	typ        Type
	constValue *Value
	kind       Kind
	assignment bool
	selectors  []Selector
}

var _ Expr = (*DesignatorExpr)(nil)

func NewDesignatorExpr(q *QualIdent) *DesignatorExpr {
	return &DesignatorExpr{
		qualIdent: q,
	}
}

func (n *DesignatorExpr) Token() token.Token {
	return n.qualIdent.Token()
}

func (n *DesignatorExpr) Type() Type {
	return n.typ
}

func (n *DesignatorExpr) Kind() Kind {
	return n.kind
}

func (n *DesignatorExpr) ConstValue() *Value {
	return n.constValue
}

func (n *DesignatorExpr) IsAssignment() bool {
	// TODO(daniel): do we need to know this during generation?
	return n.assignment
}

func (n *DesignatorExpr) QualIdent() *QualIdent {
	return n.qualIdent
}

func (n *DesignatorExpr) AddSelector(s Selector) {
	n.selectors = append(n.selectors, s)
}

func (n *DesignatorExpr) Selectors() []Selector {
	return n.selectors
}

func (n *DesignatorExpr) Visit(v AstVisitor) {
	v.VisitDesignatorExpr(n)
}

func (n *DesignatorExpr) Update(t Type, k Kind, c *Value) {
	n.typ = t
	n.kind = k
	n.constValue = c
}

// ===== Selector =============================================================

type Selector interface {
	Token() token.Token
}

// ----- DerefSelector --------------------------------------------------------

func NewDerefSelector(token token.Token) *DerefSelector {
	return &DerefSelector{
		token: token,
	}
}

type DerefSelector struct {
	token token.Token
}

var _ Selector = (*DerefSelector)(nil)

func (s *DerefSelector) Token() token.Token {
	return s.token
}

// ----- DotSelector ----------------------------------------------------------

func NewDotSelector(token, ident token.Token) *DotSelector {
	return &DotSelector{
		token: token,
		ident: ident,
	}
}

type DotSelector struct {
	token token.Token
	ident token.Token
}

var _ Selector = (*DotSelector)(nil)

func (s *DotSelector) Token() token.Token {
	return s.token
}

func (s *DotSelector) Ident() token.Token {
	return s.ident
}

// ----- BracketSelector ------------------------------------------------------

func NewBracketSelector(token token.Token, exprs []Expr) *BracketSelector {
	return &BracketSelector{
		token: token,
		exprs: exprs,
	}
}

type BracketSelector struct {
	token token.Token
	exprs []Expr
}

var _ Selector = (*BracketSelector)(nil)

func (s *BracketSelector) Token() token.Token {
	return s.token
}

func (s *BracketSelector) Exprs() []Expr {
	return s.exprs
}
