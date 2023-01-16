package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TokenType string

const (
	TokenNumber TokenType = "number"
	TokenIdent  TokenType = "ident"
	TokenPrint  TokenType = "print"
	TokenLParen TokenType = "lparen"
	TokenRParen TokenType = "rparen"
	TokenPlus   TokenType = "add"
	TokenMinus  TokenType = "sub"
	TokenComma  TokenType = "comma"
	TokenEOF    TokenType = "eof"
)

var mapCharToToken = map[byte]TokenType{
	'(': TokenLParen,
	')': TokenRParen,
	'+': TokenPlus,
	'-': TokenMinus,
	',': TokenComma,
}

type Range struct {
	FromRow, FromCol int
	ToRow, ToCol     int
}

func (r Range) String() string {
	return fmt.Sprintf("(%d, %d) -> (%d, %d)", r.FromRow, r.FromCol, r.ToRow, r.ToCol)
}

type Tokens []Token

func (t Tokens) String() string {
	var parts []string

	parts = append(parts, "Tokens:")

	for _, v := range t {
		parts = append(parts, v.String())
	}

	return strings.Join(parts, "\n")
}

type Token struct {
	Type   TokenType
	File   string
	Range  Range
	Text   string
	Number int
}

func (t Token) String() string {
	return fmt.Sprintf("%s:%d:%d:\t%v\t%q\t%d\t%v", t.File, t.Range.FromRow, t.Range.FromCol, t.Type, t.Text, t.Number, t.Range)
}

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

			tokenType = TokenIdent

			// TODO(daniel): move this to semantic analysis?
			switch string(bs[starti:i]) {
			case "print":
				tokenType = TokenPrint
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
