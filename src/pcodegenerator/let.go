package pcodegenerator

import (
	"turtlego/src/ast"
	"turtlego/src/pcode"
)

func (g *Generator) genLetInit(n ast.Node) Register {
	node := n.(*ast.LetInit)
	result_reg := g.appendCodeFor(node.Expr)

	location_reg := g.GetRegister()
	g.WriteInstruction(pcode.MOV_REG_REG, location_reg.RegisterNumber, pcode.STACK_FRAME_POINTER_REG)

	loc_on_stack := node.LocationInfo.LocationOnStack
	g.WriteInstruction(pcode.ADD_REG_INT_INT, location_reg.RegisterNumber, loc_on_stack*STACK_VAR_SIZE)

	g.WriteInstruction(pcode.MOV_REG_ADDRESS_REG, location_reg.RegisterNumber, result_reg.RegisterNumber)

	g.ReleaseRegister(location_reg)

	g.SymTab[node.Ident] = node.LocationInfo

	return result_reg
}
