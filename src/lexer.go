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

		// increment the cursor.
		next := func() {
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
				next()

				return true
			}

			return false
		}

		// return the accumulated text for the current token.
		text := func() string {
			return string(bs[starti:i])
		}

		var tokenType TokenType

		switch {
		case isNumeric(bs[i]):
			for isHex(bs[i]) {
				next()
			}

			if strings.ContainsAny(text(), "ABCDEF") {
				switch {
				case peek('H'):
					tokenType = TokenInteger
				case peek('X'):
					tokenType = TokenString
				default:
					return result, fmt.Errorf("expected 'H' or 'X' after hex digits: %q", text())
				}
			} else if peek('.') {
				for isNumeric(bs[i]) {
					next()
				}

				// TODO(daniel): support scale factor.
				tokenType = TokenReal
			} else if peek('X') {
				tokenType = TokenString
			} else {
				peek('H')

				tokenType = TokenInteger
			}
		case isAlpha(bs[i]):
			for isAlphaNumeric(bs[i]) {
				next()
			}

			tokenType = checkKeyword(text())
		case bs[i] == '"':
			// NOTE(daniel): special case for strings.
			next()

			// TODO(daniel): handle escape characters?
			// TODO(daniel): this accepts multi-line strings. Are those a thing?
			for !peek('"') {
				next()
			}

			tokenType = TokenString
		case bs[i] == '.':
			// NOTE(daniel): special case for ranges.
			next()

			if peek('.') {
				tokenType = TokenDotDot
			} else {
				tokenType = TokenDot
			}
		case bs[i] == '<' || bs[i] == '>':
			// NOTE(daniel): special case for relationships.
			next()
			peek('=')

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
		case bs[i] == '(':
			// NOTE(daniel): special case for comments.
			// TODO(daniel): do we want to store comments in the token stream, either as
			// separate tokens or as annotations on e.g. the next token? Storing them would
			// allow us to use them in the parser or generator, e.g. for compiler directives.
			// But on the other hand, we'd have to thread them through the whole pipeline.
			next()

			if peek('*') {
				// TODO(daniel): handle nested comments?
				for !peek(')') {
					for !peek('*') {
						next()
					}
				}

				continue
			} else {
				tokenType = TokenLParen
			}
		default:
			if t, ok := mapCharToToken[bs[i]]; ok {
				next()

				tokenType = t
			} else {
				// skip whitespace.
				switch bs[i] {
				case '\n', '\r', '\t', ' ':
					next()

					continue
				default:
					return result, fmt.Errorf("unknown character '%c'", bs[i])
				}
			}
		}

		// NOTE(daniel): add the identified token.
		var (
			txt  string
			inum int
			fnum float64
		)

		switch tokenType {
		case TokenInvalid:
			return result, fmt.Errorf("%s:%d:%d: read invalid token", name, startr, startc)
		case TokenInteger:
			txt = text()

			i, err := decodeInteger(txt)
			if err != nil {
				return result, err
			}

			inum = i
		case TokenReal:
			txt = text()

			f, err := decodeReal(txt)
			if err != nil {
				return result, err
			}

			fnum = f
		case TokenString:
			txt = text()

			t, err := decodeString(txt)
			if err != nil {
				return result, err
			}

			txt = t
		case TokenBoolean:
			txt = text()
			inum = 0

			if txt == "TRUE" {
				inum = 1
			}
		default:
			txt = text()
		}

		result = append(result, Token{
			Type:  tokenType,
			File:  name,
			Range: Range{startr, startc, row, col},
			Text:  txt,
			Int:   inum,
			Real:  fnum,
		})
	}

	// NOTE(daniel): add EOF token as a terminator.
	result = append(result, eof(name, row, col))

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

func eof(name string, row, col int) Token {
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
