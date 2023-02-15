package parser

import (
	"errors"
	"fmt"

	"github.com/corani/silver-octo-sniffle/ast"
	"github.com/corani/silver-octo-sniffle/reporter"
	"github.com/corani/silver-octo-sniffle/token"
)

var ErrParsing = errors.New("error parsing")

type Parser struct {
	out    *reporter.Reporter
	tokens token.Tokens
	index  int
}

func Parse(out *reporter.Reporter, tokens token.Tokens) (ast.Node, error) {
	p := &Parser{
		out:    out,
		tokens: tokens,
		index:  0,
	}

	return p.parseModule()
}

func (p *Parser) parseModule() (ast.Node, error) {
	t, err := p.require(token.TokenMODULE)
	if err != nil {
		return nil, err
	}

	name, err := p.require(token.TokenIdent)
	if err != nil {
		p.findNextSyncPoint(token.TokenSemicolon, token.TokenCONST, token.TokenTYPE, token.TokenVAR, token.TokenPROCEDURE, token.TokenBEGIN, token.TokenEND)
	}

	if _, err := p.require(token.TokenSemicolon); err != nil {
		p.findNextSyncPoint(token.TokenCONST, token.TokenTYPE, token.TokenVAR, token.TokenPROCEDURE, token.TokenBEGIN, token.TokenEND)
	}

	var (
		decls []ast.Decl
		stmts ast.Node
	)

	if p.expect(token.TokenCONST) {
		for _, v := range p.parseConsts() {
			decls = append(decls, v)
		}
	}

	if p.expect(token.TokenTYPE) {
		for _, v := range p.parseTypes() {
			decls = append(decls, v)
		}
	}

	if p.expect(token.TokenVAR) {
		for _, v := range p.parseVars() {
			decls = append(decls, v)
		}
	}

	if p.expect(token.TokenPROCEDURE) {
		p.parseProcedures()
	}

	if p.expect(token.TokenBEGIN) {
		stmts = p.parseStmtSequence(token.TokenEND)
	}

	if _, err := p.require(token.TokenEND); err != nil {
		p.findNextSyncPoint(token.TokenIdent, token.TokenDot)
	}

	if v, err := p.require(token.TokenIdent); err != nil {
		p.findNextSyncPoint(token.TokenDot)
	} else {
		if v.Text != name.Text {
			p.out.Errorf(p.currentToken(), "unexpected identifier %q at end of module", v.Text)
			p.out.Infof(p.currentToken(), "module name is %q", name.Text)
		}
	}

	if _, err := p.require(token.TokenDot); err != nil {
		p.findNextSyncPoint(token.TokenEOF)
	}

	_, _ = p.require(token.TokenEOF)

	return ast.NewModule(t, name.Text, stmts, decls), nil
}

func (p *Parser) parseConsts() []*ast.ConstDecl {
	var consts []*ast.ConstDecl

	for p.currentType() == token.TokenIdent {
		t, _ := p.require(token.TokenIdent)

		if _, err := p.require(token.TokenEQ); err != nil {
			p.findNextSyncPoint(token.TokenSemicolon,
				token.TokenVAR, token.TokenPROCEDURE, token.TokenBEGIN, token.TokenEND)

			break
		}

		expr := p.parseExpr()

		if _, err := p.require(token.TokenSemicolon); err != nil {
			p.findNextSyncPoint(token.TokenSemicolon,
				token.TokenVAR, token.TokenPROCEDURE, token.TokenBEGIN, token.TokenEND)

			break
		}

		consts = append(consts, ast.NewConstDecl(t, expr))
	}

	return consts
}

func (p *Parser) parseTypes() []ast.TypeDecl {
	var types []ast.TypeDecl

	for p.currentType() == token.TokenIdent {
		t, _ := p.require(token.TokenIdent)

		if _, err := p.require(token.TokenEQ); err != nil {
			p.findNextSyncPoint(token.TokenSemicolon)

			continue
		}

		typeDecl := p.parseType(t)

		if _, err := p.require(token.TokenSemicolon); err != nil {
			p.findNextSyncPoint(token.TokenSemicolon)

			continue
		}

		types = append(types, typeDecl)
	}

	return types
}

func (p *Parser) parseVars() []*ast.VarDecl {
	var vars []*ast.VarDecl

	for p.currentType() == token.TokenIdent {
		varIdents := p.parseIdentList()

		if _, err := p.require(token.TokenColon); err != nil {
			p.findNextSyncPoint(token.TokenSemicolon,
				token.TokenPROCEDURE, token.TokenBEGIN, token.TokenEND)

			break
		}

		typeDecl := p.parseType(token.Token{})

		if _, err := p.require(token.TokenSemicolon); err != nil {
			p.findNextSyncPoint(token.TokenSemicolon,
				token.TokenPROCEDURE, token.TokenBEGIN, token.TokenEND)

			break
		}

		for _, varIdent := range varIdents {
			vars = append(vars, ast.NewVarDecl(varIdent, typeDecl))
		}
	}

	return vars
}

func (p *Parser) parseIdentList() token.Tokens {
	var varIdents token.Tokens

	for {
		varIdent, err := p.require(token.TokenIdent)
		if err != nil {
			p.findNextSyncPoint(token.TokenColon, token.TokenSemicolon,
				token.TokenPROCEDURE, token.TokenBEGIN, token.TokenEND)

			break
		}

		varIdents = append(varIdents, varIdent)

		if !p.expect(token.TokenComma) {
			break
		}
	}

	return varIdents
}

func (p *Parser) parseType(name token.Token) ast.TypeDecl {
	// TODO(daniel): break these out into separate methods?
	switch p.currentType() {
	case token.TokenIdent:
		t, _ := p.require(token.TokenIdent)

		return ast.NewTypeRef(t)
	case token.TokenPOINTER:
		p.consume(token.TokenPOINTER)

		if _, err := p.require(token.TokenTO); err != nil {
			p.findNextSyncPoint(token.TokenSemicolon)

			return ast.NewInvalidTypeDecl(name)
		}

		to := p.parseType(token.Token{})

		return ast.NewTypePointerDecl(name, to)
	case token.TokenARRAY:
	case token.TokenRECORD:
	case token.TokenPROCEDURE:
	}

	p.out.Errorf(p.currentToken(), "invalid type in type declaration: %v", p.currentType())
	p.findNextSyncPoint(token.TokenSemicolon)

	return ast.NewInvalidTypeDecl(name)
}

func (p *Parser) parseStmtSequence(terminator ...token.TokenType) ast.Stmt {
	var stmts []ast.Stmt

	for !p.tokenIs(terminator...) {
		stmt := p.parseStmt()

		stmts = append(stmts, stmt)

		if !p.tokenIs(terminator...) {
			if _, err := p.require(token.TokenSemicolon); err != nil {
				p.findNextSyncPoint(token.TokenSemicolon, token.TokenEND)

				break
			}
		}
	}

	return ast.NewStmtSequence(stmts)
}

func (p *Parser) parseProcedures() {
	for {
		p.out.Errorf(p.currentToken(), "PROCEDURE is unimplemented")
		p.findNextSyncPoint(token.TokenPROCEDURE, token.TokenBEGIN)

		if !p.expect(token.TokenPROCEDURE) {
			break
		}
	}
}

func (p *Parser) parseStmt() ast.Stmt {
	switch p.currentType() {
	case token.TokenIdent:
		des := p.parseDesignator()

		switch p.currentType() {
		case token.TokenAssign:
			return p.parseAssignStmt(des)
		default:
			expr := p.parseCallExpr(des)

			return ast.NewExprStmt(expr)
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
		t := p.currentToken()

		p.out.Errorf(t, "unexpected token: %v", t.Text)
		p.out.Infof(t, "expected a statement")
		p.findNextSyncPoint(token.TokenSemicolon, token.TokenEND)

		return ast.NewInvalidStmt(t)
	}
}

func (p *Parser) parseAssignStmt(des *ast.DesignatorExpr) ast.Stmt {
	p.consume(token.TokenAssign)

	expr := p.parseExpr()

	return ast.NewAssignStmt(des, expr)
}

func (p *Parser) parseIfStmt() ast.Stmt {
	t, _ := p.require(token.TokenIF)

	return p.parseIfTail(t)
}

func (p *Parser) parseIfTail(t token.Token) ast.Stmt {
	expr := p.parseExpr()

	if _, err := p.require(token.TokenTHEN); err != nil {
		p.findNextSyncPoint(token.TokenEND)
	}

	var trueBlock, falseBlock ast.Stmt

	trueBlock = p.parseStmtSequence(token.TokenELSIF, token.TokenELSE, token.TokenEND)

	switch p.currentType() {
	case token.TokenELSIF:
		e, _ := p.require(token.TokenELSIF)

		falseBlock = p.parseIfTail(e)
	case token.TokenELSE:
		p.consume(token.TokenELSE)

		falseBlock = p.parseStmtSequence(token.TokenEND)

		p.consume(token.TokenEND)
	case token.TokenEND:
		p.consume(token.TokenEND)

		falseBlock = ast.NewStmtSequence(nil)
	default:
		p.out.Errorf(p.currentToken(), "unexpected token after `IF`: %v", p.currentType())
		p.out.Infof(p.currentToken(), "expected `ELSIF`, `ELSE` or `END`")

		p.findNextSyncPoint(token.TokenEND)
	}

	return ast.NewIfStmt(t, expr, trueBlock, falseBlock)
}

func (p *Parser) parseRepeatStmt() ast.Stmt {
	t, _ := p.require(token.TokenREPEAT)

	stmts := p.parseStmtSequence(token.TokenUNTIL)

	if _, err := p.require(token.TokenUNTIL); err != nil {
		p.findNextSyncPoint(token.TokenSemicolon, token.TokenEND)
	}

	expr := p.parseExpr()

	cond := ast.NewCondition(expr, stmts)

	return ast.NewRepeatStmt(t, cond)
}

func (p *Parser) parseWhileStmt() ast.Stmt {
	t, _ := p.require(token.TokenWHILE)

	var conds []ast.Condition

	for {
		expr := p.parseExpr()

		if _, err := p.require(token.TokenDO); err != nil {
			p.findNextSyncPoint(token.TokenELSIF, token.TokenEND, token.TokenSemicolon)
		}

		stmts := p.parseStmtSequence(token.TokenELSIF, token.TokenEND)

		conds = append(conds, ast.NewCondition(expr, stmts))

		if !p.expect(token.TokenELSIF) {
			break
		}
	}

	if _, err := p.require(token.TokenEND); err != nil {
		p.findNextSyncPoint(token.TokenSemicolon)
	}

	return ast.NewWhileStmt(t, conds)
}

func (p *Parser) parseForStmt() ast.Stmt {
	t, _ := p.require(token.TokenFOR)

	iter, err := p.require(token.TokenIdent)
	if err != nil {
		p.findNextSyncPoint(token.TokenAssign, token.TokenTO, token.TokenBY, token.TokenDO, token.TokenEND, token.TokenSemicolon)
	}

	if _, err := p.require(token.TokenAssign); err != nil {
		p.findNextSyncPoint(token.TokenTO, token.TokenBY, token.TokenDO, token.TokenEND, token.TokenSemicolon)
	}

	from := p.parseExpr()

	if _, err := p.require(token.TokenTO); err != nil {
		p.findNextSyncPoint(token.TokenBY, token.TokenDO, token.TokenEND, token.TokenSemicolon)
	}

	to := p.parseExpr()

	var by ast.Expr

	if p.expect(token.TokenBY) {
		by = p.parseExpr()
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
		p.findNextSyncPoint(token.TokenEND, token.TokenSemicolon)
	}

	stmt := p.parseStmtSequence(token.TokenEND)

	if _, err := p.require(token.TokenEND); err != nil {
		p.findNextSyncPoint(token.TokenSemicolon)
	}

	return ast.NewForStmt(t, iter, from, to, by, stmt)
}

func (p *Parser) parseCallExpr(des *ast.DesignatorExpr) ast.Expr {
	p.consume(token.TokenLParen)

	var args []ast.Expr

	for p.currentType() != token.TokenRParen {
		expr := p.parseExpr()

		args = append(args, expr)

		if !p.expect(token.TokenComma) {
			break
		}
	}

	if _, err := p.require(token.TokenRParen); err != nil {
		p.findNextSyncPoint(token.TokenSemicolon, token.TokenEND)
	}

	return ast.NewCallExpr(des, args)
}

func (p *Parser) parseExpr() ast.Expr {
	// expr := simpleExpr [ '=' | '#' | '<' | '<=' | '>' | '>=' | 'IN' | 'IS'  simpleExpr ]
	lhs := p.parseSimpleExpr()

	for op := p.currentToken(); op.IsRelation(); op = p.currentToken() {
		p.index++

		rhs := p.parseSimpleExpr()
		lhs = ast.NewBinaryExpr(op, lhs, rhs)
	}

	return lhs
}

func (p *Parser) parseSimpleExpr() ast.Expr {
	// simpleExpr := term { '+' | '-' | 'OR' term }
	addOperators := []token.TokenType{
		token.TokenPlus, token.TokenMinus, token.TokenOR,
	}

	lhs := p.parseTerm()

	for op := p.currentToken(); p.tokenIs(addOperators...); op = p.currentToken() {
		p.index++

		rhs := p.parseTerm()
		lhs = ast.NewBinaryExpr(op, lhs, rhs)
	}

	return lhs
}

func (p *Parser) parseTerm() ast.Expr {
	// term := factor { '*' | '/' | 'DIV' | 'MUL' | '&' factor }
	mulOperators := []token.TokenType{
		token.TokenAsterisk, token.TokenSlash, token.TokenDIV, token.TokenMOD, token.TokenAmpersand,
	}

	lhs := p.parseFactor()

	for op := p.currentToken(); p.tokenIs(mulOperators...); op = p.currentToken() {
		p.index++

		rhs := p.parseFactor()
		lhs = ast.NewBinaryExpr(op, lhs, rhs)
	}

	return lhs
}

func (p *Parser) parseFactor() ast.Expr {
	// factor := set | designator [actualParameters] | number | string | boolean | '(' expr ')' | '~' factor
	switch p.currentType() {
	case token.TokenIdent:
		des := p.parseDesignator()

		if p.currentType() == token.TokenLParen {
			return p.parseCallExpr(des)
		}

		return des
	case token.TokenInteger, token.TokenReal, token.TokenMinus, token.TokenPlus:
		return p.parseNumberExpr()
	case token.TokenString:
		return p.parseStringLiteral()
	case token.TokenBoolean:
		return p.parseBooleanLiteral()
	case token.TokenTilde:
		return p.parseNotExpr()
	case token.TokenLParen:
		p.consume(token.TokenLParen)

		expr := p.parseExpr()

		if _, err := p.require(token.TokenRParen); err != nil {
			p.findNextSyncPoint(token.TokenRParen, token.TokenSemicolon, token.TokenEND)
		}

		return expr
	case token.TokenLBrace:
		return p.parseSetLiteral()
	default:
		t := p.currentToken()

		p.out.Errorf(t, "unexpected token in factor: %s", t.Text)
		p.out.Infof(t, "expected an expression")
		p.findNextSyncPoint(token.TokenSemicolon, token.TokenEND)

		return ast.NewInvalidExpr(t)
	}
}

func (p *Parser) parseDesignator() *ast.DesignatorExpr {
	ident, _ := p.require(token.TokenIdent)

	des := ast.NewDesignatorExpr(ident)

	for done := false; !done; {
		switch p.currentType() {
		case token.TokenDot:
			t, _ := p.require(token.TokenDot)
			i, _ := p.require(token.TokenIdent)

			des.AddSelector(ast.NewDotSelector(t, i))
		case token.TokenLBracket:
			t, _ := p.require(token.TokenLBracket)

			var exprs []ast.Expr
			for {
				exprs = append(exprs, p.parseExpr())

				if !p.expect(token.TokenComma) {
					break
				}
			}

			_, _ = p.require(token.TokenRBracket)

			des.AddSelector(ast.NewBracketSelector(t, exprs))
		case token.TokenCaret:
			t, _ := p.require(token.TokenCaret)

			des.AddSelector(ast.NewDerefSelector(t))
		case token.TokenLParen:
			// TODO(daniel): ambiguous, as type-guards look like calls.
			done = true
		default:
			done = true
		}
	}

	return des
}

func (p *Parser) parseNumberExpr() ast.Expr {
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
		p.out.Errorf(p.currentToken(), "expected number, got %q", p.currentToken().Text)

		return ast.NewInvalidExpr(p.currentToken())
	}

	num := ast.Expr(ast.NewNumberExpr(t))

	if op.Type == token.TokenMinus {
		// NOTE(daniel): There is no negate for integers in LLVM IR, so we treat
		// `-x` as `0 - x`. That means we don't need generator code either.
		t.Int = 0
		t.Text = ""

		zero := ast.NewNumberExpr(t)

		num = ast.NewBinaryExpr(op, zero, num)
	}

	return num
}

func (p *Parser) parseStringLiteral() ast.Expr {
	t, _ := p.require(token.TokenString)

	if len(t.Text) == 1 {
		return ast.NewCharExpr(t)
	} else {
		return ast.NewStringExpr(t)
	}
}

func (p *Parser) parseBooleanLiteral() ast.Expr {
	t, _ := p.require(token.TokenBoolean)

	return ast.NewBooleanExpr(t)
}

func (p *Parser) parseSetLiteral() ast.Expr {
	t, _ := p.require(token.TokenLBrace)

	var bits []byte

	terminators := []token.TokenType{token.TokenRBrace, token.TokenSemicolon, token.TokenEND}

	for !p.tokenIs(terminators...) {
		i1, err := p.require(token.TokenInteger)
		if err != nil {
			p.findNextSyncPoint(token.TokenRBrace, terminators...)

			return ast.NewInvalidExpr(t)
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
			p.findNextSyncPoint(token.TokenRBrace, terminators...)

			return ast.NewInvalidExpr(t)
		}

		for i := i1.Int; i <= i2.Int; i++ {
			bits = append(bits, byte(i))
		}

		if p.expect(token.TokenComma) {
			continue
		}
	}

	if _, err := p.require(token.TokenRBrace); err != nil {
		p.findNextSyncPoint(token.TokenRBrace, terminators...)
	}

	return ast.NewSetExpr(t, bits)
}

func (p *Parser) parseNotExpr() ast.Expr {
	t, _ := p.require(token.TokenTilde)

	expr := p.parseExpr()

	return ast.NewNotExpr(t, expr)
}

func (p *Parser) findNextSyncPoint(arg token.TokenType, args ...token.TokenType) {
	args = append(args, arg)

	for p.index < len(p.tokens)-1 {
		if p.tokenIs(args...) {
			break
		}

		p.index++
	}
}

func (p *Parser) require(exp token.TokenType) (token.Token, error) {
	if p.currentType() == exp {
		token := p.currentToken()
		p.index++

		return token, nil
	}

	p.out.Errorf(p.currentToken(), "unexpected token: %v", p.currentToken().Text)
	p.out.Infof(p.currentToken(), "expected: %v", exp)

	return token.Token{}, ErrParsing
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
