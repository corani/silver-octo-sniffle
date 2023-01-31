package main

import (
	"fmt"
	"strings"
)

type TokenType string

const (
	TokenInvalid   TokenType = ""
	TokenInteger   TokenType = "integer"
	TokenReal      TokenType = "real"
	TokenString    TokenType = "string"
	TokenBoolean   TokenType = "boolean"
	TokenNil       TokenType = "nil"
	TokenLParen    TokenType = "lparen"
	TokenRParen    TokenType = "rparen"
	TokenLBrace    TokenType = "lbrace"
	TokenRBrace    TokenType = "rbrace"
	TokenLBracket  TokenType = "lbracket"
	TokenRBracket  TokenType = "rbracket"
	TokenPlus      TokenType = "plus"
	TokenMinus     TokenType = "minus"
	TokenAsterisk  TokenType = "asterisk"
	TokenSlash     TokenType = "slash"
	TokenComma     TokenType = "comma"
	TokenAnd       TokenType = "and"
	TokenNot       TokenType = "not"
	TokenColon     TokenType = "colon"
	TokenSemicolon TokenType = "semicolon"
	TokenDot       TokenType = "dot"
	TokenDotDot    TokenType = "dotdot"
	TokenCaret     TokenType = "caret"
	TokenBar       TokenType = "bar"
	TokenIdent     TokenType = "ident"
	TokenIDiv      TokenType = "idiv"
	TokenMod       TokenType = "mod"
	TokenOr        TokenType = "or"
	TokenEQ        TokenType = "eq"
	TokenNE        TokenType = "ne"
	TokenLT        TokenType = "lt"
	TokenLE        TokenType = "le"
	TokenGE        TokenType = "ge"
	TokenGT        TokenType = "gt"
	TokenIN        TokenType = "in"
	TokenIS        TokenType = "is"
	TokenEOF       TokenType = "eof"
)

var mapCharToToken = map[byte]TokenType{
	'(': TokenLParen,
	')': TokenRParen,
	'{': TokenLBrace,
	'}': TokenRBrace,
	'[': TokenLBracket,
	']': TokenRBracket,
	'+': TokenPlus,     // addition, set union
	'-': TokenMinus,    // subtraction, set difference
	'*': TokenAsterisk, // multiplication, set intersection
	'/': TokenSlash,    // REAL division, symmetric set difference
	',': TokenComma,
	'&': TokenAnd, // logical AND
	'~': TokenNot, // logical NOT
	':': TokenColon,
	';': TokenSemicolon,
	'=': TokenEQ,
	'#': TokenNE,
	'^': TokenCaret, // dereference
	'|': TokenBar,
}

// TODO(daniel): are keywords case-sensitive (i.e. must they be all upper-case)?
var mapIdentToToken = map[string]TokenType{
	"DIV":   TokenIDiv, // INTEGER division
	"MOD":   TokenMod,
	"OR":    TokenOr,
	"TRUE":  TokenBoolean,
	"FALSE": TokenBoolean,
	"NIL":   TokenNil,
	"IN":    TokenIN,
	"IS":    TokenIS,
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
	Type  TokenType
	File  string
	Range Range
	Text  string
	Int   int
	Real  float64
}

func (t Token) String() string {
	return fmt.Sprintf("%s:%d:%d:\t%v\t%q\t%d\t%v", t.File, t.Range.FromRow, t.Range.FromCol, t.Type, t.Text, t.Int, t.Range)
}

func (t Token) isRelation() bool {
	relations := []TokenType{TokenEQ, TokenNE, TokenLT, TokenLE, TokenGE, TokenGT, TokenIN, TokenIS}

	for _, v := range relations {
		if t.Type == v {
			return true
		}
	}

	return false
}

func (t Token) isNumber() bool {
	return t.Type == TokenInteger || t.Type == TokenReal
}
