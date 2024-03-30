package main

import (
	"bufio"
	"log"
	"os"

	// "github.com/HicaroD/telia-lang/ast"
	"github.com/HicaroD/telia-lang/codegen"
	"github.com/HicaroD/telia-lang/lexer"
	"github.com/HicaroD/telia-lang/parser"
	"github.com/HicaroD/telia-lang/sema"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("error: no input files")
	}
	filename := args[0]

	file, err := os.Open(filename)
	// TODO(errors)
	if err != nil {
		log.Fatalf("unable to open file: %s due to error '%s'", filename, err)
	}

	reader := bufio.NewReader(file)

	lex := lexer.NewLexer(filename, reader)
	tokens := lex.Tokenize()
	// for i := range tokens {
	// 	fmt.Printf("%s %s\n", tokens[i].Kind, tokens[i].Lexeme)
	// }

	parser := parser.NewParser(tokens)
	astNodes, err := parser.Parse()
	if err != nil {
		// TODO(errors)
		log.Fatal(err)
	}

	sema := sema.NewSema(astNodes)
	err = sema.Analyze()
	if err != nil {
		log.Fatal(err)
	}

	codegen := codegen.NewCodegen(astNodes)
	err = codegen.Generate()
	if err != nil {
		log.Fatal(err)
	}
}
