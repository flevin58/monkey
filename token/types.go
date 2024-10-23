package token

type TokenType int

const (
	ILLEGAL TokenType = iota
	EOF

	// Identifiers
	IDENTIFIER
	INTEGER
	STRING

	// Operators
	ASSIGN
	PLUS
	MINUS
	ASTERISK
	SLASH

	// Logical
	EQ
	NEQ
	BANG
	GT
	LT

	// Delimiters
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE

	// Keywords
	FUNCTION
	LET
	RETURN

	// Other
	COMMENT
	TRUE
	FALSE
	IF
	ELSE
)

func (t TokenType) String() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case IDENTIFIER:
		return "IDENTIFIER"
	case INTEGER:
		return "INTEGER"
	case STRING:
		return "STRING"
	case COMMENT:
		return "COMMENT"
	case ASSIGN:
		return "ASSIGN"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case ASTERISK:
		return "ASTERISK"
	case SLASH:
		return "SLASH"
	case EQ:
		return "EQ"
	case NEQ:
		return "NEQ"
	case BANG:
		return "BANG"
	case GT:
		return "GT"
	case LT:
		return "LT"
	case COMMA:
		return "COMMA"
	case SEMICOLON:
		return "SEMICOLON"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case LBRACE:
		return "LBRACE"
	case RBRACE:
		return "RBRACE"
	case FUNCTION:
		return "FUNCTION"
	case LET:
		return "LET"
	case RETURN:
		return "RETURN"
	case TRUE:
		return "TRUE"
	case FALSE:
		return "FALSE"
	case IF:
		return "IF"
	case ELSE:
		return "ELSE"
	default:
		return ""
	}
}
