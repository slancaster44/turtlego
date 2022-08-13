package ast

import (
	"fmt"
	"turtlego/src/tokens"
)

//Notes on parsing
//Both TrueExpr and FalseExpr must return the same type
//Cond must return a boolean

type IfElse struct {
	Cond      Node
	TrueExpr  Node
	FalseExpr Node
	Tok       tokens.Token
}

func (i *IfElse) GetTok() tokens.Token {
	return i.Tok
}

func (i *IfElse) NodeType() byte {
	return IFEL_NT
}

func (i *IfElse) TypeGenerated() byte {
	return i.TrueExpr.TypeGenerated()
}

func (i *IfElse) Stringify(tab string) string {
	ret_val := tab + "<if/else>\n"

	ret_val += tab + "<condition>\n"

	ret_val += i.Cond.Stringify(tab + "\t")

	ret_val += tab + "</condition>\n"
	ret_val += tab + "<true_expr>\n"

	ret_val += i.TrueExpr.Stringify(tab + "\t")

	ret_val += tab + "</true_expr>\n"
	ret_val += tab + "<false_expr>\n"

	ret_val += i.FalseExpr.Stringify(tab + "\t")

	ret_val += tab + "</false_expr>\n"
	ret_val += tab + "</if/else>\n"

	return ret_val
}

func (i *IfElse) PrintAll(tab string) {
	fmt.Print(i.Stringify(tab))
}
