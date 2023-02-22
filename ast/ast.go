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
	VisitReturnStmt(*ReturnStmt)
	VisitExprStmt(*ExprStmt)
	VisitInvalidExpr(*InvalidExpr)
	VisitCallExpr(*CallExpr)
	VisitBinaryExpr(*BinaryExpr)
	VisitDesignatorExpr(*DesignatorExpr)
	VisitNotExpr(*NotExpr)
	VisitNumberLit(*NumberLit)
	VisitStringLit(*StringLit)
	VisitCharLit(*CharLit)
	VisitBooleanLit(*BooleanLit)
	VisitSetLit(*SetLit)
	VisitConstDecl(*ConstDecl)
	VisitTypeBaseDecl(*TypeBaseDecl)
	VisitTypePointerDecl(*TypePointerDecl)
	VisitVarDecl(*VarDecl)
	VisitProcDecl(*ProcDecl)
	VisitTypeRef(*TypeRef)
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
	token token.Token
	name  string
	block Stmt
	decls []Decl
}

var _ Node = (*Module)(nil)

// TODO(daniel): maybe instead of passing all decls in the constructor, add some `NewXxx` methods?
func NewModule(token token.Token, n string, s Stmt, d []Decl) *Module {
	return &Module{
		token: token,
		name:  n,
		block: s,
		decls: d,
	}
}

func (n *Module) Token() token.Token {
	return n.token
}

func (n *Module) Block() Stmt {
	return n.block
}

func (n *Module) Decls() []Decl {
	return n.decls
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
	designator *DesignatorExpr
	expr       Expr
}

var _ Stmt = (*AssignStmt)(nil)

func NewAssignStmt(d *DesignatorExpr, e Expr) *AssignStmt {
	d.assignment = true

	return &AssignStmt{
		designator: d,
		expr:       e,
	}
}

func (n *AssignStmt) Token() token.Token {
	return n.designator.Token()
}

func (n *AssignStmt) Designator() *DesignatorExpr {
	return n.designator
}

func (n *AssignStmt) Expr() Expr {
	return n.expr
}

func (n *AssignStmt) Visit(v AstVisitor) {
	v.VisitAssignStmt(n)
}

// ----- ReturnStmt -----------------------------------------------------------

type ReturnStmt struct {
	token token.Token
	expr  Expr
}

var _ Stmt = (*ReturnStmt)(nil)

func NewReturnStmt(t token.Token, e Expr) *ReturnStmt {
	return &ReturnStmt{
		token: t,
		expr:  e,
	}
}

func (n *ReturnStmt) Token() token.Token {
	return n.token
}

func (n *ReturnStmt) Expr() Expr {
	return n.expr
}

func (n *ReturnStmt) Visit(v AstVisitor) {
	v.VisitReturnStmt(n)
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

func (n *CallExpr) Designator() *DesignatorExpr {
	return n.designator
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

// ----- NumberLit -----------------------------------------------------------

type NumberLit struct {
	token      token.Token
	typ        Type
	constValue *Value
}

var _ Expr = (*NumberLit)(nil)

func NewNumberExpr(t token.Token) *NumberLit {
	return &NumberLit{
		token: t,
	}
}

func (n *NumberLit) Token() token.Token {
	return n.token
}

func (n *NumberLit) Type() Type {
	return n.typ
}

func (n *NumberLit) ConstValue() *Value {
	return n.constValue
}

func (n *NumberLit) Visit(v AstVisitor) {
	v.VisitNumberLit(n)
}

func (n *NumberLit) Update(t Type, c *Value) {
	n.typ = t
	n.constValue = c
}

// ----- StringLit -----------------------------------------------------------

type StringLit struct {
	token      token.Token
	constValue *Value
}

var _ Expr = (*StringLit)(nil)

func NewStringExpr(t token.Token) *StringLit {
	return &StringLit{
		token: t,
	}
}

func (n *StringLit) Token() token.Token {
	return n.token
}

func (n *StringLit) Type() Type {
	return TypeString
}

func (n *StringLit) ConstValue() *Value {
	return n.constValue
}

func (n *StringLit) Visit(v AstVisitor) {
	v.VisitStringLit(n)
}

func (n *StringLit) Update(c *Value) {
	n.constValue = c
}

// ----- CharLit -------------------------------------------------------------

type CharLit struct {
	token      token.Token
	constValue *Value
}

var _ Expr = (*CharLit)(nil)

func NewCharExpr(t token.Token) *CharLit {
	return &CharLit{
		token: t,
	}
}

func (n *CharLit) Token() token.Token {
	return n.token
}

func (n *CharLit) Type() Type {
	return TypeChar
}

func (n *CharLit) ConstValue() *Value {
	return n.constValue
}

func (n *CharLit) Visit(v AstVisitor) {
	v.VisitCharLit(n)
}

func (n *CharLit) Update(c *Value) {
	n.constValue = c
}

// ----- BooleanLit ----------------------------------------------------------

type BooleanLit struct {
	token      token.Token
	constValue *Value
}

var _ Expr = (*BooleanLit)(nil)

func NewBooleanExpr(t token.Token) *BooleanLit {
	return &BooleanLit{
		token: t,
	}
}

func (n *BooleanLit) Token() token.Token {
	return n.token
}

func (n *BooleanLit) Type() Type {
	return TypeBoolean
}

func (n *BooleanLit) ConstValue() *Value {
	return n.constValue
}

func (n *BooleanLit) Visit(v AstVisitor) {
	v.VisitBooleanLit(n)
}

func (n *BooleanLit) Update(c *Value) {
	n.constValue = c
}

// ----- SetLit --------------------------------------------------------------

type SetLit struct {
	token      token.Token
	bits       []byte
	constValue *Value
}

var _ Expr = (*SetLit)(nil)

func NewSetExpr(t token.Token, b []byte) *SetLit {
	return &SetLit{
		token: t,
		bits:  b,
	}
}

func (n *SetLit) Token() token.Token {
	return n.token
}

func (n *SetLit) Type() Type {
	return TypeSet
}

func (n *SetLit) ConstValue() *Value {
	return n.constValue
}

func (n *SetLit) Bits() []byte {
	return n.bits
}

func (n *SetLit) Visit(v AstVisitor) {
	v.VisitSetLit(n)
}

func (n *SetLit) Update(c *Value) {
	n.constValue = c
}
