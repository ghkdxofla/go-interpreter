package lexer

import "go-interpreter/token"

// TODO: ch의 byte를 rune으로 변경하면 유니코드, utf-8 지원 가능함
type Lexer struct {
	input        string
	position     int  // 현재 위치
	readPosition int  // 다음 위치
	ch           byte // 현재 문자
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case 0:
		tok = newToken(token.EOF, l.ch)
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpIdent(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	var literal string

	if ch == 0 {
		literal = ""
	} else {
		literal = string(ch)
	}

	return token.Token{Type: tokenType, Literal: literal}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII에서 'NUL'에 해당. 읽기 전 상태 또는 EOF를 의미.
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
	position := l.position // 글자의 처음 시작점 저장
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position] // 글자 추출
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' // 이 식 대로면 var-iable 등의 변수명은 불가
}