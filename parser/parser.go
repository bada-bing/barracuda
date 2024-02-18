package parser

import (
	"barracuda/ast"
	"barracuda/lexer"
	"barracuda/token"
	"fmt"
)

type Parser struct {
	l *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token

	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	// read two tokens so that both current and peek token are set
	p.NextToken()
	p.NextToken()

	return p

}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) NextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currentToken}

	if !p.expectPeekTokenType(token.IDENTIF) {
		return nil
	}
	// else if ok expectPeekTokenType reads NextToken()
	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.expectPeekTokenType(token.ASSIGN) {
		return nil
	}
	// TODO: we are skipping expressions until we encounter SEMICOLON
	for !p.currentTokenIs(token.SEMICOLON) {
		p.NextToken()
	}
	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currentToken}

	p.NextToken()

	// current tests will timeout if there is no semicolon at the end
	for !p.currentTokenIs(token.SEMICOLON) {
		p.NextToken()
	}

	return stmt
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{} // pointer of new Program instance
	program.Statements = []ast.Statement{}

	for p.currentToken.Type != token.EOF {
		stmt := p.ParseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.NextToken()
	}

	return program
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) expectPeekTokenType(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.NextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
