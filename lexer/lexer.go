package lexer

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/corani/silver-octo-sniffle/reporter"
	"github.com/corani/silver-octo-sniffle/token"
)

var ErrLexer = errors.New("error during lexing")

func Lex(out *reporter.Reporter, name string, bs []byte) (token.Tokens, error) {
	var (
		result token.Tokens
		err    error
	)

	row := 1
	col := 1

	for i := 0; i < len(bs); {
		starti := i
		startr := row
		startc := col

		// NOTE(daniel): some helper functions:

		// increment the cursor.
		accept := func() {
			if i < len(bs) && bs[i] == '\n' {
				row++
				col = 0
			} else {
				col++
			}

			i++
		}

		// if the current char matches, increment the cursor.
		peek := func(c byte) bool {
			if i < len(bs) && bs[i] == c {
				accept()

				return true
			}

			return false
		}

		// return the accumulated text for the current token.
		text := func() string {
			return string(bs[starti:i])
		}

		while := func(fn func(byte) bool) {
			for fn(bs[i]) {
				accept()
			}
		}

		report := func(format string, args ...any) {
			err = ErrLexer
			loc := token.Token{
				File:  name,
				Range: token.NewRange(startr, startc, row, col),
			}

			out.Errorf(loc, format, args...)
		}

		// NOTE(daniel): identify token.
		var tokenType token.TokenType

		switch {
		case isNumeric(bs[i]): // number
			while(isHex)

			if strings.ContainsAny(text(), "ABCDEF") {
				// definitely hex digits, so they must be prefixed with 'H' or 'X'.
				switch {
				case peek('H'):
					tokenType = token.TokenInteger
				case peek('X'):
					tokenType = token.TokenString
				default:
					report("expected 'H' or 'X' after hex digits: %q", text())
					continue
				}
			} else if peek('.') {
				// special case: back-track as we found a ".." after the integer, so it's not a real.
				if peek('.') {
					tokenType = token.TokenInteger
					i -= 2

					break
				}

				// it's a real number.
				while(isNumeric)

				// NOTE(daniel): scan exponent.
				if peek('E') {
					if !peek('+') && !peek('-') {
						report("expected '+' or '-' after real exponent: %q", text())
						continue
					}

					while(isNumeric)
				}

				tokenType = token.TokenReal
			} else if peek('X') {
				tokenType = token.TokenString
			} else {
				// otherwise assume it's an integer number.
				peek('H') // accept optional 'H' qualifier.

				tokenType = token.TokenInteger
			}
		case isAlpha(bs[i]): // identifier or keyword
			while(isAlphaNumeric)

			tokenType = token.CheckKeyword(text())
		case peek('"'): // string
			// TODO(daniel): handle escape characters?
			// TODO(daniel): this accepts multi-line strings. Are those a thing?
			for !peek('"') {
				accept()
			}

			tokenType = token.TokenString
		case peek('.'): // dot or range
			if peek('.') {
				tokenType = token.TokenDotDot
			} else {
				tokenType = token.TokenDot
			}
		case peek('<') || peek('>'): // relationship
			peek('=') // accept following '='

			switch text() {
			case "<":
				tokenType = token.TokenLT
			case "<=":
				tokenType = token.TokenLE
			case ">=":
				tokenType = token.TokenGE
			case ">":
				tokenType = token.TokenGT
			}
		case peek('('): // parenthesis or comment
			// TODO(daniel): do we want to store comments in the token stream, either as
			// separate tokens or as annotations on e.g. the next token? Storing them would
			// allow us to use them in the parser or generator, e.g. for compiler directives.
			// But on the other hand, we'd have to thread them through the whole pipeline.
			if peek('*') {
				// TODO(daniel): handle nested comments?
				for !peek(')') {
					for !peek('*') {
						accept()
					}
				}

				continue
			} else {
				tokenType = token.TokenLParen
			}
		case peek(':'): // colon or assignment
			if peek('=') {
				tokenType = token.TokenAssign
			} else {
				tokenType = token.TokenColon
			}
		case peek('\n') || peek('\r') || peek('\t') || peek(' '): // whitespace
			continue
		default:
			if t, ok := token.CheckChar(bs[i]); ok {
				accept()

				tokenType = t
			} else {
				report("unknown character '%c'", bs[i])
				accept()

				continue
			}
		}

		// NOTE(daniel): add the identified token.
		token, err1 := makeToken(text(), tokenType, name, token.NewRange(startr, startc, row, col))
		if err1 != nil {
			report("%v", err1)
			continue
		}

		result = append(result, token)
	}

	// NOTE(daniel): add EOF token as a terminator.
	result = append(result, makeEOF(name, row, col))

	return result, err
}

func decodeInteger(txt string) (int, error) {
	base := 10

	// NOTE(daniel): decode hex digits.
	if strings.HasSuffix(txt, "H") {
		base = 16
		txt = strings.TrimSuffix(txt, "H")
	}

	if i, err := strconv.ParseInt(txt, base, 64); err != nil {
		return 0, err
	} else {
		return int(i), nil
	}
}

func decodeReal(txt string) (float64, error) {
	return strconv.ParseFloat(txt, 64)
}

func decodeString(txt string) (string, error) {
	if strings.HasPrefix(txt, "\"") {
		// slice off the enclosing quotes.
		txt = strings.TrimPrefix(txt, "\"")
		txt = strings.TrimSuffix(txt, "\"")

		return txt, nil
	} else if strings.HasSuffix(txt, "X") {
		// decode hex string
		txt = strings.TrimSuffix(txt, "X")

		bs, err := hex.DecodeString(txt)
		if err != nil {
			return "", err
		}

		return string(bs), nil
	} else {
		return "", fmt.Errorf("read invalid string: %q", txt)
	}
}

func makeToken(txt string, tokenType token.TokenType, name string, loc token.Range) (token.Token, error) {
	var (
		ival int
		fval float64
		bval bool
	)

	switch tokenType {
	case token.TokenInvalid:
		return token.Token{}, fmt.Errorf("read invalid token")
	case token.TokenInteger:
		i, err := decodeInteger(txt)
		if err != nil {
			return token.Token{}, err
		}

		ival = i
	case token.TokenReal:
		f, err := decodeReal(txt)
		if err != nil {
			return token.Token{}, err
		}

		fval = f
	case token.TokenString:
		t, err := decodeString(txt)
		if err != nil {
			return token.Token{}, err
		}

		txt = t
	case token.TokenBoolean:
		if txt == "TRUE" {
			bval = true
		}
	}

	return token.Token{
		Type:  tokenType,
		File:  name,
		Range: loc,
		Text:  txt,
		Bool:  bval,
		Int:   ival,
		Real:  fval,
	}, nil
}

func makeEOF(name string, row, col int) token.Token {
	return token.Token{
		Type:  token.TokenEOF,
		File:  name,
		Range: token.NewRange(row, col),
		Text:  "",
		Int:   0,
	}
}

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func isNumeric(b byte) bool {
	return (b >= '0' && b <= '9')
}

func isHex(b byte) bool {
	return isNumeric(b) || (b >= 'A' && b <= 'F')
}

func isAlphaNumeric(b byte) bool {
	return isAlpha(b) || isNumeric(b)
}
