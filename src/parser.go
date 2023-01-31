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
	// expr := simpleExpr [ '=' | '#' | '<' | '<=' | '>' | '>=' | 'IN' | 'IS'  simpleExpr ]
	lhs, err := p.parseSimpleExpr()
	if err != nil {
		return lhs, err
	}

	for op := p.currentToken(); op.isRelation(); op = p.currentToken() {
		p.index++

		rhs, err := p.parseSimpleExpr()
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

func (p *Parser) parseSimpleExpr() (Expr, error) {
	// simpleExpr := term { '+' | '-' | 'OR' term }
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
	mulOperators := []TokenType{TokenAsterisk, TokenSlash, TokenIDiv, TokenMod, TokenAnd}

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
	// factor := number | string | boolean | '(' expr ')' | '~' factor
	switch p.currentType() {
	case TokenInteger, TokenReal, TokenMinus:
		return p.parseNumberExpr()
	case TokenString:
		return p.parseStringLiteral()
	case TokenBoolean:
		return p.parseBooleanLiteral()
	case TokenNot:
		return p.parseNotExpr()
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
		panic(fmt.Sprintf("unexpected token in factor: %q", p.currentType()))
	}
}

func (p *Parser) parseNumberExpr() (Expr, error) {
	var (
		op Token
		t  Token
	)

	switch p.currentType() {
	case TokenMinus:
		// Unary Minus.
		op, _ = p.require(TokenMinus)
	case TokenPlus:
		// Unary Plus is a no-op so we ignore it.
		_, _ = p.require(TokenPlus)
	}

	switch p.currentType() {
	case TokenInteger:
		t, _ = p.require(TokenInteger)
	case TokenReal:
		t, _ = p.require(TokenReal)
	default:
		return nil, fmt.Errorf("expected number, got %v", p.currentType())
	}

	num := &NumberExpr{token: t}

	if op.Type == TokenInvalid {
		return num, nil
	}

	// NOTE(daniel): There is no negate for integers in LLVM IR, so we treat
	// `-x` as `0 - x`. That means we don't need generator code either.
	zero := &NumberExpr{token: t}
	zero.token.Int = 0
	zero.token.Text = ""

	return &BinaryExpr{
		token: op,
		args:  []Expr{zero, num},
	}, nil
}

func (p *Parser) parseStringLiteral() (Expr, error) {
	t, err := p.require(TokenString)
	if err != nil {
		return nil, err
	}

	return &StringExpr{token: t}, nil
}

func (p *Parser) parseBooleanLiteral() (Expr, error) {
	t, err := p.require(TokenBoolean)
	if err != nil {
		return nil, err
	}

	return &BooleanExpr{token: t}, nil
}

func (p *Parser) parseNotExpr() (Expr, error) {
	t, err := p.require(TokenNot)
	if err != nil {
		return nil, err
	}

	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	return &NotExpr{
		token: t,
		expr:  expr,
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

func (p *Parser) tokenIs(opts ...TokenType) bool {
	current := p.currentType()

	for _, opt := range opts {
		if opt == current {
			return true
		}
	}

	return false
}
