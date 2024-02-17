package ast

import "barracuda/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node // Expression embeds Node interface
	expressionNode()
}

// root node of every AST produced by parser
type Program struct {
	// every program is a series of statements
	Statements []Statement
}

type LetStatement struct {
	Token token.Token // the LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token // the IDENTIF token
	Value string
}

func (i *Identifier) expressionNode() {} // to simplify parsing Identifier implements Expression interface
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
