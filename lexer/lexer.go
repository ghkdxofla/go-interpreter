package lexer

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

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII에서 'NUL'에 해당. 읽기 전 상태 또는 EOF를 의미.
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
