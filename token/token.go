package token

import (
	"fmt"
)

type Token struct {
	Type    TokenType
	Literal string
}

// Returns a string representation of a token
func (t Token) String() string {
	var stringToken string
	switch t.Type {
	case INTEGER, IDENTIFIER, ILLEGAL:
		stringToken = fmt.Sprintf("%s(%s)", t.Type.String(), t.Literal)
	default:
		stringToken = t.Type.String()
	}
	return stringToken
}

// Creates a new Token given its type and literal string
func New(tokenType TokenType, literal string) Token {
	return Token{Type: tokenType, Literal: literal}
}

func DetermineTokenType(id string) TokenType {
	if tokenType, found := keywords[id]; found {
		return tokenType
	}
	return IDENTIFIER
}
