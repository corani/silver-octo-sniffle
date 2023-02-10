package parser

import (
	"fmt"

	"github.com/corani/silver-octo-sniffle/ast"
	"github.com/corani/silver-octo-sniffle/token"
)

type Parser struct {
	tokens token.Tokens
	index  int
	errors []error
}

func Parse(tokens token.Tokens) (ast.Node, error) {
	p := &Parser{
		tokens: tokens,
		index:  0,
		errors: nil,
	}

	if node, err := p.parseModule(); err != nil {
		return node, err
	} else {
		return node, p.Error()
	}
}

func (p *Parser) Error() error {
	if len(p.errors) == 0 {
		return nil
	}

	txt := "parse errors"

	for _, v := range p.errors {
		txt = txt + ": " + v.Error()
	}

	return fmt.Errorf(txt)
}

func (p *Parser) parseModule() (ast.Node, error) {
	t, err := p.require(token.TokenMODULE)
	if err != nil {
		return nil, err
	}

	name, err := p.require(token.TokenIdent)
	if err != nil {
		return nil, err
	}

	if _, err := p.require(token.TokenSemicolon); err != nil {
		return nil, err
	}

	var (
		consts []*ast.ConstDecl
		types  []*ast.TypeDecl
		vars   []*ast.VarDecl
		stmts  ast.Node
	)

	if p.expect(token.TokenCONST) {
		consts, err = p.parseConsts()
		if err != nil {
			return nil, err
		}
	}

	if p.expect(token.TokenTYPE) {
		types, err = p.parseTypes()
		if err != nil {
			return nil, err
		}
	}

	if p.expect(token.TokenVAR) {
		vars, err = p.parseVars()
		if err != nil {
			return nil, err
		}
	}

	for p.currentType() == token.TokenPROCEDURE {
		return nil, fmt.Errorf("PROCEDURE is unimplemented")
	}

	if p.expect(token.TokenBEGIN) {
		stmts, err = p.parseStmtSequence(token.TokenEND)
		if err != nil {
			return nil, err
		}
	}

	if _, err := p.require(token.TokenEND); err != nil {
		return nil, err
	}

	if v, err := p.require(token.TokenIdent); err != nil {
		return nil, err
	} else {
		if v.Text != name.Text {
			return nil, fmt.Errorf("unexpected identifier %q at end of module %q", v.Text, name.Text)
		}
	}

	if _, err := p.require(token.TokenDot); err != nil {
		return nil, err
	}

	if _, err := p.require(token.TokenEOF); err != nil {
		return nil, err
	}

	return ast.NewModule(t, name.Text, stmts, consts, types, vars), nil
}

func (p *Parser) parseConsts() ([]*ast.ConstDecl, error) {
	var consts []*ast.ConstDecl

	for p.currentType() == token.TokenIdent {
		t, _ := p.require(token.TokenIdent)

		if _, err := p.require(token.TokenEQ); err != nil {
			return nil, err
		}

		expr, err := p.parseExpr()
		if err != nil {
			return nil, err
		}

		if _, err := p.require(token.TokenSemicolon); err != nil {
			return nil, err
		}

		consts = append(consts, ast.NewConstDecl(t, expr))
	}

	return consts, nil
}

func (p *Parser) parseTypes() ([]*ast.TypeDecl, error) {
	var types []*ast.TypeDecl

	for p.currentType() == token.TokenIdent {
		t, _ := p.require(token.TokenIdent)

		if _, err := p.require(token.TokenEQ); err != nil {
			return nil, err
		}

		switch p.currentType() {
		case token.TokenARRAY:
			p.errors = append(p.errors, fmt.Errorf("ARRAY types are not implemented"))
		case token.TokenRECORD:
			p.errors = append(p.errors, fmt.Errorf("RECORD types are not implemented"))
		case token.TokenPOINTER:
			p.errors = append(p.errors, fmt.Errorf("POINTER types are not implemented"))
		case token.TokenPROCEDURE:
			p.errors = append(p.errors, fmt.Errorf("PROCEDURE types are not implemented"))
		default:
			return nil, fmt.Errorf("type declarations can only be `ARRAY`, `RECORD`, `POINTER` or `PROCEDURE`")
		}

		// TODO(daniel): the following is wrong!
		typToken := p.currentToken()
		for !p.expect(token.TokenSemicolon) {
			p.index++
		}

		types = append(types, ast.NewTypeDecl(t, typToken))
	}

	return types, nil
}

func (p *Parser) parseVars() ([]*ast.VarDecl, error) {
	var vars []*ast.VarDecl

	for p.currentType() == token.TokenIdent {
		varIdents, err := p.parseIdentList()
		if err != nil {
			return nil, err
		}

		if _, err := p.require(token.TokenColon); err != nil {
			return nil, err
		}

		typeIdent, err := p.parseType()
		if err != nil {
			return nil, err
		}

		if _, err := p.require(token.TokenSemicolon); err != nil {
			return nil, err
		}

		for _, varIdent := range varIdents {
			vars = append(vars, ast.NewVarDecl(varIdent, typeIdent))
		}
	}

	return vars, nil
}

func (p *Parser) parseIdentList() (token.Tokens, error) {
	var varIdents token.Tokens

	for {
		varIdent, err := p.require(token.TokenIdent)
		if err != nil {
			return nil, err
		}

		varIdents = append(varIdents, varIdent)

		if !p.expect(token.TokenComma) {
			break
		}
	}

	return varIdents, nil
}

func (p *Parser) parseType() (token.Token, error) {
	// TODO(daniel): complexer types.
	return p.require(token.TokenIdent)
}

func (p *Parser) parseStmtSequence(terminator ...token.TokenType) (ast.Stmt, error) {
	var stmts []ast.Stmt

	for !p.tokenIs(terminator...) {
		stmt, err := p.parseStmt()
		if err != nil {
			return nil, err
		}

		stmts = append(stmts, stmt)

		if !p.tokenIs(terminator...) {
			if _, err := p.require(token.TokenSemicolon); err != nil {
				return nil, err
			}
		}
	}

	return ast.NewStmtSequence(stmts), nil
}

func (p *Parser) parseStmt() (ast.Stmt, error) {
	switch p.currentType() {
	case token.TokenIdent:
		ident, _ := p.require(token.TokenIdent)

		switch p.currentType() {
		case token.TokenAssign:
			return p.parseAssignStmt(ident)
		default:
			expr, err := p.parseCallExpr(ast.NewDesignatorExpr(ident))
			if err != nil {
				return nil, err
			}

			return ast.NewExprStmt(expr), nil
		}
	case token.TokenIF:
		return p.parseIfStmt()
	case token.TokenREPEAT:
		return p.parseRepeatStmt()
	case token.TokenWHILE:
		return p.parseWhileStmt()
	case token.TokenFOR:
		return p.parseForStmt()
	default:
		return nil, fmt.Errorf("unexpected token: %v", p.tokens[p.index].Type)
	}
}

func (p *Parser) parseAssignStmt(ident token.Token) (ast.Stmt, error) {
	p.consume(token.TokenAssign)

	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	return ast.NewAssignStmt(ident, expr), nil
}

func (p *Parser) parseIfStmt() (ast.Stmt, error) {
	t, err := p.require(token.TokenIF)
	if err != nil {
		return nil, err
	}

	return p.parseIfTail(t)
}

func (p *Parser) parseIfTail(t token.Token) (ast.Stmt, error) {
	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	if _, err := p.require(token.TokenTHEN); err != nil {
		return nil, err
	}

	var trueBlock, falseBlock ast.Stmt

	trueBlock, err = p.parseStmtSequence(token.TokenELSIF, token.TokenELSE, token.TokenEND)
	if err != nil {
		return nil, err
	}

	switch p.currentType() {
	case token.TokenELSIF:
		e, _ := p.require(token.TokenELSIF)

		falseBlock, err = p.parseIfTail(e)
		if err != nil {
			return nil, err
		}
	case token.TokenELSE:
		p.consume(token.TokenELSE)

		falseBlock, err = p.parseStmtSequence(token.TokenEND)
		if err != nil {
			return nil, err
		}

		p.consume(token.TokenEND)
	case token.TokenEND:
		p.consume(token.TokenEND)

		falseBlock = ast.NewStmtSequence(nil)
	default:
		return nil, fmt.Errorf("expected ELSIF, ELSE or END after IF, got %v", p.currentType())
	}

	return ast.NewIfStmt(t, expr, trueBlock, falseBlock), nil
}

func (p *Parser) parseRepeatStmt() (ast.Stmt, error) {
	t, err := p.require(token.TokenREPEAT)
	if err != nil {
		return nil, err
	}

	stmts, err := p.parseStmtSequence(token.TokenUNTIL)
	if err != nil {
		return nil, err
	}

	if _, err := p.require(token.TokenUNTIL); err != nil {
		return nil, err
	}

	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	cond := ast.NewCondition(expr, stmts)

	return ast.NewRepeatStmt(t, cond), nil
}

func (p *Parser) parseWhileStmt() (ast.Stmt, error) {
	t, err := p.require(token.TokenWHILE)
	if err != nil {
		return nil, err
	}

	var conds []ast.Condition

	for {
		expr, err := p.parseExpr()
		if err != nil {
			return nil, err
		}

		if _, err := p.require(token.TokenDO); err != nil {
			return nil, err
		}

		stmts, err := p.parseStmtSequence(token.TokenELSIF, token.TokenEND)
		if err != nil {
			return nil, err
		}

		conds = append(conds, ast.NewCondition(expr, stmts))

		if !p.expect(token.TokenELSIF) {
			break
		}
	}

	if _, err := p.require(token.TokenEND); err != nil {
		return nil, err
	}

	return ast.NewWhileStmt(t, conds), nil
}

func (p *Parser) parseForStmt() (ast.Stmt, error) {
	t, err := p.require(token.TokenFOR)
	if err != nil {
		return nil, err
	}

	iter, err := p.require(token.TokenIdent)
	if err != nil {
		return nil, err
	}

	if _, err := p.require(token.TokenAssign); err != nil {
		return nil, err
	}

	from, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	if _, err := p.require(token.TokenTO); err != nil {
		return nil, err
	}

	to, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	var by ast.Expr

	if p.expect(token.TokenBY) {
		by, err = p.parseExpr()
		if err != nil {
			return nil, err
		}
	} else {
		// NOTE(daniel): if no `BY` is specified, synthesize a `BY 1`.
		by = ast.NewNumberExpr(token.Token{
			File:  t.File,
			Range: t.Range,
			Type:  token.TokenInteger,
			Text:  "1",
			Int:   1,
		})
	}

	if _, err := p.require(token.TokenDO); err != nil {
		return nil, err
	}

	stmt, err := p.parseStmtSequence(token.TokenEND)
	if err != nil {
		return nil, err
	}

	if _, err := p.require(token.TokenEND); err != nil {
		return nil, err
	}

	return ast.NewForStmt(t, iter, from, to, by, stmt), nil
}

func (p *Parser) parseCallExpr(designator *ast.DesignatorExpr) (ast.Expr, error) {
	if _, err := p.require(token.TokenLParen); err != nil {
		return nil, err
	}

	var args []ast.Expr

	for {
		expr, err := p.parseExpr()
		if err != nil {
			return nil, err
		}

		args = append(args, expr)

		if !p.expect(token.TokenComma) {
			break
		}
	}

	if _, err := p.require(token.TokenRParen); err != nil {
		return nil, err
	}

	return ast.NewCallExpr(designator, args), nil
}

func (p *Parser) parseExpr() (ast.Expr, error) {
	// expr := simpleExpr [ '=' | '#' | '<' | '<=' | '>' | '>=' | 'IN' | 'IS'  simpleExpr ]
	lhs, err := p.parseSimpleExpr()
	if err != nil {
		return lhs, err
	}

	for op := p.currentToken(); op.IsRelation(); op = p.currentToken() {
		p.index++

		rhs, err := p.parseSimpleExpr()
		if err != nil {
			return rhs, err
		}

		lhs = ast.NewBinaryExpr(op, lhs, rhs)
	}

	return lhs, nil
}

func (p *Parser) parseSimpleExpr() (ast.Expr, error) {
	// simpleExpr := term { '+' | '-' | 'OR' term }
	addOperators := []token.TokenType{
		token.TokenPlus, token.TokenMinus, token.TokenOR,
	}

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

		lhs = ast.NewBinaryExpr(op, lhs, rhs)
	}

	return lhs, nil
}

func (p *Parser) parseTerm() (ast.Expr, error) {
	// term := factor { '*' | '/' | 'DIV' | 'MUL' | '&' factor }
	mulOperators := []token.TokenType{
		token.TokenAsterisk, token.TokenSlash, token.TokenDIV, token.TokenMOD, token.TokenAmpersand,
	}

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

		lhs = ast.NewBinaryExpr(op, lhs, rhs)
	}

	return lhs, nil
}

func (p *Parser) parseFactor() (ast.Expr, error) {
	// factor := set | designator [actualParameters] | number | string | boolean | '(' expr ')' | '~' factor
	switch p.currentType() {
	case token.TokenIdent:
		d, err := p.parseDesignator()
		if err != nil {
			return nil, err
		}

		if p.currentType() != token.TokenLParen {
			return d, nil
		}

		return p.parseCallExpr(d)
	case token.TokenInteger, token.TokenReal, token.TokenMinus, token.TokenPlus:
		return p.parseNumberExpr()
	case token.TokenString:
		return p.parseStringLiteral()
	case token.TokenBoolean:
		return p.parseBooleanLiteral()
	case token.TokenTilde:
		return p.parseNotExpr()
	case token.TokenLParen:
		if _, err := p.require(token.TokenLParen); err != nil {
			return nil, err
		}

		expr, err := p.parseExpr()
		if err != nil {
			return expr, err
		}

		if _, err := p.require(token.TokenRParen); err != nil {
			return nil, err
		}

		return expr, nil
	case token.TokenLBrace:
		return p.parseSetLiteral()
	default:
		return nil, fmt.Errorf("unexpected token in factor: %q", p.currentType())
	}
}

func (p *Parser) parseDesignator() (*ast.DesignatorExpr, error) {
	ident, _ := p.require(token.TokenIdent)

	return ast.NewDesignatorExpr(ident), nil
}

func (p *Parser) parseNumberExpr() (ast.Expr, error) {
	var (
		op token.Token
		t  token.Token
	)

	switch p.currentType() {
	case token.TokenMinus:
		// Unary Minus.
		op, _ = p.require(token.TokenMinus)
	case token.TokenPlus:
		// Unary Plus is a no-op so we ignore it.
		p.consume(token.TokenPlus)
	}

	switch p.currentType() {
	case token.TokenInteger:
		t, _ = p.require(token.TokenInteger)
	case token.TokenReal:
		t, _ = p.require(token.TokenReal)
	default:
		return nil, fmt.Errorf("expected number, got %v", p.currentType())
	}

	num := ast.NewNumberExpr(t)

	if op.Type == token.TokenInvalid {
		return num, nil
	}

	// NOTE(daniel): There is no negate for integers in LLVM IR, so we treat
	// `-x` as `0 - x`. That means we don't need generator code either.
	t.Int = 0
	t.Text = ""

	zero := ast.NewNumberExpr(t)

	return ast.NewBinaryExpr(op, zero, num), nil
}

func (p *Parser) parseStringLiteral() (ast.Expr, error) {
	t, err := p.require(token.TokenString)
	if err != nil {
		return nil, err
	}

	if len(t.Text) == 1 {
		return ast.NewCharExpr(t), nil
	} else {
		return ast.NewStringExpr(t), nil
	}
}

func (p *Parser) parseBooleanLiteral() (ast.Expr, error) {
	t, err := p.require(token.TokenBoolean)
	if err != nil {
		return nil, err
	}

	return ast.NewBooleanExpr(t), nil
}

func (p *Parser) parseSetLiteral() (ast.Expr, error) {
	t, err := p.require(token.TokenLBrace)
	if err != nil {
		return nil, err
	}

	var bits []byte

	for p.currentType() != token.TokenRBrace {
		i1, err := p.require(token.TokenInteger)
		if err != nil {
			return nil, err
		}

		if !p.expect(token.TokenDotDot) {
			bits = append(bits, byte(i1.Int))

			if p.expect(token.TokenComma) {
				continue
			}

			break
		}

		i2, err := p.require(token.TokenInteger)
		if err != nil {
			return nil, err
		}

		for i := i1.Int; i <= i2.Int; i++ {
			bits = append(bits, byte(i))
		}

		if p.expect(token.TokenComma) {
			continue
		}
	}

	if _, err := p.require(token.TokenRBrace); err != nil {
		return nil, err
	}

	return ast.NewSetExpr(t, bits), nil
}

func (p *Parser) parseNotExpr() (ast.Expr, error) {
	t, err := p.require(token.TokenTilde)
	if err != nil {
		return nil, err
	}

	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	return ast.NewNotExpr(t, expr), nil
}

func (p *Parser) require(exp token.TokenType) (token.Token, error) {
	if p.currentType() == exp {
		token := p.currentToken()
		p.index++

		return token, nil
	}

	return token.Token{}, fmt.Errorf("unexpected token: %v (expected %v)", p.tokens[p.index].Type, exp)
}

func (p *Parser) expect(exp token.TokenType) bool {
	if p.currentType() == exp {
		p.index++

		return true
	}

	return false
}

func (p *Parser) consume(exp token.TokenType) {
	if p.currentType() != exp {
		panic(fmt.Sprintf("tried to consume %v but got %v", exp, p.currentType()))
	}

	p.index++
}

func (p *Parser) currentToken() token.Token {
	return p.tokens[p.index]
}

func (p *Parser) currentType() token.TokenType {
	return p.tokens[p.index].Type
}

func (p *Parser) tokenIs(opts ...token.TokenType) bool {
	current := p.currentType()

	for _, opt := range opts {
		if opt == current {
			return true
		}
	}

	return false
}
