package main

import "fmt"

type Type int

const (
	TypeVoid Type = iota
	TypeInt64
	TypeFloat64
	TypeString
	TypeBoolean
)

func (t Type) String() string {
	switch t {
	case TypeVoid:
		return "void"
	case TypeInt64:
		return "i64"
	case TypeFloat64:
		return "f64"
	case TypeString:
		return "string"
	case TypeBoolean:
		return "boolean"
	default:
		return fmt.Sprintf("undefined=%d", int(t))
	}
}

type Visitor interface {
	VisitModule(*Module)
	VisitExprStmt(*ExprStmt)
	VisitCallExpr(*CallExpr)
	VisitBinaryExpr(*BinaryExpr)
	VisitNumberExpr(*NumberExpr)
	VisitStringExpr(*StringExpr)
	VisitBooleanExpr(*BooleanExpr)
}

type Node interface {
	Token() Token
	Visit(Visitor)
}

type Expr interface {
	Node
	Type() Type
}

type Stmt interface {
	Node
}

type Module struct {
	stmts []Stmt
}

var _ Node = (*Module)(nil)

func (n *Module) Token() Token {
	if len(n.stmts) == 0 {
		return Token{}
	}

	return n.stmts[0].Token()
}

func (n *Module) Visit(v Visitor) {
	v.VisitModule(n)
}

type NumberExpr struct {
	token Token
	typ   Type
}

var _ Expr = (*NumberExpr)(nil)

func (n *NumberExpr) Token() Token {
	return n.token
}

func (n *NumberExpr) Type() Type {
	return n.typ
}

func (n *NumberExpr) Visit(v Visitor) {
	v.VisitNumberExpr(n)
}

type StringExpr struct {
	token Token
}

var _ Expr = (*StringExpr)(nil)

func (n *StringExpr) Token() Token {
	return n.token
}

func (n *StringExpr) Type() Type {
	return TypeString
}

func (n *StringExpr) Visit(v Visitor) {
	v.VisitStringExpr(n)
}

type BooleanExpr struct {
	token Token
}

var _ Expr = (*BooleanExpr)(nil)

func (n *BooleanExpr) Token() Token {
	return n.token
}

func (n *BooleanExpr) Type() Type {
	return TypeBoolean
}

func (n *BooleanExpr) Visit(v Visitor) {
	v.VisitBooleanExpr(n)
}

type BinaryExpr struct {
	token Token
	args  []Expr
	typ   Type
}

var _ Expr = (*BinaryExpr)(nil)

func (n *BinaryExpr) Token() Token {
	return n.token
}

func (n *BinaryExpr) Type() Type {
	return n.typ
}

func (n *BinaryExpr) Visit(v Visitor) {
	v.VisitBinaryExpr(n)
}

type CallExpr struct {
	token Token
	typ   Type
	args  []Expr
}

var _ Expr = (*CallExpr)(nil)

func (n *CallExpr) Token() Token {
	return n.token
}

func (n *CallExpr) Type() Type {
	return n.typ
}

func (n *CallExpr) Visit(v Visitor) {
	v.VisitCallExpr(n)
}

type ExprStmt struct {
	expr Expr
}

var _ Stmt = (*ExprStmt)(nil)

func (n *ExprStmt) Token() Token {
	return n.expr.Token()
}

func (n *ExprStmt) Visit(v Visitor) {
	v.VisitExprStmt(n)
}
