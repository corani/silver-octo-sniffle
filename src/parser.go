package main

import "fmt"

type Visitor interface {
	VisitNumberExpr(*NumberExpr)
	VisitBinaryExpr(*BinaryExpr)
	VisitPrintStmt(*PrintStmt)
}

type Node interface {
	Token() Token
	Visit(Visitor)
}

type NumberExpr struct {
	token Token
}

func (n *NumberExpr) Token() Token {
	return n.token
}

func (n *NumberExpr) Visit(v Visitor) {
	v.VisitNumberExpr(n)
}

type BinaryExpr struct {
	token Token
	args  []Node
}

func (n *BinaryExpr) Token() Token {
	return n.token
}

func (n *BinaryExpr) Visit(v Visitor) {
	v.VisitBinaryExpr(n)
}

type PrintStmt struct {
	token Token
	args  []Node
}

func (n *PrintStmt) Token() Token {
	return n.token
}

func (n *PrintStmt) Visit(v Visitor) {
	v.VisitPrintStmt(n)
}

type Parser struct {
	tokens Tokens
	index  int
}

func parse(tokens Tokens) (Node, error) {
	p := &Parser{
		tokens: tokens,
		index:  0,
	}

	return p.parse()
}

func (p *Parser) parse() (Node, error) {
	switch p.currentType() {
	case TokenPrint:
		return p.parsePrintStmt()
	default:
		return nil, fmt.Errorf("unexpected token: %v", p.tokens[p.index])
	}
}

func (p *Parser) parsePrintStmt() (Node, error) {
	node := &PrintStmt{
		token: p.currentToken(),
		args:  nil,
	}

	p.index++

	if _, err := p.require(TokenLParen); err != nil {
		return node, err
	}

	for {
		expr, err := p.parseExpr()
		if err != nil {
			return nil, err
		}

		node.args = append(node.args, expr)

		if !p.expect(TokenComma) {
			break
		}
	}

	if _, err := p.require(TokenRParen); err != nil {
		return node, err
	}

	return node, nil
}

func (p *Parser) parseExpr() (Node, error) {
	// expr := term { '+' | '-' term }
	lhs, err := p.parseTerm()
	if err != nil {
		return lhs, err
	}

	// NOTE(daniel): parse a chain of `+`/`-` operators
	for op := p.currentToken(); op.Type == TokenPlus || op.Type == TokenMinus; op = p.currentToken() {
		p.index++

		rhs, err := p.parseTerm()
		if err != nil {
			return rhs, err
		}

		lhs = &BinaryExpr{
			token: op,
			args:  []Node{lhs, rhs},
		}
	}

	return lhs, nil
}

func (p *Parser) parseTerm() (Node, error) {
	// term := factor { '*' | '/' factor }
	lhs, err := p.parseFactor()
	if err != nil {
		return lhs, err
	}

	// NOTE(daniel): parse a chain of `*`/`/` operators
	for op := p.currentToken(); op.Type == TokenStar || op.Type == TokenSlash; op = p.currentToken() {
		p.index++

		rhs, err := p.parseFactor()
		if err != nil {
			return rhs, err
		}

		lhs = &BinaryExpr{
			token: op,
			args:  []Node{lhs, rhs},
		}
	}

	return lhs, nil
}

func (p *Parser) parseFactor() (Node, error) {
	// factor := '(' expr ')' | number
	if p.currentType() != TokenLParen {
		return p.parseNumberExpr()
	}

	if _, err := p.require(TokenLParen); err != nil {
		return nil, err
	}

	expr, err := p.parseExpr()
	if err != nil {
		return expr, err
	}

	if _, err := p.require(TokenRParen); err != nil {
		return nil, err
	}

	return expr, nil
}

func (p *Parser) parseNumberExpr() (Node, error) {
	var op Token

	switch p.currentType() {
	case TokenMinus:
		// Unary Minus.
		op, _ = p.require(TokenMinus)
	case TokenPlus:
		// Unary Plus is a no-op so we ignore it.
		_, _ = p.require(TokenPlus)
	}

	t, err := p.require(TokenNumber)
	if err != nil {
		return nil, err
	}

	num := &NumberExpr{token: t}

	if op.Type == TokenInvalid {
		return num, nil
	}

	// NOTE(daniel): There is no negate for integers in LLVM IR, so we treat
	// `-x` as `0 - x`. That means we don't need generator code either.
	zero := &NumberExpr{token: t}
	zero.token.Number = 0
	zero.token.Text = ""

	return &BinaryExpr{
		token: op,
		args:  []Node{zero, num},
	}, nil
}

func (p *Parser) require(exp TokenType) (Token, error) {
	if p.currentType() == exp {
		token := p.currentToken()
		p.index++

		return token, nil
	}

	return Token{}, fmt.Errorf("unexpected token: %v (expected %v)", p.tokens[p.index], exp)
}

func (p *Parser) expect(exp TokenType) bool {
	if p.currentType() == exp {
		p.index++

		return true
	}

	return false
}

func (p *Parser) currentToken() Token {
	return p.tokens[p.index]
}

func (p *Parser) currentType() TokenType {
	return p.tokens[p.index].Type
}
