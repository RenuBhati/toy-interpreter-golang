package ast

import (
	"monkey/token"
)

// Here we have three interfaces called Node, Statement and Expression.
// Every node in our AST has to implement the Node interface,
//meaning it has to provide a TokenLiteral() method that returns the literal value of the token
// it’s associated with.

type Node interface {
	TokenLiteral() string  //TokenLiteral() method that returns the literal value of the token it’s associated with.
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type Identifier struct {
	Token token.Token
	Value string
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

