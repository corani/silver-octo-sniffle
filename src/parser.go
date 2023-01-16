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
	switch p.tokens[p.index].Type {
	case TokenPrint:
		return p.parsePrint()
	default:
		return nil, fmt.Errorf("unexpected token: %v", p.tokens[p.index])
	}
}

func (p *Parser) parsePrint() (Node, error) {
	node := &PrintStmt{
		token: p.tokens[p.index],
		args:  nil,
	}

	p.index++

	if _, err := p.require(TokenLParen); err != nil {
		return node, err
	}

	var args []Node
	var op Token

	for p.tokens[p.index].Type != TokenRParen {
		switch p.tokens[p.index].Type {
		case TokenNumber:
			arg, err := p.parseNumber()
			if err != nil {
				return node, err
			}

			args = append(args, arg)

			if len(args) == 2 && op.Type != "" {
				switch op.Type {
				case TokenPlus:
					plus := &BinaryExpr{
						token: op,
						args:  args,
					}

					args = []Node{plus}
					op = Token{}
				}
			}
		case TokenPlus:
			op = p.tokens[p.index]
			p.index++
		}
	}

	node.args = args

	if _, err := p.require(TokenRParen); err != nil {
		return node, err
	}

	return node, nil
}

func (p *Parser) parseNumber() (Node, error) {
	node := &NumberExpr{
		token: p.tokens[p.index],
	}

	p.index++

	return node, nil
}

func (p *Parser) require(exp TokenType) (Token, error) {
	if p.tokens[p.index].Type == exp {
		token := p.tokens[p.index]
		p.index++

		return token, nil
	}

	return Token{}, fmt.Errorf("unexpected token: %v (expected %v)", p.tokens[p.index], exp)
}
