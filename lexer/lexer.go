package lexer

import "go-interpreter/token"

type Lexer struct {
	input        string
	position     int  // current position in input (it points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	// [...]
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier() // if is a letter at start, delegate it to a identifier function
			return tok
		} else { // otherwise is illegal
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
}
