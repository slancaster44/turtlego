package pcodegenerator

import (
	"turtlego/src/ast"
	"turtlego/src/pcode"
)

func (g *Generator) genLetInit(n ast.Node) Register {
	node := n.(*ast.LetInit)
	result_reg := g.appendCodeFor(node.Expr)

	location_reg := g.GetRegister()
	ins := &pcode.Instruction{pcode.MOV_REG_REG, []int{location_reg.RegisterNumber, pcode.STACK_FRAME_POINTER_REG}}
	g.Program.WriteInstruction(ins)

	ins = &pcode.Instruction{pcode.ADD_REG_INT_INT,
		[]int{location_reg.RegisterNumber, node.LocationInfo.LocationOnStack * STACK_VAR_SIZE}}
	g.Program.WriteInstruction(ins)

	ins = &pcode.Instruction{pcode.MOV_REG_ADDRESS_REG, []int{location_reg.RegisterNumber, result_reg.RegisterNumber}}
	g.Program.WriteInstruction(ins)

	g.ReleaseRegister(location_reg)

	g.SymTab[node.Ident] = node.LocationInfo

	return result_reg
}
