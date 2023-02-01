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
	TokenNIL       TokenType = "nil"
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
	TokenAmpersand TokenType = "ampersand"
	TokenTilde     TokenType = "tilde"
	TokenColon     TokenType = "colon"
	TokenSemicolon TokenType = "semicolon"
	TokenDot       TokenType = "dot"
	TokenDotDot    TokenType = "dotdot"
	TokenCaret     TokenType = "caret"
	TokenBar       TokenType = "bar"
	TokenIdent     TokenType = "ident"
	TokenAssign    TokenType = "assign"
	TokenDIV       TokenType = "div"
	TokenMOD       TokenType = "mod"
	TokenOR        TokenType = "or"
	TokenEQ        TokenType = "eq"
	TokenNE        TokenType = "ne"
	TokenLT        TokenType = "lt"
	TokenLE        TokenType = "le"
	TokenGE        TokenType = "ge"
	TokenGT        TokenType = "gt"
	TokenIN        TokenType = "in"
	TokenIS        TokenType = "is"
	TokenMODULE    TokenType = "module"
	TokenBEGIN     TokenType = "begin"
	TokenEND       TokenType = "end"
	TokenIMPORT    TokenType = "import"
	TokenCONST     TokenType = "const"
	TokenTYPE      TokenType = "type"
	TokenVAR       TokenType = "var"
	TokenARRAY     TokenType = "array"
	TokenOF        TokenType = "of"
	TokenRECORD    TokenType = "record"
	TokenPOINTER   TokenType = "pointer"
	TokenTO        TokenType = "to"
	TokenPROCEDURE TokenType = "procedure"
	TokenRETURN    TokenType = "return"
	TokenIF        TokenType = "if"
	TokenTHEN      TokenType = "then"
	TokenELSIF     TokenType = "elsif"
	TokenELSE      TokenType = "else"
	TokenCASE      TokenType = "case"
	TokenWHILE     TokenType = "while"
	TokenDO        TokenType = "do"
	TokenREPEAT    TokenType = "repeat"
	TokenUNTIL     TokenType = "until"
	TokenFOR       TokenType = "for"
	TokenBY        TokenType = "by"
	TokenEOF       TokenType = "eof"
)

var mapCharToToken = map[byte]TokenType{
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
	'&': TokenAmpersand, // logical AND
	'~': TokenTilde,     // logical NOT
	';': TokenSemicolon,
	'=': TokenEQ,
	'#': TokenNE,
	'^': TokenCaret, // dereference
	'|': TokenBar,
}

// TODO(daniel): are keywords case-sensitive (i.e. must they be all upper-case)?
var mapIdentToToken = map[string]TokenType{
	"DIV":       TokenDIV, // INTEGER division
	"MOD":       TokenMOD,
	"OR":        TokenOR,
	"TRUE":      TokenBoolean,
	"FALSE":     TokenBoolean,
	"NIL":       TokenNIL,
	"IN":        TokenIN,
	"IS":        TokenIS,
	"MODULE":    TokenMODULE,
	"BEGIN":     TokenBEGIN,
	"END":       TokenEND,
	"IMPORT":    TokenIMPORT,
	"CONST":     TokenCONST,
	"TYPE":      TokenTYPE,
	"VAR":       TokenVAR,
	"ARRAY":     TokenARRAY,
	"OF":        TokenOF,
	"RECORD":    TokenRECORD,
	"POINTER":   TokenPOINTER,
	"TO":        TokenTO,
	"PROCEDURE": TokenPROCEDURE,
	"RETURN":    TokenRETURN,
	"IF":        TokenIF,
	"THEN":      TokenTHEN,
	"ELSIF":     TokenELSIF,
	"ELSE":      TokenELSE,
	"CASE":      TokenCASE,
	"WHILE":     TokenWHILE,
	"DO":        TokenDO,
	"REPEAT":    TokenREPEAT,
	"UNTIL":     TokenUNTIL,
	"FOR":       TokenFOR,
	"BY":        TokenBY,
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
	Bool  bool
	Int   int
	Real  float64
}

func (t Token) String() string {
	return fmt.Sprintf("%s:%d:%d:\t%v\t%q\t%v\t%d\t%f\t%v", t.File, t.Range.FromRow, t.Range.FromCol, t.Type, t.Text, t.Bool, t.Int, t.Real, t.Range)
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
