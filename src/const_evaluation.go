package main

import "fmt"

func evaluateNumberExpr(lit Token) (*Value, error) {
	switch lit.Type {
	case TokenInteger:
		return &Value{
			typ:     TypeInt64,
			integer: lit.Int,
		}, nil
	case TokenReal:
		return &Value{
			typ:  TypeFloat64,
			real: lit.Real,
		}, nil
	default:
		return nil, fmt.Errorf("invalid number literal: %q", lit.Text)
	}
}

func evaluateCharExpr(lit Token) *Value {
	if len(lit.Text) == 0 {
		return nil
	}

	return &Value{
		typ:  TypeChar,
		char: byte(lit.Text[0]),
	}
}

func evaluateStringExpr(lit Token) *Value {
	return &Value{
		typ: TypeString,
		str: lit.Text,
	}
}

func evaluateBooleanExpr(lit Token) (*Value, error) {
	var v bool

	switch lit.Text {
	case "TRUE":
		v = true
	case "FALSE":
		v = false
	default:
		return nil, fmt.Errorf("invalid boolean literal: %q", lit.Text)
	}

	return &Value{
		typ:     TypeBoolean,
		boolean: v,
	}, nil
}

func evaluateSetExpr(expr *SetExpr) *Value {
	var val int64

	for _, bit := range expr.bits {
		val |= 1 << bit
	}

	// TODO(daniel): use a uint64 for type set?
	return &Value{
		typ:     TypeSet,
		integer: int(val),
	}
}

func evaluateNotExpr(expr Expr) (*Value, error) {
	if expr.Type() != TypeBoolean {
		return nil, fmt.Errorf("`~` not supported for type %v", expr.Type())
	}

	if c := expr.ConstValue(); c != nil {
		return &Value{
			typ:     TypeBoolean,
			boolean: !c.Bool(),
		}, nil
	}

	//nolint:nilnil
	return nil, nil
}

func evaluateBinaryExpr(op TokenType, target Type, lhs, rhs *Value) *Value {
	// TODO(daniel): check if set operations are correct!
	switch op {
	case TokenPlus:
		switch target {
		case TypeInt64:
			return &Value{
				typ:     target,
				integer: lhs.Int() + rhs.Int(),
			}
		case TypeFloat64:
			return &Value{
				typ:  target,
				real: lhs.Real() + rhs.Real(),
			}
		case TypeSet: // set union
			return &Value{
				typ:     target,
				integer: lhs.Int() | rhs.Int(),
			}
		default:
			return nil
		}
	case TokenMinus:
		switch target {
		case TypeInt64:
			return &Value{
				typ:     target,
				integer: lhs.Int() - rhs.Int(),
			}
		case TypeFloat64:
			return &Value{
				typ:  target,
				real: lhs.Real() - rhs.Real(),
			}
		case TypeSet: // set difference
			return &Value{
				typ:     target,
				integer: lhs.Int() &^ rhs.Int(),
			}
		default:
			return nil
		}
	case TokenAsterisk:
		switch target {
		case TypeInt64:
			return &Value{
				typ:     target,
				integer: lhs.Int() * rhs.Int(),
			}
		case TypeFloat64:
			return &Value{
				typ:  target,
				real: lhs.Real() * rhs.Real(),
			}
		case TypeSet: // set intersection
			return &Value{
				typ:     target,
				integer: lhs.Int() & rhs.Int(),
			}
		default:
			return nil
		}
	case TokenSlash:
		switch target {
		case TypeFloat64: // REAL division
			return &Value{
				typ:  target,
				real: lhs.Real() / rhs.Real(),
			}
		case TypeSet: // symmetric set difference
			return &Value{
				typ:     target,
				integer: lhs.Int() ^ rhs.Int(),
			}
		default:
			return nil
		}
	case TokenDIV: // INTEGER division
		switch target {
		case TypeInt64:
			return &Value{
				typ:     target,
				integer: lhs.Int() / rhs.Int(),
			}
		default:
			return nil
		}
	case TokenMOD:
		switch target {
		case TypeInt64:
			return &Value{
				typ:     target,
				integer: lhs.Int() % rhs.Int(),
			}
		default:
			return nil
		}
	case TokenOR:
		switch target {
		case TypeBoolean:
			return &Value{
				typ:     target,
				boolean: lhs.Bool() || rhs.Bool(),
			}
		default:
			return nil
		}
	case TokenAmpersand:
		switch target {
		case TypeBoolean:
			return &Value{
				typ:     target,
				boolean: lhs.Bool() && rhs.Bool(),
			}
		default:
			return nil
		}
	case TokenEQ:
		switch target {
		case TypeBoolean:
			if lhs.Type() == TypeFloat64 || rhs.Type() == TypeFloat64 {
				return &Value{
					typ:     target,
					boolean: lhs.Real() == rhs.Real(),
				}
			} else {
				return &Value{
					typ:     target,
					boolean: lhs.Int() == rhs.Int(),
				}
			}
		default:
			return nil
		}
	case TokenNE:
		switch target {
		case TypeBoolean:
			if lhs.Type() == TypeFloat64 || rhs.Type() == TypeFloat64 {
				return &Value{
					typ:     target,
					boolean: lhs.Real() != rhs.Real(),
				}
			} else {
				return &Value{
					typ:     target,
					boolean: lhs.Int() != rhs.Int(),
				}
			}
		default:
			return nil
		}
	case TokenLT:
		switch target {
		case TypeBoolean:
			if lhs.Type() == TypeFloat64 || rhs.Type() == TypeFloat64 {
				return &Value{
					typ:     target,
					boolean: lhs.Real() < rhs.Real(),
				}
			} else {
				return &Value{
					typ:     target,
					boolean: lhs.Int() < rhs.Int(),
				}
			}
		default:
			return nil
		}
	case TokenLE:
		switch target {
		case TypeBoolean:
			if lhs.Type() == TypeFloat64 || rhs.Type() == TypeFloat64 {
				return &Value{
					typ:     target,
					boolean: lhs.Real() <= rhs.Real(),
				}
			} else {
				return &Value{
					typ:     target,
					boolean: lhs.Int() <= rhs.Int(),
				}
			}
		default:
			return nil
		}
	case TokenGE:
		switch target {
		case TypeBoolean:
			if lhs.Type() == TypeFloat64 || rhs.Type() == TypeFloat64 {
				return &Value{
					typ:     target,
					boolean: lhs.Real() >= rhs.Real(),
				}
			} else {
				return &Value{
					typ:     target,
					boolean: lhs.Int() >= rhs.Int(),
				}
			}
		default:
			return nil
		}
	case TokenGT:
		switch target {
		case TypeBoolean:
			if lhs.Type() == TypeFloat64 || rhs.Type() == TypeFloat64 {
				return &Value{
					typ:     target,
					boolean: lhs.Real() > rhs.Real(),
				}
			} else {
				return &Value{
					typ:     target,
					boolean: lhs.Int() > rhs.Int(),
				}
			}
		default:
			return nil
		}
	case TokenIN:
		if lhs.Type() == TypeInt64 && rhs.Type() == TypeSet {
			return &Value{
				typ:     target,
				boolean: rhs.Int()&(1<<lhs.Int()) != 0,
			}
		} else {
			return nil
		}
	case TokenIS:
		return nil
	default:
		return nil
	}
}
