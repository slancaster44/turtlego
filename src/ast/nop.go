package ast

import (
	"fmt"
	"turtlego/src/tokens"
)

type Nop struct {
	Tok tokens.Token
}

func (n *Nop) TypeGenerated() byte {
	return NO_TYPE
}

func (n *Nop) GetTok() tokens.Token {
	return n.Tok
}

func (n *Nop) NodeType() byte {
	return NOP_NT
}

func (n *Nop) Stringify(tab string) string {
	return tab + "<nop>"
}

func (n *Nop) PrintAll(tab string) {
	fmt.Print(n.Stringify(tab))
}
