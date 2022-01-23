package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}


const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifier + Literal
	IDENT = "IDENT"
	INT   = "INT"

	// Operator
	ASSIGN = "="
	PLUS   = "+"

	// Sperator
	COMMA     = ","
	SEMICOLON = ";"

	// bracket
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Reserved keyword
	FUNCTION = "FUNCTION"
	LET      = "LET"
)