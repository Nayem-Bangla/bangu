package ast

import (
	"bangu/token"
	"bytes"
)

type Node interface { // TokenLiteral returns the literal value of the token associated with the node.
	TokenLiteral() string
	String() string
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

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

// TokenLiteral returns the literal value of the token associated with the ExpressionStatement node.
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// statementNode is a marker method to distinguish ExpressionStatement as a statement.
func (es *ExpressionStatement) statementNode() {}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// Stirng returns a string representation of the Program node.
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (i *Identifier) String() string {
	return i.Value
}
