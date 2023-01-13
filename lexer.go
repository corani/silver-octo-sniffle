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
	TokenPlus   TokenType = "plus"
	TokenEOF    TokenType = "eof"
)

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

		switch {
		case bs[i] >= '0' && bs[i] <= '9':
			for bs[i] >= '0' && bs[i] <= '9' {
				i++
				col++
			}

			text := string(bs[starti:i])
			num, err := strconv.Atoi(text)
			if err != nil {
				return result, err
			}

			result = append(result, Token{
				Type:   TokenNumber,
				File:   name,
				Range:  Range{startr, startc, row, col},
				Text:   text,
				Number: num,
			})
		case bs[i] >= 'a' && bs[i] <= 'z':
			for (bs[i] >= '0' && bs[i] <= '9') || (bs[i] >= 'a' && bs[i] <= 'z') {
				i++
				col++
			}

			text := string(bs[starti:i])
			typ := TokenIdent

			switch text {
			case "print":
				typ = TokenPrint
			}

			result = append(result, Token{
				Type:  typ,
				File:  name,
				Range: Range{startr, startc, row, col},
				Text:  text,
			})
		default:
			switch bs[i] {
			case '(':
				i++
				col++

				result = append(result, Token{
					Type:   TokenLParen,
					File:   name,
					Range:  Range{startr, startc, row, col},
					Text:   string(bs[starti:i]),
					Number: 0,
				})
			case ')':
				i++
				col++

				result = append(result, Token{
					Type:   TokenRParen,
					File:   name,
					Range:  Range{startr, startc, row, col},
					Text:   string(bs[starti:i]),
					Number: 0,
				})
			case '+':
				i++
				col++

				result = append(result, Token{
					Type:   TokenPlus,
					File:   name,
					Range:  Range{startr, startc, row, col},
					Text:   string(bs[starti:i]),
					Number: 0,
				})
			case '\n':
				i++
				row++
				col = 0
			case '\r', ' ', '\t':
				i++
			default:
				return result, fmt.Errorf("unknown character '%c'", bs[i])
			}
		}
	}

	result = append(result, Token{
		Type:   TokenEOF,
		File:   name,
		Range:  Range{row, col, row, col},
		Text:   "",
		Number: 0,
	})

	return result, nil
}
