package ast

import "github.com/corani/silver-octo-sniffle/token"

type AstVisitor interface {
	VisitModule(*Module)
	VisitInvalidStmt(*InvalidStmt)
	VisitStmtSequence(*StmtSequence)
	VisitIfStmt(*IfStmt)
	VisitRepeatStmt(*RepeatStmt)
	VisitWhileStmt(*WhileStmt)
	VisitForStmt(*ForStmt)
	VisitAssignStmt(*AssignStmt)
	VisitExprStmt(*ExprStmt)
	VisitInvalidExpr(*InvalidExpr)
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
	Token() token.Token
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
	token  token.Token
	name   string
	block  Stmt
	consts []*ConstDecl
	types  []*TypeDecl
	vars   []*VarDecl
}

var _ Node = (*Module)(nil)

// TODO(daniel): maybe instead of passing all decls in the constructor, add some `NewXxx` methods?
func NewModule(token token.Token, n string, s Stmt, c []*ConstDecl, t []*TypeDecl, v []*VarDecl) *Module {
	return &Module{
		token:  token,
		name:   n,
		block:  s,
		consts: c,
		types:  t,
		vars:   v,
	}
}

func (n *Module) Token() token.Token {
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

// ----- InvalidStatement ---------------------------------------------------------

type InvalidStmt struct {
	token token.Token
}

var _ Stmt = (*InvalidStmt)(nil)

func NewInvalidStmt(t token.Token) *InvalidStmt {
	return &InvalidStmt{
		token: t,
	}
}

func (n *InvalidStmt) Token() token.Token {
	return n.token
}

func (n *InvalidStmt) Visit(v AstVisitor) {
	v.VisitInvalidStmt(n)
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

func (n *StmtSequence) Token() token.Token {
	if len(n.stmts) == 0 {
		return token.Token{}
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
	token      token.Token
	expr       Expr
	trueBlock  Stmt
	falseBlock Stmt
}

var _ Stmt = (*IfStmt)(nil)

func NewIfStmt(t token.Token, e Expr, tb, fb Stmt) *IfStmt {
	return &IfStmt{
		token:      t,
		expr:       e,
		trueBlock:  tb,
		falseBlock: fb,
	}
}

func (n *IfStmt) Token() token.Token {
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
	token token.Token
	cond  Condition
}

var _ Stmt = (*RepeatStmt)(nil)

func NewRepeatStmt(t token.Token, c Condition) *RepeatStmt {
	return &RepeatStmt{
		token: t,
		cond:  c,
	}
}

func (n *RepeatStmt) Token() token.Token {
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
	token token.Token
	conds []Condition
}

var _ Stmt = (*WhileStmt)(nil)

func NewWhileStmt(t token.Token, c []Condition) *WhileStmt {
	return &WhileStmt{
		token: t,
		conds: c,
	}
}

func (n *WhileStmt) Token() token.Token {
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
	token token.Token
	iter  token.Token
	from  Expr
	to    Expr
	by    Expr
	stmt  Stmt
}

var _ Stmt = (*ForStmt)(nil)

func NewForStmt(t, i token.Token, fr, to, by Expr, s Stmt) *ForStmt {
	return &ForStmt{
		token: t,
		iter:  i,
		from:  fr,
		to:    to,
		by:    by,
		stmt:  s,
	}
}

func (n *ForStmt) Token() token.Token {
	return n.token
}

func (n *ForStmt) Iter() token.Token {
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
	token token.Token
	expr  Expr
}

var _ Stmt = (*AssignStmt)(nil)

func NewAssignStmt(t token.Token, e Expr) *AssignStmt {
	return &AssignStmt{
		token: t,
		expr:  e,
	}
}

func (n *AssignStmt) Token() token.Token {
	return n.token
}

func (n *AssignStmt) Expr() Expr {
	return n.expr
}

func (n *AssignStmt) Visit(v AstVisitor) {
	v.VisitAssignStmt(n)
}

// ----- InvalidExpr -----------------------------------------------------------

type InvalidExpr struct {
	token      token.Token
	typ        Type
	constValue *Value
}

var _ Expr = (*InvalidExpr)(nil)

func NewInvalidExpr(t token.Token) *InvalidExpr {
	return &InvalidExpr{
		token: t,
	}
}

func (n *InvalidExpr) Token() token.Token {
	return n.token
}

func (n *InvalidExpr) Type() Type {
	return n.typ
}

func (n *InvalidExpr) ConstValue() *Value {
	return n.constValue
}

func (n *InvalidExpr) Visit(v AstVisitor) {
	v.VisitInvalidExpr(n)
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

func (n *ExprStmt) Token() token.Token {
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
	token      token.Token
	typ        Type
	constValue *Value
	kind       Kind
}

var _ Expr = (*DesignatorExpr)(nil)

func NewDesignatorExpr(t token.Token) *DesignatorExpr {
	return &DesignatorExpr{
		token: t,
	}
}

func (n *DesignatorExpr) Token() token.Token {
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
	token      token.Token
	typ        Type
	constValue *Value
}

var _ Expr = (*NumberExpr)(nil)

func NewNumberExpr(t token.Token) *NumberExpr {
	return &NumberExpr{
		token: t,
	}
}

func (n *NumberExpr) Token() token.Token {
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
	token      token.Token
	constValue *Value
}

var _ Expr = (*CharExpr)(nil)

func NewCharExpr(t token.Token) *CharExpr {
	return &CharExpr{
		token: t,
	}
}

func (n *CharExpr) Token() token.Token {
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
	token      token.Token
	constValue *Value
}

var _ Expr = (*StringExpr)(nil)

func NewStringExpr(t token.Token) *StringExpr {
	return &StringExpr{
		token: t,
	}
}

func (n *StringExpr) Token() token.Token {
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
	token      token.Token
	constValue *Value
}

var _ Expr = (*BooleanExpr)(nil)

func NewBooleanExpr(t token.Token) *BooleanExpr {
	return &BooleanExpr{
		token: t,
	}
}

func (n *BooleanExpr) Token() token.Token {
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
	token      token.Token
	bits       []byte
	constValue *Value
}

var _ Expr = (*SetExpr)(nil)

func NewSetExpr(t token.Token, b []byte) *SetExpr {
	return &SetExpr{
		token: t,
		bits:  b,
	}
}

func (n *SetExpr) Token() token.Token {
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
	token      token.Token
	expr       Expr
	constValue *Value
}

var _ Expr = (*NotExpr)(nil)

func NewNotExpr(t token.Token, e Expr) *NotExpr {
	return &NotExpr{
		token: t,
		expr:  e,
	}
}

func (n *NotExpr) Token() token.Token {
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
	token      token.Token
	args       []Expr
	typ        Type
	constValue *Value
}

var _ Expr = (*BinaryExpr)(nil)

func NewBinaryExpr(t token.Token, lhs, rhs Expr) *BinaryExpr {
	return &BinaryExpr{
		token: t,
		args:  []Expr{lhs, rhs},
	}
}

func (n *BinaryExpr) Token() token.Token {
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

func (n *CallExpr) Token() token.Token {
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
