package main

import (
	"fmt"
	"os"
	"turtlego/src/lexer"
	"turtlego/src/parser"
	"turtlego/src/pcodegnerator"
	"turtlego/src/source"
)

var version = "v0.2.0 -- Development"

func main() {
	if len(os.Args) < 2 {
		printHelp()
	}

	switch os.Args[1] {
	case "-l":
		printToks()
	case "-p":
		printTree()
	case "-c":
		printPCode()
	default:
		printHelp()
	}
}

func printHelp() {
	msg := `
Welcome to the TurtleGo Turtle Compiler
Version Information: ` + version + `

    turtle [option] | turtle [option] [file]
    -h :: print this message and exit
    -l :: lex file, and print tokens
    -p :: parse file and print tree
    -e :: complete evaluation of the file
    `
	fmt.Println(msg)
	os.Exit(0)
}

func printToks() {
	if len(os.Args) < 3 {
		printHelp()
	}

	src := source.New(os.Args[2])
	lex := lexer.New(src)

	for !lex.IsDone() {
		fmt.Printf("%v -> %v\n", lex.CurTok.Alias, lex.CurTok)
		lex.MoveUp()
	}
}

func printTree() {
	if len(os.Args) < 3 {
		printHelp()
	}

	src := source.New(os.Args[2])
	lex := lexer.New(src)
	prs := parser.New(lex)
	prs.ParseProgram()

	for _, i := range prs.Tree {
		i.PrintAll("")
	}
}

func printPCode() {
	if len(os.Args) < 3 {
		printHelp()
	}

	src := source.New(os.Args[2])
	lex := lexer.New(src)
	prs := parser.New(lex)
	prs.ParseProgram()

	pg := pcodegnerator.NewGenerator(prs.Tree)
	pg.GenPCode()

	for _, i := range pg.Output.Instructions {
		i.Print()
	}
}
