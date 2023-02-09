package ast

import (
	"fmt"

	"github.com/corani/silver-octo-sniffle/lex"
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
	token lex.Token
	expr  Expr
	typ   Type
	value *Value
}

func NewConstDecl(t lex.Token, e Expr) *ConstDecl {
	return &ConstDecl{
		token: t,
		expr:  e,
	}
}

func (d *ConstDecl) Token() lex.Token {
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
	token     lex.Token
	typeToken lex.Token
	typ       Type
}

func NewTypeDecl(token, typeToken lex.Token) *TypeDecl {
	return &TypeDecl{
		token:     token,
		typeToken: typeToken,
	}
}

func (d *TypeDecl) Token() lex.Token {
	return d.token
}

func (d *TypeDecl) TypeToken() lex.Token {
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
	token     lex.Token
	typeToken lex.Token
	typ       Type
}

func NewVarDecl(token, typeToken lex.Token) *VarDecl {
	return &VarDecl{
		token:     token,
		typeToken: typeToken,
	}
}

func (d *VarDecl) Token() lex.Token {
	return d.token
}

func (d *VarDecl) TypeToken() lex.Token {
	return d.typeToken
}

func (d *VarDecl) Type() Type {
	return d.typ
}

func (d *VarDecl) Update(t Type) {
	d.typ = t
}
