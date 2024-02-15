package lexer

import (
	"barracuda/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	// has two whitespaces in a row
	input :=
		`
		let five  = 5;
		let ten = 10;
		
		let add = fn(x,y){
			x+y};

		let result = add(five, ten);

		!-/*5;
		5 < 10 > 5;

		if (5 < 10) {
			return true;
		} else {
			return false;
		}

		10 == 10;
		10 != 9;

			`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// first statement
		{token.LET, "let"},
		{token.IDENTIF, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		// second statement
		{token.LET, "let"},
		{token.IDENTIF, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		// third statement
		{token.LET, "let"},
		{token.IDENTIF, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENTIF, "x"},
		{token.COMMA, ","},
		{token.IDENTIF, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIF, "x"},
		{token.PLUS, "+"},
		{token.IDENTIF, "y"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		// fourth statement
		{token.LET, "let"},
		{token.IDENTIF, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIF, "add"},
		{token.LPAREN, "("},
		{token.IDENTIF, "five"},
		{token.COMMA, ","},
		{token.IDENTIF, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		// fifth "statement"
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		// ...
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong, expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - token literal wrong, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
