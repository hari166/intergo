package token

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

const (
	SEMI_COLON = ";"
	EQUALS     = "="
	COMMA      = ","
	ADD        = "+"
	SUB        = "-"
	LT         = "<"
	GT         = ">"
	EXC        = "!"
	AST        = "*"
	FWDSLASH   = "/"
	NOT_EQ     = "!="
	EQ         = "=="

	FUNCTION = "FX"
	INT      = "INT"
	IDENT    = "IDENT"
	LET      = "LET"
	IF       = "IF"
	RETURN   = "RETURN"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

var keywords = map[string]TokenType{
	"fx":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
