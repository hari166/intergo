package lexer

import (
	"intergo/pkg/token"
	"testing"
)

func TestNext(t *testing.T) {
	input := `
	let five = 5
	let ten = 10
	let add = fx(x, y) {
		x + y
	}
	
	let result = add(five, ten)
	!-/*5
	5 < 10 > 5
	if (5 < 10 ) {
		return true 
	}
	else {
		return false
	}
	10 == 10
	10 != 9
	`

	tests := []struct {
		expectedType  token.TokenType
		expectedValue string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.EQUALS, "="},
		{token.INT, "5"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.EQUALS, "="},
		{token.INT, "10"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.EQUALS, "="},
		{token.FUNCTION, "fx"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.ADD, "+"},
		{token.IDENT, "y"},
		{token.RBRACE, "}"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.EQUALS, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.EXC, "!"},
		{token.SUB, "-"},
		{token.FWDSLASH, "/"},
		{token.AST, "*"},
		{token.INT, "5"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		//fmt.Printf("Token %d: Expected=%q, Got=%q\n", i, tt.expectedValue, tok.Value)
		if tok.Type != tt.expectedType {
			t.Fatalf("Tests[%d]: Expected %q, got %q", i, tt.expectedType, tok.Type)
		}
		if tok.Value != tt.expectedValue {
			t.Fatalf("Tests[%d]: Expected %q, got %q", i, tt.expectedValue, tok.Value)
		}
	}

}
