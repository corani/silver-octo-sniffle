package ast

import (
	"fmt"

	"github.com/corani/silver-octo-sniffle/token"
)

type Kind int

const (
	KindUndefined Kind = iota
	KindVar
	KindConst
	KindType
	KindProc
	KindModule
)

func (k Kind) String() string {
	switch k {
	case KindUndefined:
		return "undefined"
	case KindVar:
		return "variable"
	case KindConst:
		return "constant"
	case KindType:
		return "type"
	case KindProc:
		return "procedure"
	case KindModule:
		return "module"
	default:
		return fmt.Sprintf("undefined=%d", int(k))
	}
}

type Symbol interface {
	Kind() Kind
	Type() Type
}

func GetBuiltinSymbols() map[string]Symbol {
	return map[string]Symbol{
		"INTEGER": newBuiltinTypeBaseDecl("INTEGER", TypeInt64),
		"REAL":    newBuiltinTypeBaseDecl("REAL", TypeFloat64),
		"BOOLEAN": newBuiltinTypeBaseDecl("BOOLEAN", TypeBoolean),
		"CHAR":    newBuiltinTypeBaseDecl("CHAR", TypeChar),
		"SET":     newBuiltinTypeBaseDecl("SET", TypeSet),
	}
}

type Decl interface {
	Symbol
	Node
}

// ----- ConstDecl ------------------------------------------------------------

func NewConstDecl(t token.Token, e Expr) *ConstDecl {
	return &ConstDecl{
		token: t,
		expr:  e,
	}
}

type ConstDecl struct {
	token token.Token
	expr  Expr
	typ   Type
	value *Value
}

var _ Decl = (*ConstDecl)(nil)

func (d *ConstDecl) Kind() Kind {
	return KindConst
}

func (d *ConstDecl) Token() token.Token {
	return d.token
}

func (d *ConstDecl) Expr() Expr {
	return d.expr
}

func (d *ConstDecl) Type() Type {
	return d.typ
}

func (d *ConstDecl) Value() *Value {
	return d.value
}

func (d *ConstDecl) Update(t Type, v *Value) {
	d.typ = t
	d.value = v
}

func (d *ConstDecl) Visit(v AstVisitor) {
	v.VisitConstDecl(d)
}

// ----- VarDecl --------------------------------------------------------------

func NewVarDecl(token token.Token, typeDecl TypeDecl) *VarDecl {
	switch d := typeDecl.(type) {
	case *TypeRef:
		return &VarDecl{
			token:   token,
			typeRef: d,
		}
	default:
		return &VarDecl{
			token:    token,
			typeDecl: d,
		}
	}
}

type VarDecl struct {
	token    token.Token
	typeRef  *TypeRef
	typeDecl TypeDecl
	typ      Type
}

var _ Decl = (*VarDecl)(nil)

func (d *VarDecl) Kind() Kind {
	return KindVar
}

func (d *VarDecl) Token() token.Token {
	return d.token
}

func (d *VarDecl) TypeRef() *TypeRef {
	return d.typeRef
}

func (d *VarDecl) TypeDecl() TypeDecl {
	return d.typeDecl
}

func (d *VarDecl) Type() Type {
	return d.typ
}

func (d *VarDecl) Update(td TypeDecl, t Type) {
	d.typeRef = nil
	d.typeDecl = td
	d.typ = t
}

func (d *VarDecl) Visit(v AstVisitor) {
	v.VisitVarDecl(d)
}

// ----- ProcDecl -------------------------------------------------------------

func NewProcDecl(
	token token.Token, params []*VarDecl, ret TypeDecl, decls []Decl, stmts Stmt,
) *ProcDecl {
	return &ProcDecl{
		token:  token,
		params: params,
		ret:    ret,
		decls:  decls,
		stmts:  stmts,
	}
}

type ProcDecl struct {
	token  token.Token
	params []*VarDecl
	ret    TypeDecl
	decls  []Decl
	stmts  Stmt
	typ    Type
}

var _ Decl = (*ProcDecl)(nil)

func (d *ProcDecl) Kind() Kind {
	return KindProc
}

func (d *ProcDecl) Token() token.Token {
	return d.token
}

func (d *ProcDecl) Type() Type {
	return d.typ
}

func (d *ProcDecl) Params() []*VarDecl {
	return d.params
}

func (d *ProcDecl) ReturnType() TypeDecl {
	return d.ret
}

func (d *ProcDecl) Decls() []Decl {
	return d.decls
}

func (d *ProcDecl) Stmts() Stmt {
	return d.stmts
}

func (d *ProcDecl) Visit(v AstVisitor) {
	v.VisitProcDecl(d)
}
