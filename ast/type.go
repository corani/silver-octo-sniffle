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

// ----- TypeRef --------------------------------------------------------------

func NewTypeRef(q *QualIdent) *TypeRef {
	return &TypeRef{
		qualIdent: q,
	}
}

type TypeRef struct {
	qualIdent *QualIdent
	typeDecl  TypeDecl
}

var _ TypeDecl = (*TypeRef)(nil)

func (t *TypeRef) Token() token.Token {
	return t.qualIdent.Token()
}

func (t *TypeRef) Visit(v AstVisitor) {
	v.VisitTypeRef(t)
}

func (t *TypeRef) Kind() Kind {
	return KindUndefined
}

func (t *TypeRef) Type() Type {
	return TypeVoid
}

func (t *TypeRef) TypeDecl() TypeDecl {
	return t.typeDecl
}

func (t *TypeRef) Update(decl TypeDecl) {
	t.typeDecl = decl
}

// ----- TypeBaseDecl -----------------------------------------------------------------------------

func NewInvalidTypeDecl(token token.Token) *TypeBaseDecl {
	return &TypeBaseDecl{
		token: token,
	}
}

func newBuiltinTypeBaseDecl(name string, typ Type) *TypeBaseDecl {
	t := token.Token{
		Type: token.TokenIdent,
		File: "<builtin>",
		Text: name,
	}

	return &TypeBaseDecl{
		token:   t,
		typeRef: nil,
		typ:     typ,
	}
}

type TypeBaseDecl struct {
	token   token.Token
	typeRef *TypeRef
	typ     Type
}

var _ TypeDecl = (*TypeBaseDecl)(nil)

func (d *TypeBaseDecl) Kind() Kind {
	return KindType
}

func (d *TypeBaseDecl) Token() token.Token {
	return d.token
}

func (d *TypeBaseDecl) Type() Type {
	return d.typ
}

func (d *TypeBaseDecl) TypeRef() *TypeRef {
	return d.typeRef
}

func (d *TypeBaseDecl) Update(token token.Token, typ Type) {
	d.typeRef = nil
	d.token = token
	d.typ = typ
}

func (d *TypeBaseDecl) Visit(v AstVisitor) {
	v.VisitTypeBaseDecl(d)
}

// ----- PointerTypeDecl --------------------------------------------------------------------------

func NewTypePointerDecl(token token.Token, to TypeDecl) *TypePointerDecl {
	if ref, ok := to.(*TypeRef); ok {
		return &TypePointerDecl{
			token:   token,
			typeRef: ref,
		}
	}

	return &TypePointerDecl{
		token: token,
		to:    to,
	}
}

type TypePointerDecl struct {
	token   token.Token
	typeRef *TypeRef
	to      TypeDecl
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

func (d *TypePointerDecl) TypeRef() *TypeRef {
	return d.typeRef
}

func (d *TypePointerDecl) To() TypeDecl {
	return d.to
}

func (d *TypePointerDecl) Update(to TypeDecl) {
	d.typeRef = nil
	d.to = to
}

func (d *TypePointerDecl) Visit(v AstVisitor) {
	v.VisitTypePointerDecl(d)
}
