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
		case numeric(bs[i]):
			// TODO(daniel): support REAL numbers.
			for numeric(bs[i]) {
				next()
			}

			tokenType = TokenNumber
		case alpha(bs[i]):
			for alphanumeric(bs[i]) {
				next()
			}

			tokenType = checkKeyword(text())
		case bs[i] == '(':
			// NOTE(daniel): special case, to handle skipping over comments.
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
		var num int

		if tokenType == TokenNumber {
			if i, err := strconv.Atoi(text()); err != nil {
				return result, err
			} else {
				num = i
			}
		}

		result = append(result, Token{
			Type:   tokenType,
			File:   name,
			Range:  Range{startr, startc, row, col},
			Text:   text(),
			Number: num,
		})
	}

	// NOTE(daniel): add EOF token as a terminator.
	result = append(result, eof(name, row, col))

	return result, nil
}

func eof(name string, row, col int) Token {
	return Token{
		Type:   TokenEOF,
		File:   name,
		Range:  Range{row, col, row, col},
		Text:   "",
		Number: 0,
	}
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
