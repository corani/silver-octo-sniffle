package ast

import (
	"github.com/corani/silver-octo-sniffle/lex"
)

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

type Node interface {
	Token() lex.Token
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

// ----- Condition ------------------------------------------------------------

type Condition struct {
	expr Expr
	stmt Stmt
}

func (c Condition) Expr() Expr {
	return c.expr
}

func (c Condition) Stmt() Stmt {
	return c.stmt
}

func NewCondition(e Expr, s Stmt) Condition {
	return Condition{
		expr: e,
		stmt: s,
	}
}

// ----- Module ---------------------------------------------------------------

type Module struct {
	token  lex.Token
	name   string
	block  Stmt
	consts []*ConstDecl
	types  []*TypeDecl
	vars   []*VarDecl
}

var _ Node = (*Module)(nil)

// TODO(daniel): maybe instead of passing all decls in the constructor, add some `NewXxx` methods?
func NewModule(token lex.Token, n string, s Stmt, c []*ConstDecl, t []*TypeDecl, v []*VarDecl) *Module {
	return &Module{
		token:  token,
		name:   n,
		block:  s,
		consts: c,
		types:  t,
		vars:   v,
	}
}

func (n *Module) Token() lex.Token {
	return n.token
}

func (n *Module) Block() Stmt {
	return n.block
}

func (n *Module) Consts() []*ConstDecl {
	return n.consts
}

func (n *Module) Types() []*TypeDecl {
	return n.types
}

func (n *Module) Vars() []*VarDecl {
	return n.vars
}

func (n *Module) Visit(v AstVisitor) {
	v.VisitModule(n)
}

// ----- StmtSequence ---------------------------------------------------------

type StmtSequence struct {
	stmts []Stmt
}

var _ Stmt = (*StmtSequence)(nil)

func NewStmtSequence(s []Stmt) *StmtSequence {
	return &StmtSequence{
		stmts: s,
	}
}

func (n *StmtSequence) Token() lex.Token {
	if len(n.stmts) == 0 {
		return lex.Token{}
	}

	return n.stmts[0].Token()
}

func (n *StmtSequence) Stmts() []Stmt {
	return n.stmts
}

func (n *StmtSequence) Visit(v AstVisitor) {
	v.VisitStmtSequence(n)
}

// ----- IfStmt ---------------------------------------------------------------

type IfStmt struct {
	token      lex.Token
	expr       Expr
	trueBlock  Stmt
	falseBlock Stmt
}

var _ Stmt = (*IfStmt)(nil)

func NewIfStmt(t lex.Token, e Expr, tb, fb Stmt) *IfStmt {
	return &IfStmt{
		token:      t,
		expr:       e,
		trueBlock:  tb,
		falseBlock: fb,
	}
}

func (n *IfStmt) Token() lex.Token {
	return n.token
}

func (n *IfStmt) Expr() Expr {
	return n.expr
}

func (n *IfStmt) TrueBlock() Stmt {
	return n.trueBlock
}

func (n *IfStmt) FalseBlock() Stmt {
	return n.falseBlock
}

func (n *IfStmt) Visit(v AstVisitor) {
	v.VisitIfStmt(n)
}

// ----- RepeatStmt -----------------------------------------------------------

type RepeatStmt struct {
	token lex.Token
	cond  Condition
}

var _ Stmt = (*RepeatStmt)(nil)

func NewRepeatStmt(t lex.Token, c Condition) *RepeatStmt {
	return &RepeatStmt{
		token: t,
		cond:  c,
	}
}

func (n *RepeatStmt) Token() lex.Token {
	return n.token
}

func (n *RepeatStmt) Condition() Condition {
	return n.cond
}

func (n *RepeatStmt) Visit(v AstVisitor) {
	v.VisitRepeatStmt(n)
}

// ----- WhileStmt ------------------------------------------------------------

type WhileStmt struct {
	token lex.Token
	conds []Condition
}

var _ Stmt = (*WhileStmt)(nil)

func NewWhileStmt(t lex.Token, c []Condition) *WhileStmt {
	return &WhileStmt{
		token: t,
		conds: c,
	}
}

func (n *WhileStmt) Token() lex.Token {
	return n.token
}

func (n *WhileStmt) Conditions() []Condition {
	return n.conds
}

func (n *WhileStmt) Visit(v AstVisitor) {
	v.VisitWhileStmt(n)
}

// ----- ForStmt --------------------------------------------------------------

type ForStmt struct {
	token lex.Token
	iter  lex.Token
	from  Expr
	to    Expr
	by    Expr
	stmt  Stmt
}

var _ Stmt = (*ForStmt)(nil)

func NewForStmt(t, i lex.Token, fr, to, by Expr, s Stmt) *ForStmt {
	return &ForStmt{
		token: t,
		iter:  i,
		from:  fr,
		to:    to,
		by:    by,
		stmt:  s,
	}
}

func (n *ForStmt) Token() lex.Token {
	return n.token
}

func (n *ForStmt) Iter() lex.Token {
	return n.iter
}

func (n *ForStmt) From() Expr {
	return n.from
}

func (n *ForStmt) To() Expr {
	return n.to
}

func (n *ForStmt) By() Expr {
	return n.by
}

func (n *ForStmt) Stmt() Stmt {
	return n.stmt
}

func (n *ForStmt) Visit(v AstVisitor) {
	v.VisitForStmt(n)
}

// ----- AssignStmt -----------------------------------------------------------

type AssignStmt struct {
	token lex.Token
	expr  Expr
}

var _ Stmt = (*AssignStmt)(nil)

func NewAssignStmt(t lex.Token, e Expr) *AssignStmt {
	return &AssignStmt{
		token: t,
		expr:  e,
	}
}

func (n *AssignStmt) Token() lex.Token {
	return n.token
}

func (n *AssignStmt) Expr() Expr {
	return n.expr
}

func (n *AssignStmt) Visit(v AstVisitor) {
	v.VisitAssignStmt(n)
}

// ----- ExprStmt -------------------------------------------------------------

type ExprStmt struct {
	expr Expr
}

var _ Stmt = (*ExprStmt)(nil)

func NewExprStmt(e Expr) *ExprStmt {
	return &ExprStmt{
		expr: e,
	}
}

func (n *ExprStmt) Token() lex.Token {
	return n.expr.Token()
}

func (n *ExprStmt) Expr() Expr {
	return n.expr
}

func (n *ExprStmt) Visit(v AstVisitor) {
	v.VisitExprStmt(n)
}

// ----- DesignatorExpr -------------------------------------------------------

type DesignatorExpr struct {
	token      lex.Token
	typ        Type
	constValue *Value
	kind       Kind
}

var _ Expr = (*DesignatorExpr)(nil)

func NewDesignatorExpr(t lex.Token) *DesignatorExpr {
	return &DesignatorExpr{
		token: t,
	}
}

func (n *DesignatorExpr) Token() lex.Token {
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

func (n *DesignatorExpr) Update(t Type, k Kind, c *Value) {
	n.typ = t
	n.kind = k
	n.constValue = c
}

// ----- NumberExpr -----------------------------------------------------------

type NumberExpr struct {
	token      lex.Token
	typ        Type
	constValue *Value
}

var _ Expr = (*NumberExpr)(nil)

func NewNumberExpr(t lex.Token) *NumberExpr {
	return &NumberExpr{
		token: t,
	}
}

func (n *NumberExpr) Token() lex.Token {
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

func (n *NumberExpr) Update(t Type, c *Value) {
	n.typ = t
	n.constValue = c
}

// ----- CharExpr -------------------------------------------------------------

type CharExpr struct {
	token      lex.Token
	constValue *Value
}

var _ Expr = (*CharExpr)(nil)

func NewCharExpr(t lex.Token) *CharExpr {
	return &CharExpr{
		token: t,
	}
}

func (n *CharExpr) Token() lex.Token {
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

func (n *CharExpr) Update(c *Value) {
	n.constValue = c
}

// ----- StringExpr -----------------------------------------------------------

type StringExpr struct {
	token      lex.Token
	constValue *Value
}

var _ Expr = (*StringExpr)(nil)

func NewStringExpr(t lex.Token) *StringExpr {
	return &StringExpr{
		token: t,
	}
}

func (n *StringExpr) Token() lex.Token {
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

func (n *StringExpr) Update(c *Value) {
	n.constValue = c
}

// ----- BooleanExpr ----------------------------------------------------------

type BooleanExpr struct {
	token      lex.Token
	constValue *Value
}

var _ Expr = (*BooleanExpr)(nil)

func NewBooleanExpr(t lex.Token) *BooleanExpr {
	return &BooleanExpr{
		token: t,
	}
}

func (n *BooleanExpr) Token() lex.Token {
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

func (n *BooleanExpr) Update(c *Value) {
	n.constValue = c
}

// ----- SetExpr --------------------------------------------------------------

type SetExpr struct {
	token      lex.Token
	bits       []byte
	constValue *Value
}

var _ Expr = (*SetExpr)(nil)

func NewSetExpr(t lex.Token, b []byte) *SetExpr {
	return &SetExpr{
		token: t,
		bits:  b,
	}
}

func (n *SetExpr) Token() lex.Token {
	return n.token
}

func (n *SetExpr) Type() Type {
	return TypeSet
}

func (n *SetExpr) ConstValue() *Value {
	return n.constValue
}

func (n *SetExpr) Bits() []byte {
	return n.bits
}

func (n *SetExpr) Visit(v AstVisitor) {
	v.VisitSetExpr(n)
}

func (n *SetExpr) Update(c *Value) {
	n.constValue = c
}

// ----- NotExpr --------------------------------------------------------------

type NotExpr struct {
	token      lex.Token
	expr       Expr
	constValue *Value
}

var _ Expr = (*NotExpr)(nil)

func NewNotExpr(t lex.Token, e Expr) *NotExpr {
	return &NotExpr{
		token: t,
		expr:  e,
	}
}

func (n *NotExpr) Token() lex.Token {
	return n.token
}

func (n *NotExpr) Type() Type {
	return TypeBoolean
}

func (n *NotExpr) ConstValue() *Value {
	return n.constValue
}

func (n *NotExpr) Expr() Expr {
	return n.expr
}

func (n *NotExpr) Visit(v AstVisitor) {
	v.VisitNotExpr(n)
}

func (n *NotExpr) Update(c *Value) {
	n.constValue = c
}

// ----- BinaryExpr -----------------------------------------------------------

type BinaryExpr struct {
	token      lex.Token
	args       []Expr
	typ        Type
	constValue *Value
}

var _ Expr = (*BinaryExpr)(nil)

func NewBinaryExpr(t lex.Token, lhs, rhs Expr) *BinaryExpr {
	return &BinaryExpr{
		token: t,
		args:  []Expr{lhs, rhs},
	}
}

func (n *BinaryExpr) Token() lex.Token {
	return n.token
}

func (n *BinaryExpr) Type() Type {
	return n.typ
}

func (n *BinaryExpr) ConstValue() *Value {
	return n.constValue
}

func (n *BinaryExpr) Args() []Expr {
	return n.args
}

func (n *BinaryExpr) Lhs() Expr {
	return n.args[0]
}

func (n *BinaryExpr) Rhs() Expr {
	return n.args[1]
}

func (n *BinaryExpr) Visit(v AstVisitor) {
	v.VisitBinaryExpr(n)
}

func (n *BinaryExpr) Update(t Type, c *Value) {
	n.typ = t
	n.constValue = c
}

// ----- CallExpr -------------------------------------------------------------

type CallExpr struct {
	designator *DesignatorExpr
	typ        Type
	args       []Expr
	constValue *Value
	builtin    Function
}

var _ Expr = (*CallExpr)(nil)

func NewCallExpr(d *DesignatorExpr, a []Expr) *CallExpr {
	return &CallExpr{
		designator: d,
		args:       a,
	}
}

func (n *CallExpr) Token() lex.Token {
	return n.designator.Token()
}

func (n *CallExpr) Type() Type {
	return n.typ
}

func (n *CallExpr) ConstValue() *Value {
	return n.constValue
}

func (n *CallExpr) Args() []Expr {
	return n.args
}

func (n *CallExpr) Builtin() Function {
	return n.builtin
}

func (n *CallExpr) Visit(v AstVisitor) {
	v.VisitCallExpr(n)
}

func (n *CallExpr) Update(t Type, c *Value, b Function) {
	n.typ = t
	n.constValue = c
	n.builtin = b
}
