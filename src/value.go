package main

import "fmt"

type Type int

const (
	TypeVoid Type = iota
	TypeInt64
	TypeFloat64
	TypeString
	TypeBoolean
	TypeChar
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
	default:
		return fmt.Sprintf("undefined=%d", int(t))
	}
}

type Value struct {
	typ     Type
	integer int
	real    float64
	boolean bool
	str     string
	char    byte
}

func (v Value) Type() Type {
	return v.typ
}

func (v Value) Int() int {
	switch v.typ {
	case TypeInt64:
		return v.integer
	case TypeFloat64:
		return int(v.real)
	case TypeBoolean:
		if v.boolean {
			return 1
		}

		return 0
	case TypeChar:
		return int(v.char)
	default:
		return 0
	}
}

func (v Value) Real() float64 {
	switch v.typ {
	case TypeInt64:
		return float64(v.integer)
	case TypeFloat64:
		return v.real
	default:
		return 0
	}
}

func (v Value) Char() byte {
	switch v.typ {
	case TypeChar:
		return v.char
	case TypeInt64:
		return byte(v.integer)
	case TypeString:
		if len(v.str) > 0 {
			return byte(v.str[0])
		}

		return 0
	default:
		return 0
	}
}

func (v Value) String() string {
	switch v.typ {
	case TypeString:
		return v.str
	case TypeChar:
		return string([]byte{v.char})
	default:
		return ""
	}
}

func (v Value) Bool() bool {
	return v.boolean
}
