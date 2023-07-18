package lexer

import (
	"intergo/pkg/token"
)

type Lexer struct {
	input    string
	position int
	readPos  int
	ch       byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.position = l.readPos
	l.readPos += 1
}
func (l *Lexer) ignoreWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.ignoreWhiteSpace()
	switch l.ch {
	case '=':
		if l.peek() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Value: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.EQUALS, l.ch)
		}
	case '+':
		tok = newToken(token.ADD, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '-':
		tok = newToken(token.SUB, l.ch)
	case '!':
		if l.peek() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Value: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.EXC, l.ch)
		}
	case '/':
		tok = newToken(token.FWDSLASH, l.ch)
	case '*':
		tok = newToken(token.AST, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMI_COLON, l.ch)
	case 0:
		tok.Value = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Value = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Value)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Value = l.readNum()
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Value: string(ch)}
}
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
func (l *Lexer) readNum() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
func (l *Lexer) peek() byte {
	if l.readPos >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPos]
	}
}
