package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func lex(name string, bs []byte) (Tokens, error) {
	var result Tokens

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

		// NOTE(daniel): identify token.
		var tokenType TokenType

		switch {
		case isNumeric(bs[i]): // number
			while(isHex)

			if strings.ContainsAny(text(), "ABCDEF") {
				// definitely hex digits, so they must be prefixed with 'H' or 'X'.
				switch {
				case peek('H'):
					tokenType = TokenInteger
				case peek('X'):
					tokenType = TokenString
				default:
					return result, fmt.Errorf("expected 'H' or 'X' after hex digits: %q", text())
				}
			} else if peek('.') {
				// it's a real number.
				while(isNumeric)

				// NOTE(daniel): scan exponent.
				if peek('E') {
					if !peek('+') && !peek('-') {
						return result, fmt.Errorf("expected '+' or '-' after real exponent: %q", text())
					}

					while(isNumeric)
				}

				tokenType = TokenReal
			} else if peek('X') {
				tokenType = TokenString
			} else {
				// otherwise assume it's an integer number.
				peek('H') // accept optional 'H' qualifier.

				tokenType = TokenInteger
			}
		case isAlpha(bs[i]): // identifier or keyword
			while(isAlphaNumeric)

			tokenType = checkKeyword(text())
		case peek('"'): // string
			// TODO(daniel): handle escape characters?
			// TODO(daniel): this accepts multi-line strings. Are those a thing?
			for !peek('"') {
				accept()
			}

			tokenType = TokenString
		case peek('.'): // dot or range
			if peek('.') {
				tokenType = TokenDotDot
			} else {
				tokenType = TokenDot
			}
		case peek('<') || peek('>'): // relationship
			peek('=') // accept following '='

			switch text() {
			case "<":
				tokenType = TokenLT
			case "<=":
				tokenType = TokenLE
			case ">=":
				tokenType = TokenGE
			case ">":
				tokenType = TokenGT
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
				tokenType = TokenLParen
			}
		case peek(':'): // colon or assignment
			if peek('=') {
				tokenType = TokenAssign
			} else {
				tokenType = TokenColon
			}
		case peek('\n') || peek('\r') || peek('\t') || peek(' '): // whitespace
			continue
		default:
			if t, ok := mapCharToToken[bs[i]]; ok {
				accept()

				tokenType = t
			} else {
				return result, fmt.Errorf("unknown character '%c'", bs[i])
			}
		}

		// NOTE(daniel): add the identified token.
		token, err := makeToken(text(), tokenType, name, Range{startr, startc, row, col})
		if err != nil {
			return result, fmt.Errorf("%s:%d:%d: %w", name, startr, startc, err)
		}

		result = append(result, token)
	}

	// NOTE(daniel): add EOF token as a terminator.
	result = append(result, makeEOF(name, row, col))

	return result, nil
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

func makeToken(txt string, tokenType TokenType, name string, loc Range) (Token, error) {
	var (
		ival int
		fval float64
		bval bool
	)

	switch tokenType {
	case TokenInvalid:
		return Token{}, fmt.Errorf("read invalid token")
	case TokenInteger:
		i, err := decodeInteger(txt)
		if err != nil {
			return Token{}, err
		}

		ival = i
	case TokenReal:
		f, err := decodeReal(txt)
		if err != nil {
			return Token{}, err
		}

		fval = f
	case TokenString:
		t, err := decodeString(txt)
		if err != nil {
			return Token{}, err
		}

		txt = t
	case TokenBoolean:
		if txt == "TRUE" {
			bval = true
		}
	}

	return Token{
		Type:  tokenType,
		File:  name,
		Range: loc,
		Text:  txt,
		Bool:  bval,
		Int:   ival,
		Real:  fval,
	}, nil
}

func makeEOF(name string, row, col int) Token {
	return Token{
		Type:  TokenEOF,
		File:  name,
		Range: Range{row, col, row, col},
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
