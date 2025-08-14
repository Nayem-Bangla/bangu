package ast

import "bangu/token"

type Node interface { // TokenLiteral returns the literal value of the token associated with the node.
	TokenLiteral() string
}

type Statement interface {
	Node
	// statementNode is a marker method to distinguish statements from expressions.
	statementNode()
}

type Expression interface {
	Node
	// expressionNode is a marker method to distinguish expressions from statements.
	expressionNode()
}

type Program struct {
	Statements []Statement // Statements is a slice of Statement nodes.
}

type ReturnStatement struct {
	Token       token.Token // The token.RETURN token.
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// TokenLiteral returns the literal value of the token associated with the Program node.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type LetStatement struct {
	Token token.Token // The token.LET token.
	Name  *Identifier // The variable name being declared.
	Value Expression  // The value being assigned to the variable.
}

// TokenLiteral returns the literal value of the token associated with the LetStatement node.
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// statementNode is a marker method to distinguish LetStatement as a statement.
func (ls *LetStatement) statementNode() {}

type Identifier struct {
	Token token.Token // The token.IDENT token.
	Value string      // The identifier's name.
}

// TokenLiteral returns the literal value of the token associated with the Identifier node.
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// expressionNode is a marker method to distinguish Identifier as an expression.
func (i *Identifier) expressionNode() {}
