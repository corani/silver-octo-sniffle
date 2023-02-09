package check

import (
	"fmt"

	"github.com/corani/silver-octo-sniffle/ast"
	"github.com/corani/silver-octo-sniffle/lex"
)

func evaluateNumberExpr(lit lex.Token) (*ast.Value, error) {
	switch lit.Type {
	case lex.TokenInteger:
		return ast.NewInt(lit.Int), nil
	case lex.TokenReal:
		return ast.NewReal(lit.Real), nil
	default:
		return nil, fmt.Errorf("invalid number literal: %q", lit.Text)
	}
}

func evaluateCharExpr(lit lex.Token) *ast.Value {
	if len(lit.Text) == 0 {
		return nil
	}

	return ast.NewChar(byte(lit.Text[0]))
}

func evaluateStringExpr(lit lex.Token) *ast.Value {
	return ast.NewString(lit.Text)
}

func evaluateBooleanExpr(lit lex.Token) (*ast.Value, error) {
	var v bool

	switch lit.Text {
	case "TRUE":
		v = true
	case "FALSE":
		v = false
	default:
		return nil, fmt.Errorf("invalid boolean literal: %q", lit.Text)
	}

	return ast.NewBoolean(v), nil
}

func evaluateSetExpr(expr *ast.SetExpr) *ast.Value {
	var val int64

	for _, bit := range expr.Bits() {
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

func evaluateBinaryExpr(op lex.TokenType, target ast.Type, lhs, rhs *ast.Value) *ast.Value {
	switch op {
	case lex.TokenPlus:
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
	case lex.TokenMinus:
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
	case lex.TokenAsterisk:
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
	case lex.TokenSlash:
		switch target {
		case ast.TypeFloat64: // REAL division
			return ast.NewReal(lhs.Real() / rhs.Real())
		case ast.TypeSet: // symmetric set difference
			return ast.NewSet(lhs.Set() ^ rhs.Set())
		default:
			return nil
		}
	case lex.TokenDIV: // INTEGER division
		switch target {
		case ast.TypeInt64:
			return ast.NewInt(lhs.Int() / rhs.Int())
		default:
			return nil
		}
	case lex.TokenMOD:
		switch target {
		case ast.TypeInt64:
			return ast.NewInt(lhs.Int() % rhs.Int())
		default:
			return nil
		}
	case lex.TokenOR:
		switch target {
		case ast.TypeBoolean:
			return ast.NewBoolean(lhs.Bool() || rhs.Bool())
		default:
			return nil
		}
	case lex.TokenAmpersand:
		switch target {
		case ast.TypeBoolean:
			return ast.NewBoolean(lhs.Bool() && rhs.Bool())
		default:
			return nil
		}
	case lex.TokenEQ:
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
	case lex.TokenNE:
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
	case lex.TokenLT:
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
	case lex.TokenLE:
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
	case lex.TokenGE:
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
	case lex.TokenGT:
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
	case lex.TokenIN:
		if lhs.Type() == ast.TypeInt64 && rhs.Type() == ast.TypeSet {
			return ast.NewBoolean(rhs.Set()&(1<<lhs.Int()) != 0)
		} else {
			return nil
		}
	case lex.TokenIS:
		return nil
	default:
		return nil
	}
}
