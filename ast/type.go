package ast

import (
	"fmt"

	"github.com/corani/silver-octo-sniffle/token"
)

type Type int

const (
	TypeVoid Type = iota
	TypeInt64
	TypeFloat64
	TypeString
	TypeBoolean
	TypeChar
	TypeSet
	TypePointer
	TypeArray
	TypeRecord
	TypeProc
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
	case TypeChar:
		return "char"
	case TypeSet:
		return "set"
	case TypePointer:
		return "pointer"
	case TypeArray:
		return "array"
	case TypeRecord:
		return "record"
	case TypeProc:
		return "procedure"
	default:
		return fmt.Sprintf("undefined=%d", int(t))
	}
}

type TypeDecl interface {
	Decl
}

// ----- TypeBaseDecl -----------------------------------------------------------------------------

func NewInvalidTypeDecl(token token.Token) *TypeBaseDecl {
	return &TypeBaseDecl{
		token: token,
	}
}

func NewTypeBaseDecl(token, typeToken token.Token) *TypeBaseDecl {
	return &TypeBaseDecl{
		token:     token,
		typeToken: typeToken,
	}
}

func newBuiltinTypeBaseDecl(name string, typ Type) *TypeBaseDecl {
	t := token.Token{
		Type: token.TokenIdent,
		File: "<builtin>",
		Text: name,
	}

	return &TypeBaseDecl{
		token:     t,
		typeToken: t,
		typ:       typ,
	}
}

type TypeBaseDecl struct {
	token     token.Token
	typeToken token.Token
	typ       Type
}

var _ TypeDecl = (*TypeBaseDecl)(nil)

func (d *TypeBaseDecl) Kind() Kind {
	return KindType
}

func (d *TypeBaseDecl) Token() token.Token {
	return d.token
}

func (d *TypeBaseDecl) TypeToken() token.Token {
	return d.typeToken
}

func (d *TypeBaseDecl) Type() Type {
	return d.typ
}

func (d *TypeBaseDecl) Update(token token.Token, typ Type) {
	d.token = token
	d.typ = typ
}

func (d *TypeBaseDecl) Visit(v AstVisitor) {
	v.VisitTypeBaseDecl(d)
}

// ----- PointerTypeDecl --------------------------------------------------------------------------

func NewTypePointerDecl(token token.Token, to TypeDecl) *TypePointerDecl {
	return &TypePointerDecl{
		token: token,
		to:    to,
	}
}

type TypePointerDecl struct {
	token token.Token
	to    TypeDecl
}

var _ TypeDecl = (*TypePointerDecl)(nil)

func (d *TypePointerDecl) Kind() Kind {
	return KindType
}

func (d *TypePointerDecl) Type() Type {
	return TypePointer
}

func (d *TypePointerDecl) Token() token.Token {
	return d.token
}

func (d *TypePointerDecl) To() TypeDecl {
	return d.to
}

func (d *TypePointerDecl) Visit(v AstVisitor) {
	v.VisitTypePointerDecl(d)
}
