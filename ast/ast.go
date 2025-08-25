package ast

import (
	"bangu/token"
	"bytes"
	"strings"
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

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

type PrefixExpression struct {
	Token    token.Token // The prefix token, e.g., '!' or '-'.
	Operator string      // The operator, e.g., '!' or '-'.
	Right    Expression  // The expression that follows the operator.
}

// expressionNode is a marker method to distinguish PrefixExpression as an expression.
func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral returns the literal value of the token associated with the PrefixExpression node.
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

// String returns a string representation of the PrefixExpression node.
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

type InfixExpression struct {
	Token    token.Token // The infix token, e.g., '+', '-', '*', '/'.
	Left     Expression  // The left-hand side expression.
	Operator string      // The operator, e.g., '+', '-', '*', '/'.
	Right    Expression  // The right-hand side expression.
}

// expressionNode is a marker method to distinguish InfixExpression as an expression.
func (oe *InfixExpression) expressionNode() {}
func (oe *InfixExpression) TokenLiteral() string {
	return oe.Token.Literal
}
func (oe *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")
	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

// expressionNode is a marker method to distinguish Boolean as an expression.
func (b *Boolean) expressionNode() {}

// TokenLiteral returns the literal value of the token associated with the Boolean node.
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

// String returns a string representation of the Boolean node.
func (b *Boolean) String() string {
	return b.Token.Literal
}

type IfExpression struct {
	Token       token.Token     // The 'if' token.
	Condition   Expression      // The condition expression.
	Consequence *BlockStatement // The block of statements to execute if the condition is true.
	Alternative *BlockStatement // The block of statements to execute if the condition is false (optional).
}

// expressionNode is a marker method to distinguish IfExpression as an expression.
func (ie *IfExpression) expressionNode() {}

// TokenLiteral returns the literal value of the token associated with the IfExpression node.
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String returns a string representation of the IfExpression node.
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if ")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString(" else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

type BlockStatement struct {
	Token      token.Token // The '{' token.
	Statements []Statement // The statements in the block.
}

// statementNode is a marker method to distinguish BlockStatement as a statement.
func (bs *BlockStatement) statementNode() {}

// TokenLiteral returns the literal value of the token associated with the BlockStatement node.
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

// String returns a string representation of the BlockStatement node.
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token     // The 'fn' token.
	Parameters []*Identifier   // The parameters of the function.
	Body       *BlockStatement // The body of the function.
}

// expressionNode is a marker method to distinguish FunctionLiteral as an expression.
func (fl *FunctionLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token associated with the FunctionLiteral node.
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

// String returns a string representation of the FunctionLiteral node.
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")

	out.WriteString(fl.Body.String())

	return out.String()
}

type CallExpression struct {
	Token     token.Token  // The '(' token.
	Function  Expression   // The function being called.
	Arguments []Expression // The arguments passed to the function.
}

// expressionNode is a marker method to distinguish CallExpression as an expression.
func (ce *CallExpression) expressionNode() {}

// TokenLiteral returns the literal value of the token associated with the CallExpression node.
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

// String returns a string representation of the CallExpression node.
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

type StringLiteral struct {
	Token token.Token // The token.STRING token.
	Value string      // The string value.
}

func (sl *StringLiteral) expressionNode() {}
func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Literal
}
func (sl *StringLiteral) String() string {
	return sl.Token.Literal
}

type ArrayLiteral struct {
	Token    token.Token  // The '[' token.
	Elements []Expression // The elements of the array.
}

func (al *ArrayLiteral) expressionNode() {}
func (al *ArrayLiteral) TokenLiteral() string {
	return al.Token.Literal
}
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type IndexExpression struct {
	Token token.Token // The '[' token.
	Left  Expression  // The expression being indexed.
	Index Expression  // The index expression.
}

func (ie *IndexExpression) expressionNode() {}
func (ie *IndexExpression) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}

type HashLiteral struct {
	Token token.Token
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) expressionNode()      {}
func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }
func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
