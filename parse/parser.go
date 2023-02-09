package parse

import (
	"fmt"

	"github.com/corani/silver-octo-sniffle/ast"
	"github.com/corani/silver-octo-sniffle/lex"
)

type Parser struct {
	tokens lex.Tokens
	index  int
	errors []error
}

func Parse(tokens lex.Tokens) (ast.Node, error) {
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
	t, err := p.require(lex.TokenMODULE)
	if err != nil {
		return nil, err
	}

	name, err := p.require(lex.TokenIdent)
	if err != nil {
		return nil, err
	}

	if _, err := p.require(lex.TokenSemicolon); err != nil {
		return nil, err
	}

	var (
		consts []*ast.ConstDecl
		types  []*ast.TypeDecl
		vars   []*ast.VarDecl
		stmts  ast.Node
	)

	if p.expect(lex.TokenCONST) {
		consts, err = p.parseConsts()
		if err != nil {
			return nil, err
		}
	}

	if p.expect(lex.TokenTYPE) {
		types, err = p.parseTypes()
		if err != nil {
			return nil, err
		}
	}

	if p.expect(lex.TokenVAR) {
		vars, err = p.parseVars()
		if err != nil {
			return nil, err
		}
	}

	for p.currentType() == lex.TokenPROCEDURE {
		return nil, fmt.Errorf("PROCEDURE is unimplemented")
	}

	if p.expect(lex.TokenBEGIN) {
		stmts, err = p.parseStmtSequence(lex.TokenEND)
		if err != nil {
			return nil, err
		}
	}

	if _, err := p.require(lex.TokenEND); err != nil {
		return nil, err
	}

	if v, err := p.require(lex.TokenIdent); err != nil {
		return nil, err
	} else {
		if v.Text != name.Text {
			return nil, fmt.Errorf("unexpected identifier %q at end of module %q", v.Text, name.Text)
		}
	}

	if _, err := p.require(lex.TokenDot); err != nil {
		return nil, err
	}

	if _, err := p.require(lex.TokenEOF); err != nil {
		return nil, err
	}

	return ast.NewModule(t, name.Text, stmts, consts, types, vars), nil
}

func (p *Parser) parseConsts() ([]*ast.ConstDecl, error) {
	var consts []*ast.ConstDecl

	for p.currentType() == lex.TokenIdent {
		t, _ := p.require(lex.TokenIdent)

		if _, err := p.require(lex.TokenEQ); err != nil {
			return nil, err
		}

		expr, err := p.parseExpr()
		if err != nil {
			return nil, err
		}

		if _, err := p.require(lex.TokenSemicolon); err != nil {
			return nil, err
		}

		consts = append(consts, ast.NewConstDecl(t, expr))
	}

	return consts, nil
}

func (p *Parser) parseTypes() ([]*ast.TypeDecl, error) {
	var types []*ast.TypeDecl

	for p.currentType() == lex.TokenIdent {
		t, _ := p.require(lex.TokenIdent)

		if _, err := p.require(lex.TokenEQ); err != nil {
			return nil, err
		}

		switch p.currentType() {
		case lex.TokenARRAY:
			p.errors = append(p.errors, fmt.Errorf("ARRAY types are not implemented"))
		case lex.TokenRECORD:
			p.errors = append(p.errors, fmt.Errorf("RECORD types are not implemented"))
		case lex.TokenPOINTER:
			p.errors = append(p.errors, fmt.Errorf("POINTER types are not implemented"))
		case lex.TokenPROCEDURE:
			p.errors = append(p.errors, fmt.Errorf("PROCEDURE types are not implemented"))
		default:
			return nil, fmt.Errorf("type declarations can only be `ARRAY`, `RECORD`, `POINTER` or `PROCEDURE`")
		}

		// TODO(daniel): the following is wrong!
		typToken := p.currentToken()
		for !p.expect(lex.TokenSemicolon) {
			p.index++
		}

		types = append(types, ast.NewTypeDecl(t, typToken))
	}

	return types, nil
}

func (p *Parser) parseVars() ([]*ast.VarDecl, error) {
	var vars []*ast.VarDecl

	for p.currentType() == lex.TokenIdent {
		varIdents, err := p.parseIdentList()
		if err != nil {
			return nil, err
		}

		if _, err := p.require(lex.TokenColon); err != nil {
			return nil, err
		}

		typeIdent, err := p.parseType()
		if err != nil {
			return nil, err
		}

		if _, err := p.require(lex.TokenSemicolon); err != nil {
			return nil, err
		}

		for _, varIdent := range varIdents {
			vars = append(vars, ast.NewVarDecl(varIdent, typeIdent))
		}
	}

	return vars, nil
}

func (p *Parser) parseIdentList() (lex.Tokens, error) {
	var varIdents lex.Tokens

	for {
		varIdent, err := p.require(lex.TokenIdent)
		if err != nil {
			return nil, err
		}

		varIdents = append(varIdents, varIdent)

		if !p.expect(lex.TokenComma) {
			break
		}
	}

	return varIdents, nil
}

func (p *Parser) parseType() (lex.Token, error) {
	// TODO(daniel): complex types.
	return p.require(lex.TokenIdent)
}

func (p *Parser) parseStmtSequence(terminator ...lex.TokenType) (ast.Stmt, error) {
	var stmts []ast.Stmt

	for !p.tokenIs(terminator...) {
		stmt, err := p.parseStmt()
		if err != nil {
			return nil, err
		}

		stmts = append(stmts, stmt)

		if !p.tokenIs(terminator...) {
			if _, err := p.require(lex.TokenSemicolon); err != nil {
				return nil, err
			}
		}
	}

	return ast.NewStmtSequence(stmts), nil
}

func (p *Parser) parseStmt() (ast.Stmt, error) {
	switch p.currentType() {
	case lex.TokenIdent:
		ident, _ := p.require(lex.TokenIdent)

		switch p.currentType() {
		case lex.TokenAssign:
			return p.parseAssignStmt(ident)
		default:
			expr, err := p.parseCallExpr(ast.NewDesignatorExpr(ident))
			if err != nil {
				return nil, err
			}

			return ast.NewExprStmt(expr), nil
		}
	case lex.TokenIF:
		return p.parseIfStmt()
	case lex.TokenREPEAT:
		return p.parseRepeatStmt()
	case lex.TokenWHILE:
		return p.parseWhileStmt()
	case lex.TokenFOR:
		return p.parseForStmt()
	default:
		return nil, fmt.Errorf("unexpected token: %v", p.tokens[p.index].Type)
	}
}

func (p *Parser) parseAssignStmt(ident lex.Token) (ast.Stmt, error) {
	p.consume(lex.TokenAssign)

	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	return ast.NewAssignStmt(ident, expr), nil
}

func (p *Parser) parseIfStmt() (ast.Stmt, error) {
	t, err := p.require(lex.TokenIF)
	if err != nil {
		return nil, err
	}

	return p.parseIfTail(t)
}

func (p *Parser) parseIfTail(t lex.Token) (ast.Stmt, error) {
	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	if _, err := p.require(lex.TokenTHEN); err != nil {
		return nil, err
	}

	var trueBlock, falseBlock ast.Stmt

	trueBlock, err = p.parseStmtSequence(lex.TokenELSIF, lex.TokenELSE, lex.TokenEND)
	if err != nil {
		return nil, err
	}

	switch p.currentType() {
	case lex.TokenELSIF:
		e, _ := p.require(lex.TokenELSIF)

		falseBlock, err = p.parseIfTail(e)
		if err != nil {
			return nil, err
		}
	case lex.TokenELSE:
		p.consume(lex.TokenELSE)

		falseBlock, err = p.parseStmtSequence(lex.TokenEND)
		if err != nil {
			return nil, err
		}

		p.consume(lex.TokenEND)
	case lex.TokenEND:
		p.consume(lex.TokenEND)

		falseBlock = ast.NewStmtSequence(nil)
	default:
		return nil, fmt.Errorf("expected ELSIF, ELSE or END after IF, got %v", p.currentType())
	}

	return ast.NewIfStmt(t, expr, trueBlock, falseBlock), nil
}

func (p *Parser) parseRepeatStmt() (ast.Stmt, error) {
	t, err := p.require(lex.TokenREPEAT)
	if err != nil {
		return nil, err
	}

	stmts, err := p.parseStmtSequence(lex.TokenUNTIL)
	if err != nil {
		return nil, err
	}

	if _, err := p.require(lex.TokenUNTIL); err != nil {
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
	t, err := p.require(lex.TokenWHILE)
	if err != nil {
		return nil, err
	}

	var conds []ast.Condition

	for {
		expr, err := p.parseExpr()
		if err != nil {
			return nil, err
		}

		if _, err := p.require(lex.TokenDO); err != nil {
			return nil, err
		}

		stmts, err := p.parseStmtSequence(lex.TokenELSIF, lex.TokenEND)
		if err != nil {
			return nil, err
		}

		conds = append(conds, ast.NewCondition(expr, stmts))

		if !p.expect(lex.TokenELSIF) {
			break
		}
	}

	if _, err := p.require(lex.TokenEND); err != nil {
		return nil, err
	}

	return ast.NewWhileStmt(t, conds), nil
}

func (p *Parser) parseForStmt() (ast.Stmt, error) {
	t, err := p.require(lex.TokenFOR)
	if err != nil {
		return nil, err
	}

	iter, err := p.require(lex.TokenIdent)
	if err != nil {
		return nil, err
	}

	if _, err := p.require(lex.TokenAssign); err != nil {
		return nil, err
	}

	from, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	if _, err := p.require(lex.TokenTO); err != nil {
		return nil, err
	}

	to, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	var by ast.Expr

	if p.expect(lex.TokenBY) {
		by, err = p.parseExpr()
		if err != nil {
			return nil, err
		}
	} else {
		// NOTE(daniel): if no `BY` is specified, synthesize a `BY 1`.
		by = ast.NewNumberExpr(lex.Token{
			File:  t.File,
			Range: t.Range,
			Type:  lex.TokenInteger,
			Text:  "1",
			Int:   1,
		})
	}

	if _, err := p.require(lex.TokenDO); err != nil {
		return nil, err
	}

	stmt, err := p.parseStmtSequence(lex.TokenEND)
	if err != nil {
		return nil, err
	}

	if _, err := p.require(lex.TokenEND); err != nil {
		return nil, err
	}

	return ast.NewForStmt(t, iter, from, to, by, stmt), nil
}

func (p *Parser) parseCallExpr(designator *ast.DesignatorExpr) (ast.Expr, error) {
	if _, err := p.require(lex.TokenLParen); err != nil {
		return nil, err
	}

	var args []ast.Expr

	for {
		expr, err := p.parseExpr()
		if err != nil {
			return nil, err
		}

		args = append(args, expr)

		if !p.expect(lex.TokenComma) {
			break
		}
	}

	if _, err := p.require(lex.TokenRParen); err != nil {
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
	addOperators := []lex.TokenType{
		lex.TokenPlus, lex.TokenMinus, lex.TokenOR,
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
	mulOperators := []lex.TokenType{
		lex.TokenAsterisk, lex.TokenSlash, lex.TokenDIV, lex.TokenMOD, lex.TokenAmpersand,
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
	case lex.TokenIdent:
		d, err := p.parseDesignator()
		if err != nil {
			return nil, err
		}

		if p.currentType() != lex.TokenLParen {
			return d, nil
		}

		return p.parseCallExpr(d)
	case lex.TokenInteger, lex.TokenReal, lex.TokenMinus, lex.TokenPlus:
		return p.parseNumberExpr()
	case lex.TokenString:
		return p.parseStringLiteral()
	case lex.TokenBoolean:
		return p.parseBooleanLiteral()
	case lex.TokenTilde:
		return p.parseNotExpr()
	case lex.TokenLParen:
		if _, err := p.require(lex.TokenLParen); err != nil {
			return nil, err
		}

		expr, err := p.parseExpr()
		if err != nil {
			return expr, err
		}

		if _, err := p.require(lex.TokenRParen); err != nil {
			return nil, err
		}

		return expr, nil
	case lex.TokenLBrace:
		return p.parseSetLiteral()
	default:
		return nil, fmt.Errorf("unexpected token in factor: %q", p.currentType())
	}
}

func (p *Parser) parseDesignator() (*ast.DesignatorExpr, error) {
	ident, _ := p.require(lex.TokenIdent)

	return ast.NewDesignatorExpr(ident), nil
}

func (p *Parser) parseNumberExpr() (ast.Expr, error) {
	var (
		op lex.Token
		t  lex.Token
	)

	switch p.currentType() {
	case lex.TokenMinus:
		// Unary Minus.
		op, _ = p.require(lex.TokenMinus)
	case lex.TokenPlus:
		// Unary Plus is a no-op so we ignore it.
		p.consume(lex.TokenPlus)
	}

	switch p.currentType() {
	case lex.TokenInteger:
		t, _ = p.require(lex.TokenInteger)
	case lex.TokenReal:
		t, _ = p.require(lex.TokenReal)
	default:
		return nil, fmt.Errorf("expected number, got %v", p.currentType())
	}

	num := ast.NewNumberExpr(t)

	if op.Type == lex.TokenInvalid {
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
	t, err := p.require(lex.TokenString)
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
	t, err := p.require(lex.TokenBoolean)
	if err != nil {
		return nil, err
	}

	return ast.NewBooleanExpr(t), nil
}

func (p *Parser) parseSetLiteral() (ast.Expr, error) {
	t, err := p.require(lex.TokenLBrace)
	if err != nil {
		return nil, err
	}

	var bits []byte

	for p.currentType() != lex.TokenRBrace {
		i1, err := p.require(lex.TokenInteger)
		if err != nil {
			return nil, err
		}

		if !p.expect(lex.TokenDotDot) {
			bits = append(bits, byte(i1.Int))

			if p.expect(lex.TokenComma) {
				continue
			}

			break
		}

		i2, err := p.require(lex.TokenInteger)
		if err != nil {
			return nil, err
		}

		for i := i1.Int; i <= i2.Int; i++ {
			bits = append(bits, byte(i))
		}

		if p.expect(lex.TokenComma) {
			continue
		}
	}

	if _, err := p.require(lex.TokenRBrace); err != nil {
		return nil, err
	}

	return ast.NewSetExpr(t, bits), nil
}

func (p *Parser) parseNotExpr() (ast.Expr, error) {
	t, err := p.require(lex.TokenTilde)
	if err != nil {
		return nil, err
	}

	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	return ast.NewNotExpr(t, expr), nil
}

func (p *Parser) require(exp lex.TokenType) (lex.Token, error) {
	if p.currentType() == exp {
		token := p.currentToken()
		p.index++

		return token, nil
	}

	return lex.Token{}, fmt.Errorf("unexpected token: %v (expected %v)", p.tokens[p.index].Type, exp)
}

func (p *Parser) expect(exp lex.TokenType) bool {
	if p.currentType() == exp {
		p.index++

		return true
	}

	return false
}

func (p *Parser) consume(exp lex.TokenType) {
	if p.currentType() != exp {
		panic(fmt.Sprintf("tried to consume %v but got %v", exp, p.currentType()))
	}

	p.index++
}

func (p *Parser) currentToken() lex.Token {
	return p.tokens[p.index]
}

func (p *Parser) currentType() lex.TokenType {
	return p.tokens[p.index].Type
}

func (p *Parser) tokenIs(opts ...lex.TokenType) bool {
	current := p.currentType()

	for _, opt := range opts {
		if opt == current {
			return true
		}
	}

	return false
}
