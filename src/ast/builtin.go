package ast

import (
	"fmt"
	"turtlego/src/tokens"
)

type Builtin struct {
	Args    []Node
	Name    string
	Tok     tokens.Token
	RetType byte
}

func (b *Builtin) GetTok() tokens.Token {
	return b.Tok
}

func (b *Builtin) TypeGenerated() byte {
	return b.RetType
}

func (b *Builtin) NodeType() byte {
	return BUILTIN_NT
}

func (b *Builtin) Stringify(tab string) string {
	retVal := tab + "<builtin='" + b.Name + "'>\n"

	for _, i := range b.Args {
		retVal += i.Stringify(tab + "\t")
	}

	retVal += tab + "</builtin>\n"
	return retVal
}

func (b *Builtin) PrintAll(tab string) {
	fmt.Print(b.Stringify(tab))
}
