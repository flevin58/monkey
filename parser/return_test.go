package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestReturnStatements(t *testing.T) {
	input := ` return 5; return 10; return 993322; `
	l := lexer.NewFromString(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got:%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}
