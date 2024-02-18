package ast

import (
	"barracuda/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name:  &Identifier{Token: token.Token{Type: token.IDENTIF, Literal: "myVar"}, Value: "myVar"},
				Value: &Identifier{Token: token.Token{Type: token.IDENTIF, Literal: "anotherVar"}, Value: "anotherVar"},
			},
		},
	}
	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. Got=%q", program.String())
	}
}
