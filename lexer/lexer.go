package lexer

import (
	"bufio"
	"monkey/token"
	"os"
	"strings"
	"text/scanner"
	"unicode"
)

type Lexer struct {
	scanner scanner.Scanner
	line    uint
	col     uint
	char    rune
}

// Returns a pointer to a Lexer from the given string
func NewFromString(inputString string) *Lexer {
	l := &Lexer{}
	l.scanner.Init(strings.NewReader(inputString))
	l.nextRune()
	return l
}

// Returns a pointer to a Lexer from the given file name
func NewFromFile(filePath string) (*Lexer, error) {
	l := &Lexer{}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	l.scanner.Init(bufio.NewReader(file))
	l.nextRune()
	return l, nil
}

// Consumes the next rune and stores it internally in char
func (l *Lexer) nextRune() {
	l.char = l.scanner.Next()
	l.col++
	if l.char == '\n' {
		l.col = 0
		l.line++
	}
}

// Gets the next rune without consuming it
func (l *Lexer) peekRune() rune {
	return l.scanner.Peek()
}

// Reads one or more runes and converts them into Token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.char {
	case '{':
		tok = token.New(token.LBRACE, "{")
	case '}':
		tok = token.New(token.RBRACE, "}")
	case '(':
		tok = token.New(token.LPAREN, "(")
	case ')':
		tok = token.New(token.RPAREN, ")")
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '=':
		if l.peekRune() == '=' {
			l.nextRune()
			tok = token.New(token.EQ, "==")
		} else {
			tok = token.New(token.ASSIGN, "=")
		}
	case '>':
		tok = token.New(token.GT, ">")
	case '<':
		tok = token.New(token.LT, "<")
	case '!':
		if l.peekRune() == '=' {
			l.nextRune()
			tok = token.New(token.NEQ, "!=")
		} else {
			tok = token.New(token.BANG, "!")
		}
	case '-':
		tok = token.New(token.MINUS, "-")
	case '+':
		tok = token.New(token.PLUS, "+")
	case '*':
		tok = token.New(token.ASTERISK, "*")
	case '/':
		switch l.peekRune() {
		case '/':
			l.nextRune()
			tok.Type = token.COMMENT
			tok.Literal = l.readLineComment()
		case '*':
			l.nextRune()
			tok.Type = token.COMMENT
			tok.Literal = l.readMultiLineComment()
		default:
			tok = token.New(token.SLASH, "/")
		}
	case ';':
		tok = token.New(token.SEMICOLON, ";")
	case ',':
		tok = token.New(token.COMMA, ",")
	case scanner.EOF:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.DetermineTokenType(tok.Literal)
			return tok
		} else if isDigit(l.char) {
			tok.Literal = l.readNumber()
			tok.Type = token.INTEGER
			return tok
		} else {
			tok = token.New(token.ILLEGAL, string(l.char))
		}
	}
	l.nextRune()
	return tok
}

func (l *Lexer) readIdentifier() string {
	var result strings.Builder
	for isLetter(l.char) {
		result.WriteRune(l.char)
		l.nextRune()
	}
	return result.String()
}

// Reads consecutive digits and return their string representation
func (l *Lexer) readNumber() string {
	var result strings.Builder
	for isDigit(l.char) {
		result.WriteRune(l.char)
		l.nextRune()
	}
	return result.String()
}

// Reads a single line comment that starts with "//" and ends at the end of the line
func (l *Lexer) readLineComment() string {
	var result strings.Builder
	result.WriteString("//")
	l.nextRune()
	for l.char != '\n' {
		if l.char == scanner.EOF {
			break
		}
		result.WriteRune(l.char)
		l.nextRune()
	}
	return result.String()
}

// Reads a multi line comment that starts with "/*" and ends with "*/"
func (l *Lexer) readMultiLineComment() string {
	var result strings.Builder
	result.WriteString("/*")
	l.nextRune()
	for l.char != scanner.EOF {
		if l.char == '*' && l.peekRune() == '/' {
			result.WriteString("*/")
			l.nextRune()
			break
		}
		result.WriteRune(l.char)
		l.nextRune()
	}
	return result.String()
}

func (l *Lexer) readString() string {
	var result strings.Builder
	l.nextRune()
	for l.char != '"' {
		if l.char == '\\' && l.peekRune() == '"' {
			result.WriteString(`\"`)
			l.nextRune()
		} else {
			result.WriteRune(l.char)
		}
		l.nextRune()
	}
	return result.String()
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.char) {
		l.nextRune()
	}
}
