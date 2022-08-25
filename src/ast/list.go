package ast

import (
	"fmt"
	"turtlego/src/tokens"
)

type List struct {
	Type   ListType
	Values []Node
	Tok    tokens.Token
}

func (l *List) GetTok() tokens.Token {
	return l.Tok
}

func (l *List) TypeGenerated() TypeInfo {
	return l.Type
}

func (l *List) NodeType() byte {
	return LIST_NT
}

func (l *List) Stringify(tab string) string {
	result := tab + "<list>\n"
	for _, v := range l.Values {
		result += v.Stringify(tab + "\t")
	}
	result += tab + "</list>\n"

	return result
}

func (l *List) PrintAll(tab string) {
	fmt.Print(l.Stringify(tab))
}
