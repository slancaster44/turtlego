package ast

import "turtlego/tokens"

type Node interface {
	Type() int
	SetId(int)
	Id() int
	Token() tokens.Token
}

type Program []Node

func GetStaticCounter() func() int {
	i := 0
	f := func() int {
		i = i + 1
		return i
	}
	return f
}
