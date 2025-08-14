package parser

import (
	"bangu/ast"
	"bangu/lexer"
	"bangu/token"
	"fmt"
)

// Parser is responsible for parsing tokens into an AST.
type Parser struct {
	lexer     *lexer.Lexer
	curToken  token.Token // The current token being parsed.
	peekToken token.Token // The next token to be parsed.
	errors    []string    // A slice to hold any parsing errors encountered.
}

// New creates a new Parser instance.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l,

		errors: []string{},
	}

	// Read two tokens, so curToken and peekToken are both set.
	// This allows the parser to look ahead one token.
	// The first call to nextToken initializes curToken.
	// The second call initializes peekToken.
	// This is necessary to avoid parsing errors when the parser expects a token.
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string {
	// Returns the slice of parsing errors encountered.
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	// Adds an error message to the parser's errors slice.
	msg := fmt.Sprintf("expected next taken to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	// Advance the lexer to the next token and update curToken and peekToken.
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

// ParseProgram parses the input tokens into an AST Program.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// Continue parsing until we reach the end of the input.
	// This loop will parse all statements in the program.
	// Each statement is parsed and added to the program's Statements slice.
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// PTODO: We're skipping the expression until we encounter a semicolon.
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	// Expect the next token to be an identifier.
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// Create a new Identifier node for the variable name.
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expression until we encounter a semicolon.
	for !p.curTokenIs(token.SEMICOLON) {
		// Advance to the next token until we reach a semicolon.
		p.nextToken()
	}

	// Return the completed LetStatement.
	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
