package pcodegenerator

import (
	"turtlego/src/ast"
	"turtlego/src/pcode"
)

const SIZE_OF_LIST_METADATA int = 2

func (g *Generator) genListInit(n ast.Node) Register {
	ret_reg := g.GetRegister()

	length_of_list := len(n.(*ast.List).Values) + SIZE_OF_LIST_METADATA

	g.WriteInstruction(pcode.ALLOC, ret_reg.RegisterNumber, length_of_list)

	return ret_reg
}
