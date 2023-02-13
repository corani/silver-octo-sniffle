package ast

import (
	"fmt"
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
	TypeArray
	TypePointer
	TypeRecord
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
	default:
		return fmt.Sprintf("undefined=%d", int(t))
	}
}
