package main

import "fmt"

type Parser struct {
	tokens Tokens
	index  int
	errors []error
}

func parse(tokens Tokens) (Node, error) {
	p := &Parser{
		tokens: tokens,
		index:  0,
		errors: nil,
	}

	if node, err := p.parse(); err != nil {
		return node, err
	} else {
		return node, p.Error()
	}
}

func (p *Parser) Error() error {
	if len(p.errors) == 0 {
		return nil
	}

	txt := "type check errors"

	for _, v := range p.errors {
		txt = txt + ": " + v.Error()
	}

	return fmt.Errorf(txt)
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

	var (
		consts []ConstDecl
		types  []TypeDecl
		vars   []VarDecl
		stmts  Node
	)

	if p.expect(TokenCONST) {
		consts, err = p.parseConsts()
		if err != nil {
			return nil, err
		}
	}

	if p.expect(TokenTYPE) {
		types, err = p.parseTypes()
		if err != nil {
			return nil, err
		}
	}

	if p.expect(TokenVAR) {
		vars, err = p.parseVars()
		if err != nil {
			return nil, err
		}
	}

	for p.currentType() == TokenPROCEDURE {
		return nil, fmt.Errorf("PROCEDURE is unimplemented")
	}

	if p.expect(TokenBEGIN) {
		stmts, err = p.parseStmtSequence(TokenEND)
		if err != nil {
			return nil, err
		}
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
		token:  t,
		name:   name.Text,
		consts: consts,
		types:  types,
		vars:   vars,
		stmts:  stmts,
	}, nil
}

func (p *Parser) parseConsts() ([]ConstDecl, error) {
	var consts []ConstDecl

	for p.currentType() == TokenIdent {
		t, _ := p.require(TokenIdent)

		if _, err := p.require(TokenEQ); err != nil {
			return nil, err
		}

		expr, err := p.parseExpr()
		if err != nil {
			return nil, err
		}

		if _, err := p.require(TokenSemicolon); err != nil {
			return nil, err
		}

		consts = append(consts, ConstDecl{
			token: t,
			expr:  expr,
		})
	}

	return consts, nil
}

func (p *Parser) parseTypes() ([]TypeDecl, error) {
	var types []TypeDecl

	for p.currentType() == TokenIdent {
		t, _ := p.require(TokenIdent)

		if _, err := p.require(TokenEQ); err != nil {
			return nil, err
		}

		switch p.currentType() {
		case TokenARRAY:
			p.errors = append(p.errors, fmt.Errorf("ARRAY types are not implemented"))
		case TokenRECORD:
			p.errors = append(p.errors, fmt.Errorf("RECORD types are not implemented"))
		case TokenPOINTER:
			p.errors = append(p.errors, fmt.Errorf("POINTER types are not implemented"))
		case TokenPROCEDURE:
			p.errors = append(p.errors, fmt.Errorf("PROCEDURE types are not implemented"))
		default:
			return nil, fmt.Errorf("type declarations can only be `ARRAY`, `RECORD`, `POINTER` or `PROCEDURE`")
		}

		// TODO(daniel): the following is wrong!
		typToken := p.currentToken()
		for !p.expect(TokenSemicolon) {
			p.index++
		}

		types = append(types, TypeDecl{
			token:    t,
			typToken: typToken,
			typ:      TypeVoid,
		})
	}

	return types, nil
}

func (p *Parser) parseVars() ([]VarDecl, error) {
	var vars []VarDecl

	for p.currentType() == TokenIdent {
		varIdents, err := p.parseIdentList()
		if err != nil {
			return nil, err
		}

		if _, err := p.require(TokenColon); err != nil {
			return nil, err
		}

		typeIdent, err := p.parseType()
		if err != nil {
			return nil, err
		}

		if _, err := p.require(TokenSemicolon); err != nil {
			return nil, err
		}

		for _, varIdent := range varIdents {
			vars = append(vars, VarDecl{
				token:    varIdent,
				typToken: typeIdent,
			})
		}
	}

	return vars, nil
}

func (p *Parser) parseIdentList() ([]Token, error) {
	var varIdents []Token

	for {
		varIdent, err := p.require(TokenIdent)
		if err != nil {
			return nil, err
		}

		varIdents = append(varIdents, varIdent)

		if !p.expect(TokenComma) {
			break
		}
	}

	return varIdents, nil
}

func (p *Parser) parseType() (Token, error) {
	// TODO(daniel): complex types.
	return p.require(TokenIdent)
}

func (p *Parser) parseStmtSequence(terminator ...TokenType) (Stmt, error) {
	node := &StmtSequence{}

	for !p.tokenIs(terminator...) {
		stmt, err := p.parseStmt()
		if err != nil {
			return nil, err
		}

		node.stmts = append(node.stmts, stmt)

		if !p.tokenIs(terminator...) {
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
		ident, _ := p.require(TokenIdent)

		switch p.currentType() {
		case TokenAssign:
			return p.parseAssignStmt(ident)
		default:
			expr, err := p.parseCallExpr(ident)
			if err != nil {
				return nil, err
			}

			return &ExprStmt{
				expr: expr,
			}, nil
		}
	case TokenIF:
		return p.parseIfStmt()
	default:
		return nil, fmt.Errorf("unexpected token: %v", p.tokens[p.index].Type)
	}
}

func (p *Parser) parseAssignStmt(ident Token) (Stmt, error) {
	p.consume(TokenAssign)

	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	return &AssignStmt{
		token: ident,
		expr:  expr,
	}, nil
}

func (p *Parser) parseIfStmt() (Stmt, error) {
	t, err := p.require(TokenIF)
	if err != nil {
		return nil, err
	}

	return p.parseIfTail(t)
}

func (p *Parser) parseIfTail(t Token) (Stmt, error) {
	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	if _, err := p.require(TokenTHEN); err != nil {
		return nil, err
	}

	var trueBlock, falseBlock Stmt

	trueBlock, err = p.parseStmtSequence(TokenELSIF, TokenELSE, TokenEND)
	if err != nil {
		return nil, err
	}

	switch p.currentType() {
	case TokenELSIF:
		e, _ := p.require(TokenELSIF)

		falseBlock, err = p.parseIfTail(e)
		if err != nil {
			return nil, err
		}
	case TokenELSE:
		p.consume(TokenELSE)

		falseBlock, err = p.parseStmtSequence(TokenEND)
		if err != nil {
			return nil, err
		}

		p.consume(TokenEND)
	case TokenEND:
		p.consume(TokenEND)

		falseBlock = &StmtSequence{}
	default:
		return nil, fmt.Errorf("expected ELSIF, ELSE or END after IF, got %v", p.currentType())
	}

	return &IfStmt{
		token:      t,
		expr:       expr,
		trueBlock:  trueBlock,
		falseBlock: falseBlock,
	}, nil
}

func (p *Parser) parseCallExpr(ident Token) (Expr, error) {
	node := &CallExpr{
		token: ident,
		args:  nil,
	}

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
	// factor := designator [actualParameters] | number | string | boolean | '(' expr ')' | '~' factor
	switch p.currentType() {
	case TokenIdent:
		return p.parseDesignator()
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

func (p *Parser) parseDesignator() (Expr, error) {
	ident, _ := p.require(TokenIdent)

	return &DesignatorExpr{
		token: ident,
	}, nil
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
		p.consume(TokenPlus)
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

func (p *Parser) consume(exp TokenType) {
	if p.currentType() != exp {
		panic(fmt.Sprintf("tried to consume %v but got %v", exp, p.currentType()))
	}

	p.index++
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
