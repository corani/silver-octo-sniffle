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

type Value struct {
	typ     Type
	integer int
	real    float64
	Bool    bool
	String  string
}

func (v Value) Type() Type {
	return v.typ
}

func (v Value) Int() int {
	switch v.typ {
	case TypeInt64:
		return v.integer
	case TypeFloat64:
		return int(v.real)
	default:
		return 0
	}
}

func (v Value) Real() float64 {
	switch v.typ {
	case TypeInt64:
		return float64(v.integer)
	case TypeFloat64:
		return v.real
	default:
		return 0
	}
}

type Visitor interface {
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
	VisitBooleanExpr(*BooleanExpr)
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
	Visit(Visitor)
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

func (n *Module) Visit(v Visitor) {
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

func (n *StmtSequence) Visit(v Visitor) {
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

func (n *IfStmt) Visit(v Visitor) {
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

func (n *RepeatStmt) Visit(v Visitor) {
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

func (n *WhileStmt) Visit(v Visitor) {
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

func (n *ForStmt) Visit(v Visitor) {
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

func (n *AssignStmt) Visit(v Visitor) {
	v.VisitAssignStmt(n)
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

func (n *DesignatorExpr) ConstValue() *Value {
	return n.constValue
}

func (n *DesignatorExpr) Visit(v Visitor) {
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

func (n *NumberExpr) Visit(v Visitor) {
	v.VisitNumberExpr(n)
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

func (n *StringExpr) Visit(v Visitor) {
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

func (n *BooleanExpr) Visit(v Visitor) {
	v.VisitBooleanExpr(n)
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

func (n *NotExpr) Visit(v Visitor) {
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

func (n *BinaryExpr) Visit(v Visitor) {
	v.VisitBinaryExpr(n)
}

type CallExpr struct {
	designator *DesignatorExpr
	typ        Type
	args       []Expr
	constValue *Value
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

func (n *CallExpr) Visit(v Visitor) {
	v.VisitCallExpr(n)
}
