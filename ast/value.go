package ast

type Value struct {
	typ     Type
	integer int
	real    float64
	boolean bool
	str     string
	char    byte
}

func NewInt(v int) *Value {
	return &Value{
		typ:     TypeInt64,
		integer: v,
	}
}

func NewReal(v float64) *Value {
	return &Value{
		typ:  TypeFloat64,
		real: v,
	}
}

func NewChar(v byte) *Value {
	return &Value{
		typ:  TypeChar,
		char: v,
	}
}

func NewString(v string) *Value {
	return &Value{
		typ: TypeString,
		str: v,
	}
}

func NewBoolean(v bool) *Value {
	return &Value{
		typ:     TypeBoolean,
		boolean: v,
	}
}

// TODO(daniel): use a uint64 for type set?
func NewSet(v int) *Value {
	return &Value{
		typ:     TypeSet,
		integer: v,
	}
}

func (v Value) Type() Type {
	return v.typ
}

func (v Value) Int() int {
	switch v.typ {
	case TypeInt64, TypeSet:
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

func (v Value) Set() int {
	return v.integer
}
