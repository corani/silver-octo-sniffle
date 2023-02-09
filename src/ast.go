package main

import "fmt"

type Kind int

const (
	KindUndefined Kind = iota
	KindVar
	KindConst
)

func (k Kind) String() string {
	switch k {
	case KindUndefined:
		return "undefined"
	case KindVar:
		return "variable"
	case KindConst:
		return "constant"
	default:
		return fmt.Sprintf("undefined=%d", int(k))
	}
}

type AstVisitor interface {
	VisitModule(*Module)
	VisitStmtSequence(*StmtSequence)
	VisitIfStmt(*IfStmt)
	VisitRepeatStmt(*RepeatStmt)
	VisitWhileStmt(*WhileStmt)
	VisitForStmt(*ForStmt)
	VisitAssignStmt(*AssignStmt)
	VisitExprStmt(*ExprStmt)
	VisitCallExpr(*CallExpr)
	VisitBinaryExpr(*BinaryExpr)
	VisitDesignatorExpr(*DesignatorExpr)
	VisitNumberExpr(*NumberExpr)
	VisitStringExpr(*StringExpr)
	VisitCharExpr(*CharExpr)
	VisitBooleanExpr(*BooleanExpr)
	VisitSetExpr(*SetExpr)
	VisitNotExpr(*NotExpr)
}

type ConstDecl struct {
	token Token
	expr  Expr
	typ   Type
	value Value
}

type TypeDecl struct {
	token    Token
	typToken Token
	typ      Type
}

type VarDecl struct {
	token    Token
	typToken Token
	typ      Type
}

type Node interface {
	Token() Token
	Visit(AstVisitor)
}

type Expr interface {
	Node
	Type() Type
	ConstValue() *Value
}

type Stmt interface {
	Node
}

type Module struct {
	token  Token
	name   string
	stmts  Stmt
	consts []ConstDecl
	types  []TypeDecl
	vars   []VarDecl
}

var _ Node = (*Module)(nil)

func (n *Module) Token() Token {
	return n.token
}

func (n *Module) Visit(v AstVisitor) {
	v.VisitModule(n)
}

type StmtSequence struct {
	stmts []Stmt
}

var _ Stmt = (*StmtSequence)(nil)

func (n *StmtSequence) Token() Token {
	if len(n.stmts) == 0 {
		return Token{}
	}

	return n.stmts[0].Token()
}

func (n *StmtSequence) Visit(v AstVisitor) {
	v.VisitStmtSequence(n)
}

type IfStmt struct {
	token      Token
	expr       Expr
	trueBlock  Stmt
	falseBlock Stmt
}

var _ Stmt = (*IfStmt)(nil)

func (n *IfStmt) Token() Token {
	return n.token
}

func (n *IfStmt) Visit(v AstVisitor) {
	v.VisitIfStmt(n)
}

type Condition struct {
	expr Expr
	stmt Stmt
}

type RepeatStmt struct {
	token Token
	cond  Condition
}

var _ Stmt = (*RepeatStmt)(nil)

func (n *RepeatStmt) Token() Token {
	return n.token
}

func (n *RepeatStmt) Visit(v AstVisitor) {
	v.VisitRepeatStmt(n)
}

type WhileStmt struct {
	token Token
	conds []Condition
}

var _ Stmt = (*WhileStmt)(nil)

func (n *WhileStmt) Token() Token {
	return n.token
}

func (n *WhileStmt) Visit(v AstVisitor) {
	v.VisitWhileStmt(n)
}

type ForStmt struct {
	token Token
	iter  Token
	from  Expr
	to    Expr
	by    Expr
	stmt  Stmt
}

var _ Stmt = (*ForStmt)(nil)

func (n *ForStmt) Token() Token {
	return n.token
}

func (n *ForStmt) Visit(v AstVisitor) {
	v.VisitForStmt(n)
}

type AssignStmt struct {
	token Token
	expr  Expr
}

var _ Stmt = (*AssignStmt)(nil)

func (n *AssignStmt) Token() Token {
	return n.token
}

func (n *AssignStmt) Visit(v AstVisitor) {
	v.VisitAssignStmt(n)
}

type ExprStmt struct {
	expr Expr
}

var _ Stmt = (*ExprStmt)(nil)

func (n *ExprStmt) Token() Token {
	return n.expr.Token()
}

func (n *ExprStmt) Visit(v AstVisitor) {
	v.VisitExprStmt(n)
}

type DesignatorExpr struct {
	token      Token
	typ        Type
	constValue *Value
	kind       Kind
}

var _ Expr = (*DesignatorExpr)(nil)

func (n *DesignatorExpr) Token() Token {
	return n.token
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

func (n *DesignatorExpr) Visit(v AstVisitor) {
	v.VisitDesignatorExpr(n)
}

type NumberExpr struct {
	token      Token
	typ        Type
	constValue *Value
}

var _ Expr = (*NumberExpr)(nil)

func (n *NumberExpr) Token() Token {
	return n.token
}

func (n *NumberExpr) Type() Type {
	return n.typ
}

func (n *NumberExpr) ConstValue() *Value {
	return n.constValue
}

func (n *NumberExpr) Visit(v AstVisitor) {
	v.VisitNumberExpr(n)
}

type CharExpr struct {
	token      Token
	constValue *Value
}

var _ Expr = (*CharExpr)(nil)

func (n *CharExpr) Token() Token {
	return n.token
}

func (n *CharExpr) Type() Type {
	return TypeChar
}

func (n *CharExpr) ConstValue() *Value {
	return n.constValue
}

func (n *CharExpr) Visit(v AstVisitor) {
	v.VisitCharExpr(n)
}

type StringExpr struct {
	token      Token
	constValue *Value
}

var _ Expr = (*StringExpr)(nil)

func (n *StringExpr) Token() Token {
	return n.token
}

func (n *StringExpr) Type() Type {
	return TypeString
}

func (n *StringExpr) ConstValue() *Value {
	return n.constValue
}

func (n *StringExpr) Visit(v AstVisitor) {
	v.VisitStringExpr(n)
}

type BooleanExpr struct {
	token      Token
	constValue *Value
}

var _ Expr = (*BooleanExpr)(nil)

func (n *BooleanExpr) Token() Token {
	return n.token
}

func (n *BooleanExpr) Type() Type {
	return TypeBoolean
}

func (n *BooleanExpr) ConstValue() *Value {
	return n.constValue
}

func (n *BooleanExpr) Visit(v AstVisitor) {
	v.VisitBooleanExpr(n)
}

type SetExpr struct {
	token      Token
	bits       []byte
	constValue *Value
}

func (n *SetExpr) Token() Token {
	return n.token
}

func (n *SetExpr) Type() Type {
	return TypeSet
}

func (n *SetExpr) ConstValue() *Value {
	return n.constValue
}

func (n *SetExpr) Visit(v AstVisitor) {
	v.VisitSetExpr(n)
}

type NotExpr struct {
	token      Token
	expr       Expr
	constValue *Value
}

var _ Expr = (*NotExpr)(nil)

func (n *NotExpr) Token() Token {
	return n.token
}

func (n *NotExpr) Type() Type {
	return TypeBoolean
}

func (n *NotExpr) ConstValue() *Value {
	return n.constValue
}

func (n *NotExpr) Visit(v AstVisitor) {
	v.VisitNotExpr(n)
}

type BinaryExpr struct {
	token      Token
	args       []Expr
	typ        Type
	constValue *Value
}

var _ Expr = (*BinaryExpr)(nil)

func (n *BinaryExpr) Token() Token {
	return n.token
}

func (n *BinaryExpr) Type() Type {
	return n.typ
}

func (n *BinaryExpr) ConstValue() *Value {
	return n.constValue
}

func (n *BinaryExpr) Visit(v AstVisitor) {
	v.VisitBinaryExpr(n)
}

type CallExpr struct {
	designator *DesignatorExpr
	typ        Type
	args       []Expr
	constValue *Value
	builtin    Function
}

var _ Expr = (*CallExpr)(nil)

func (n *CallExpr) Token() Token {
	return n.designator.Token()
}

func (n *CallExpr) Type() Type {
	return n.typ
}

func (n *CallExpr) ConstValue() *Value {
	return n.constValue
}

func (n *CallExpr) Visit(v AstVisitor) {
	v.VisitCallExpr(n)
}
