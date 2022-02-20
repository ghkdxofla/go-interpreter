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
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

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

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookUpIdent(ident string) TokenType {
	/*
	 * 참고
	 * Golang에서는 아래와 같이 map의 key 존재 여부를 확인할 수 있다
	 * if i, boolVal := map[key]; boolVal 식으로 사용하면
	 * key가 있을 때: i에는 map[key]에 해당하는 값, boolVal에는 true
	 * key가 없을 때: i에는 0, boolVal에는 false
	 * Link: https://stackoverflow.com/questions/52705926/golang-syntax-in-if-statement-with-a-map
	 */
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
