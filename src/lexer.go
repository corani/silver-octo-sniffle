package main

import (
	"fmt"
	"strconv"
)

func lex(name string, bs []byte) (Tokens, error) {
	var result Tokens

	row := 1
	col := 1

	for i := 0; i < len(bs); {
		starti := i
		startr := row
		startc := col

		var tokenType TokenType

		switch {
		case numeric(bs[i]):
			// TODO(daniel): support REAL numbers.
			for numeric(bs[i]) {
				i++
				col++
			}

			tokenType = TokenNumber
		case alpha(bs[i]):
			for alphanumeric(bs[i]) {
				i++
				col++
			}

			if v, ok := mapIdentToToken[string(bs[starti:i])]; ok {
				tokenType = v
			} else {
				tokenType = TokenIdent
			}
		default:
			if t, ok := mapCharToToken[bs[i]]; ok {
				i++
				col++
				tokenType = t
			} else {
				switch bs[i] {
				case '\n':
					i++
					row++
					col = 0
					continue
				case '\r', ' ', '\t':
					i++
					col++
					continue
				default:
					return result, fmt.Errorf("unknown character '%c'", bs[i])
				}
			}
		}

		// NOTE(daniel): add the identified token.
		text := string(bs[starti:i])
		var num int

		if tokenType == TokenNumber {
			if i, err := strconv.Atoi(text); err != nil {
				return result, err
			} else {
				num = i
			}
		}

		result = append(result, Token{
			Type:   tokenType,
			File:   name,
			Range:  Range{startr, startc, row, col},
			Text:   text,
			Number: num,
		})
	}

	// NOTE(daniel): add EOF token.
	result = append(result, Token{
		Type:   TokenEOF,
		File:   name,
		Range:  Range{row, col, row, col},
		Text:   "",
		Number: 0,
	})

	return result, nil
}

func alpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func numeric(b byte) bool {
	return (b >= '0' && b <= '9')
}

func alphanumeric(b byte) bool {
	return alpha(b) || numeric(b)
}
