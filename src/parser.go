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
	if p.currentType() == TokenMODULE {
		return p.parseModule()
	}

	stmts, err := p.parseStmtSequence(TokenEOF)
	if err != nil {
		return nil, err
	}

	if _, err := p.require(TokenEOF); err != nil {
		return nil, err
	}

	return &Module{
		token: stmts.Token(),
		name:  "",
		stmts: stmts,
	}, nil
}

func (p *Parser) parseModule() (Node, error) {
	t, err := p.require(TokenMODULE)
	if err != nil {
		return nil, err
	}

	name, err := p.require(TokenIdent)
	if err != nil {
		return nil, err
	}

	if _, err := p.require(TokenSemicolon); err != nil {
		return nil, err
	}

	var stmts Node

	if p.currentType() == TokenBEGIN {
		p.index++

		seq, err := p.parseStmtSequence(TokenEND)
		if err != nil {
			return nil, err
		}

		stmts = seq
	}

	if _, err := p.require(TokenEND); err != nil {
		return nil, err
	}

	if v, err := p.require(TokenIdent); err != nil {
		return nil, err
	} else {
		if v.Text != name.Text {
			return nil, fmt.Errorf("unexpected identifier %q at end of module %q", v.Text, name.Text)
		}
	}

	if _, err := p.require(TokenDot); err != nil {
		return nil, err
	}

	if _, err := p.require(TokenEOF); err != nil {
		return nil, err
	}

	return &Module{
		token: t,
		name:  name.Text,
		stmts: stmts,
	}, nil
}

func (p *Parser) parseStmtSequence(terminator TokenType) (Stmt, error) {
	node := &StmtSequence{}

	for p.currentType() != terminator {
		stmt, err := p.parseStmt()
		if err != nil {
			return nil, err
		}

		node.stmts = append(node.stmts, stmt)

		if p.currentType() != terminator {
			if _, err := p.require(TokenSemicolon); err != nil {
				return nil, err
			}
		}
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
	addOperators := []TokenType{TokenPlus, TokenMinus, TokenOR}

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
	mulOperators := []TokenType{TokenAsterisk, TokenSlash, TokenDIV, TokenMOD, TokenAmpersand}

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
	case TokenTilde:
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
		return nil, fmt.Errorf("unexpected token in factor: %q", p.currentType())
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
	t, err := p.require(TokenTilde)
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

	return Token{}, fmt.Errorf("unexpected token: %v (expected %v)", p.tokens[p.index].Type, exp)
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
