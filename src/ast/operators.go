package ast

import (
	"fmt"
	"turtlego/src/tokens"
)

type PrefixExpr struct {
	Expr Node
	Op   string
	Tok  tokens.Token
	Type TypeInfo
}

func (p *PrefixExpr) PrintAll(tab string) {
	fmt.Print(p.Stringify(tab))
}
func (p *PrefixExpr) Stringify(tab string) string {
	rtrnVal := ""
	rtrnVal += fmt.Sprintf("%s<Prefix Expression Operator='%s'>\n", tab, p.Op)
	rtrnVal += p.Expr.Stringify(tab + "\t")
	rtrnVal += tab + "<\\Prefix Expression>\n"
	return rtrnVal
}

func (p *PrefixExpr) GetTok() tokens.Token {
	return p.Tok
}

func (p *PrefixExpr) TypeGenerated() TypeInfo {
	return p.Type
}

func (p *PrefixExpr) NodeType() byte {
	return PREFIX_NT
}

/////////////////////////////////////////////////

type InfixExpr struct {
	Left  Node
	Right Node
	Op    string
	Tok   tokens.Token
	Type  TypeInfo
}

func (i *InfixExpr) PrintAll(tab string) {
	fmt.Print(i.Stringify(tab))
}
func (i *InfixExpr) Stringify(tab string) string {
	rtrnStr := fmt.Sprintf("%s<Infix Expression Operator='%s'>\n", tab, i.Op)

	rtrnStr += tab + "<Left>\n"
	rtrnStr += i.Left.Stringify(tab + "\t")
	rtrnStr += tab + "<\\Left>\n"

	rtrnStr += tab + "<Right>\n"
	rtrnStr += i.Right.Stringify(tab + "\t")
	rtrnStr += tab + "<\\Right>\n"

	rtrnStr += tab + "<\\Infix Expression>\n"
	return rtrnStr
}

func (i *InfixExpr) GetTok() tokens.Token {
	return i.Tok
}
func (i *InfixExpr) TypeGenerated() TypeInfo {
	return i.Type
}

func (i *InfixExpr) NodeType() byte {
	return INFIX_NT
}
