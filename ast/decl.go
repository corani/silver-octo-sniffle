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

// ----- ConstDecl ------------------------------------------------------------

type ConstDecl struct {
	token token.Token
	expr  Expr
	typ   Type
	value *Value
}

func NewConstDecl(t token.Token, e Expr) *ConstDecl {
	return &ConstDecl{
		token: t,
		expr:  e,
	}
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

// ----- TypeDecl -------------------------------------------------------------

type TypeDecl struct {
	token     token.Token
	typeToken token.Token
	typ       Type
}

func NewTypeDecl(token, typeToken token.Token) *TypeDecl {
	return &TypeDecl{
		token:     token,
		typeToken: typeToken,
	}
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

// ----- VarDecl --------------------------------------------------------------

type VarDecl struct {
	token     token.Token
	typeToken token.Token
	typ       Type
}

func NewVarDecl(token, typeToken token.Token) *VarDecl {
	return &VarDecl{
		token:     token,
		typeToken: typeToken,
	}
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
