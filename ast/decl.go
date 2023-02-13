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
		"INTEGER": newBuiltinTypeDecl("INTEGER", TypeInt64),
		"REAL":    newBuiltinTypeDecl("REAL", TypeFloat64),
		"BOOLEAN": newBuiltinTypeDecl("BOOLEAN", TypeBoolean),
		"CHAR":    newBuiltinTypeDecl("CHAR", TypeChar),
		"SET":     newBuiltinTypeDecl("SET", TypeSet),
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

// ----- TypeDecl -------------------------------------------------------------

func NewTypeDecl(token, typeToken token.Token) *TypeDecl {
	return &TypeDecl{
		token:     token,
		typeToken: typeToken,
	}
}

func newBuiltinTypeDecl(name string, typ Type) *TypeDecl {
	t := token.Token{
		Type: token.TokenIdent,
		File: "<builtin>",
		Text: name,
	}

	return &TypeDecl{
		token:     t,
		typeToken: t,
		typ:       typ,
	}
}

type TypeDecl struct {
	token     token.Token
	typeToken token.Token
	typ       Type
}

var _ Decl = (*TypeDecl)(nil)

func (d *TypeDecl) Kind() Kind {
	return KindType
}

func (d *TypeDecl) Token() token.Token {
	return d.token
}

func (d *TypeDecl) TypeToken() token.Token {
	return d.typeToken
}

func (d *TypeDecl) Type() Type {
	return d.typ
}

func (d *TypeDecl) Update(t Type) {
	d.typ = t
}

func (d *TypeDecl) Visit(v AstVisitor) {
	v.VisitTypeDecl(d)
}

// ----- VarDecl --------------------------------------------------------------

func NewVarDecl(token, typeToken token.Token) *VarDecl {
	return &VarDecl{
		token:     token,
		typeToken: typeToken,
	}
}

type VarDecl struct {
	token     token.Token
	typeToken token.Token
	typ       Type
}

var _ Decl = (*VarDecl)(nil)

func (d *VarDecl) Kind() Kind {
	return KindVar
}

func (d *VarDecl) Token() token.Token {
	return d.token
}

func (d *VarDecl) TypeToken() token.Token {
	return d.typeToken
}

func (d *VarDecl) Type() Type {
	return d.typ
}

func (d *VarDecl) Update(t Type) {
	d.typ = t
}

func (d *VarDecl) Visit(v AstVisitor) {
	v.VisitVarDecl(d)
}

// ----- VarDecl --------------------------------------------------------------

func NewProcDecl(token token.Token) *ProcDecl {
	return &ProcDecl{
		token: token,
	}
}

type ProcDecl struct {
	token token.Token
	typ   Type
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

func (d *ProcDecl) Visit(v AstVisitor) {
	v.VisitProcDecl(d)
}
