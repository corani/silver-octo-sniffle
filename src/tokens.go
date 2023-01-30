package main

import (
	"fmt"
	"strings"
)

type TokenType string

const (
	TokenInvalid   TokenType = ""
	TokenNumber    TokenType = "number"
	TokenString    TokenType = "string"
	TokenBoolean   TokenType = "boolean"
	TokenLParen    TokenType = "lparen"
	TokenRParen    TokenType = "rparen"
	TokenPlus      TokenType = "plus"
	TokenMinus     TokenType = "minus"
	TokenStar      TokenType = "star"
	TokenSlash     TokenType = "slash"
	TokenComma     TokenType = "comma"
	TokenAmpersand TokenType = "ampersand"
	TokenIdent     TokenType = "ident"
	TokenDiv       TokenType = "div"
	TokenMod       TokenType = "mod"
	TokenOr        TokenType = "or"
	TokenEOF       TokenType = "eof"
)

var mapCharToToken = map[byte]TokenType{
	'(': TokenLParen,
	')': TokenRParen,
	'+': TokenPlus,
	'-': TokenMinus,
	'*': TokenStar,
	'/': TokenSlash, // REAL division
	',': TokenComma,
	'&': TokenAmpersand, // logical AND
}

// TODO(daniel): are keywords case-sensitive (i.e. must they be all upper-case)?
var mapIdentToToken = map[string]TokenType{
	"DIV":   TokenDiv, // INTEGER division
	"MOD":   TokenMod,
	"OR":    TokenOr,
	"TRUE":  TokenBoolean,
	"FALSE": TokenBoolean,
}

func checkKeyword(s string) TokenType {
	if v, ok := mapIdentToToken[s]; ok {
		return v
	}

	return TokenIdent
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

	parts = append(parts, "## Tokens", "```tsv")

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
