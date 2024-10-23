package main

import (
	"fmt"
	"monkey/lexer"
	"monkey/repl"
	"monkey/token"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		// Here we give the filename in the command line
		l, err := lexer.NewFromFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	} else {
		repl.Start(os.Stdin, os.Stdout)
	}
}
