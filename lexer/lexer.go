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
	l.readChar() // initialize the lexer
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
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()          // if is a letter at start, delegate it to a identifier functio
			tok.Type = token.LookupIdent(tok.Literal) // check if the identifier is a keyword or not and then assign it
			return tok
		} else { // otherwise is illegal
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar() // point the read to the next char
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position // this variable is used to store the position of the first letter of the identifier
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// checks for letters and underscores
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
