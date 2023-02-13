package checker

import (
	"fmt"

	"github.com/corani/silver-octo-sniffle/ast"
	"github.com/corani/silver-octo-sniffle/token"
)

func evaluateNumberLit(lit *ast.NumberLit) (*ast.Value, error) {
	switch lit.Token().Type {
	case token.TokenInteger:
		return ast.NewInt(lit.Token().Int), nil
	case token.TokenReal:
		return ast.NewReal(lit.Token().Real), nil
	default:
		return nil, fmt.Errorf("invalid number literal: %q", lit.Token().Text)
	}
}

func evaluateCharLit(lit *ast.CharLit) *ast.Value {
	if len(lit.Token().Text) == 0 {
		return nil
	}

	return ast.NewChar(byte(lit.Token().Text[0]))
}

func evaluateStringLit(lit *ast.StringLit) *ast.Value {
	return ast.NewString(lit.Token().Text)
}

func evaluateBooleanLit(lit *ast.BooleanLit) (*ast.Value, error) {
	var v bool

	switch lit.Token().Text {
	case "TRUE":
		v = true
	case "FALSE":
		v = false
	default:
		return nil, fmt.Errorf("invalid boolean literal: %q", lit.Token().Text)
	}

	return ast.NewBoolean(v), nil
}

func evaluateSetLit(lit *ast.SetLit) *ast.Value {
	var val int64

	for _, bit := range lit.Bits() {
		val |= 1 << bit
	}

	return ast.NewSet(int(val))
}

func evaluateNotExpr(expr ast.Expr) (*ast.Value, error) {
	if expr.Type() != ast.TypeBoolean {
		return nil, fmt.Errorf("`~` not supported for type %v", expr.Type())
	}

	if c := expr.ConstValue(); c != nil {
		return ast.NewBoolean(!c.Bool()), nil
	}

	return nil, nil
}

func evaluateBinaryExpr(op token.TokenType, target ast.Type, lhs, rhs *ast.Value) *ast.Value {
	switch op {
	case token.TokenPlus:
		switch target {
		case ast.TypeInt64:
			return ast.NewInt(lhs.Int() + rhs.Int())
		case ast.TypeFloat64:
			return ast.NewReal(lhs.Real() + rhs.Real())
		case ast.TypeSet: // set union
			return ast.NewSet(lhs.Set() | rhs.Set())
		default:
			return nil
		}
	case token.TokenMinus:
		switch target {
		case ast.TypeInt64:
			return ast.NewInt(lhs.Int() - rhs.Int())
		case ast.TypeFloat64:
			return ast.NewReal(lhs.Real() - rhs.Real())
		case ast.TypeSet: // set difference
			return ast.NewSet(lhs.Set() &^ rhs.Set())
		default:
			return nil
		}
	case token.TokenAsterisk:
		switch target {
		case ast.TypeInt64:
			return ast.NewInt(lhs.Int() * rhs.Int())
		case ast.TypeFloat64:
			return ast.NewReal(lhs.Real() * rhs.Real())
		case ast.TypeSet: // set intersection
			return ast.NewSet(lhs.Set() & rhs.Set())
		default:
			return nil
		}
	case token.TokenSlash:
		switch target {
		case ast.TypeFloat64: // REAL division
			return ast.NewReal(lhs.Real() / rhs.Real())
		case ast.TypeSet: // symmetric set difference
			return ast.NewSet(lhs.Set() ^ rhs.Set())
		default:
			return nil
		}
	case token.TokenDIV: // INTEGER division
		switch target {
		case ast.TypeInt64:
			return ast.NewInt(lhs.Int() / rhs.Int())
		default:
			return nil
		}
	case token.TokenMOD:
		switch target {
		case ast.TypeInt64:
			return ast.NewInt(lhs.Int() % rhs.Int())
		default:
			return nil
		}
	case token.TokenOR:
		switch target {
		case ast.TypeBoolean:
			return ast.NewBoolean(lhs.Bool() || rhs.Bool())
		default:
			return nil
		}
	case token.TokenAmpersand:
		switch target {
		case ast.TypeBoolean:
			return ast.NewBoolean(lhs.Bool() && rhs.Bool())
		default:
			return nil
		}
	case token.TokenEQ:
		switch target {
		case ast.TypeBoolean:
			if lhs.Type() == ast.TypeFloat64 || rhs.Type() == ast.TypeFloat64 {
				return ast.NewBoolean(lhs.Real() == rhs.Real())
			} else {
				return ast.NewBoolean(lhs.Int() == rhs.Int())
			}
		default:
			return nil
		}
	case token.TokenNE:
		switch target {
		case ast.TypeBoolean:
			if lhs.Type() == ast.TypeFloat64 || rhs.Type() == ast.TypeFloat64 {
				return ast.NewBoolean(lhs.Real() != rhs.Real())
			} else {
				return ast.NewBoolean(lhs.Int() != rhs.Int())
			}
		default:
			return nil
		}
	case token.TokenLT:
		switch target {
		case ast.TypeBoolean:
			if lhs.Type() == ast.TypeFloat64 || rhs.Type() == ast.TypeFloat64 {
				return ast.NewBoolean(lhs.Real() < rhs.Real())
			} else {
				return ast.NewBoolean(lhs.Int() < rhs.Int())
			}
		default:
			return nil
		}
	case token.TokenLE:
		switch target {
		case ast.TypeBoolean:
			if lhs.Type() == ast.TypeFloat64 || rhs.Type() == ast.TypeFloat64 {
				return ast.NewBoolean(lhs.Real() <= rhs.Real())
			} else {
				return ast.NewBoolean(lhs.Int() <= rhs.Int())
			}
		default:
			return nil
		}
	case token.TokenGE:
		switch target {
		case ast.TypeBoolean:
			if lhs.Type() == ast.TypeFloat64 || rhs.Type() == ast.TypeFloat64 {
				return ast.NewBoolean(lhs.Real() >= rhs.Real())
			} else {
				return ast.NewBoolean(lhs.Int() >= rhs.Int())
			}
		default:
			return nil
		}
	case token.TokenGT:
		switch target {
		case ast.TypeBoolean:
			if lhs.Type() == ast.TypeFloat64 || rhs.Type() == ast.TypeFloat64 {
				return ast.NewBoolean(lhs.Real() > rhs.Real())
			} else {
				return ast.NewBoolean(lhs.Int() > rhs.Int())
			}
		default:
			return nil
		}
	case token.TokenIN:
		if lhs.Type() == ast.TypeInt64 && rhs.Type() == ast.TypeSet {
			return ast.NewBoolean(rhs.Set()&(1<<lhs.Int()) != 0)
		} else {
			return nil
		}
	case token.TokenIS:
		return nil
	default:
		return nil
	}
}
