package main

import "fmt"

type Type int

const (
	TypeVoid Type = iota
	TypeInt64
	TypeFloat64
)

func (t Type) String() string {
	switch t {
	case TypeVoid:
		return "string"
	case TypeInt64:
		return "i64"
	case TypeFloat64:
		return "f64"
	default:
		return fmt.Sprintf("undefined=%d", int(t))
	}
}

type Visitor interface {
	VisitModule(*Module)
	VisitPrintStmt(*PrintStmt)
	VisitBinaryExpr(*BinaryExpr)
	VisitNumberExpr(*NumberExpr)
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

type PrintStmt struct {
	token Token
	args  []Expr
}

var _ Stmt = (*PrintStmt)(nil)

func (n *PrintStmt) Token() Token {
	return n.token
}

func (n *PrintStmt) Visit(v Visitor) {
	v.VisitPrintStmt(n)
}
