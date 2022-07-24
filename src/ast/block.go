package ast

import (
	"fmt"
	"turtlego/src/tokens"
)

type Block struct {
	Exprs   []Node
	RetType byte
	Tok     tokens.Token
}

func (b *Block) GetTok() tokens.Token {
	return b.Tok
}

func (b *Block) TypeGenerated() byte {
	return b.RetType
}

func (b *Block) Stringify(tab string) string {
	ret_val := tab + "<block>\n"
	for _, i := range b.Exprs {
		ret_val += i.Stringify(tab + "\t")
	}
	ret_val += tab + "<\\block>\n"
	return ret_val
}

func (b *Block) PrintAll(tab string) {
	fmt.Printf(b.Stringify(tab))
}

func (b *Block) NodeType() byte {
	return BLOCK_NT
}
