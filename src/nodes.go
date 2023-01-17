package main

type Visitor interface {
	VisitModule(*Module)
	VisitPrintStmt(*PrintStmt)
	VisitBinaryExpr(*BinaryExpr)
	VisitNumberExpr(*NumberExpr)
}

type Node interface {
	Token() Token
	Visit(Visitor)
}

type Module struct {
	stmts []Node
}

var _ Node = (*Module)(nil)

func (n *Module) Token() Token {
	if len(n.stmts) == 0 {
		return Token{}
	}

	return n.stmts[0].Token()
}

func (n *Module) Visit(v Visitor) {
	v.VisitModule(n)
}

type NumberExpr struct {
	token Token
}

var _ Node = (*NumberExpr)(nil)

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

var _ Node = (*BinaryExpr)(nil)

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

var _ Node = (*PrintStmt)(nil)

func (n *PrintStmt) Token() Token {
	return n.token
}

func (n *PrintStmt) Visit(v Visitor) {
	v.VisitPrintStmt(n)
}
