package main

import "fmt"

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
	return p.parseModule()
}

func (p *Parser) parseModule() (Node, error) {
	node := &Module{}

	for p.currentType() != TokenEOF {
		stmt, err := p.parseStmt()
		if err != nil {
			return nil, err
		}

		node.stmts = append(node.stmts, stmt)
	}

	return node, nil
}

func (p *Parser) parseStmt() (Stmt, error) {
	switch p.currentType() {
	case TokenIdent:
		expr, err := p.parseCallExpr()
		if err != nil {
			return nil, err
		}

		return &ExprStmt{
			expr: expr,
		}, nil
	default:
		return nil, fmt.Errorf("unexpected token: %v", p.tokens[p.index])
	}
}

func (p *Parser) parseCallExpr() (Expr, error) {
	node := &CallExpr{
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

func (p *Parser) parseExpr() (Expr, error) {
	// expr := term { '+' | '-' | 'OR' term }
	addOperators := []TokenType{TokenPlus, TokenMinus, TokenOr}

	lhs, err := p.parseTerm()
	if err != nil {
		return lhs, err
	}

	for op := p.currentToken(); p.tokenIs(addOperators...); op = p.currentToken() {
		p.index++

		rhs, err := p.parseTerm()
		if err != nil {
			return rhs, err
		}

		lhs = &BinaryExpr{
			token: op,
			args:  []Expr{lhs, rhs},
		}
	}

	return lhs, nil
}

func (p *Parser) parseTerm() (Expr, error) {
	// term := factor { '*' | '/' | 'DIV' | 'MUL' | '&' factor }
	mulOperators := []TokenType{TokenStar, TokenSlash, TokenDiv, TokenMod, TokenAmpersand}

	lhs, err := p.parseFactor()
	if err != nil {
		return lhs, err
	}

	for op := p.currentToken(); p.tokenIs(mulOperators...); op = p.currentToken() {
		p.index++

		rhs, err := p.parseFactor()
		if err != nil {
			return rhs, err
		}

		lhs = &BinaryExpr{
			token: op,
			args:  []Expr{lhs, rhs},
		}
	}

	return lhs, nil
}

func (p *Parser) parseFactor() (Expr, error) {
	// factor := '(' expr ')' | number | string | boolean
	switch p.currentType() {
	case TokenNumber, TokenMinus:
		return p.parseNumberExpr()
	case TokenString:
		return p.parseStringExpr()
	case TokenBoolean:
		return p.parseBooleanExpr()
	case TokenLParen:
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
	default:
		panic(fmt.Sprintf("unexpected token in factor: %v", p.currentType()))
	}
}

func (p *Parser) parseNumberExpr() (Expr, error) {
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
		args:  []Expr{zero, num},
	}, nil
}

func (p *Parser) parseStringExpr() (Expr, error) {
	t, err := p.require(TokenString)
	if err != nil {
		return nil, err
	}

	return &StringExpr{token: t}, nil
}

func (p *Parser) parseBooleanExpr() (Expr, error) {
	t, err := p.require(TokenBoolean)
	if err != nil {
		return nil, err
	}

	return &BooleanExpr{token: t}, nil
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

func (p *Parser) tokenIs(opts ...TokenType) bool {
	current := p.currentType()

	for _, opt := range opts {
		if opt == current {
			return true
		}
	}

	return false
}
