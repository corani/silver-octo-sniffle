package main

import (
	"fmt"
	"strings"
)

type TokenType string

const (
	TokenInvalid TokenType = ""
	TokenNumber  TokenType = "number"
	TokenIdent   TokenType = "ident"
	TokenLParen  TokenType = "lparen"
	TokenRParen  TokenType = "rparen"
	TokenPlus    TokenType = "add"
	TokenMinus   TokenType = "sub"
	TokenStar    TokenType = "mul"
	TokenSlash   TokenType = "div"
	TokenComma   TokenType = "comma"
	TokenEOF     TokenType = "eof"
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

	parts = append(parts, "## Tokens", "```")

	for _, v := range t {
		parts = append(parts, v.String())
	}

	parts = append(parts, "```")

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
