package pcodegenerator

import (
	"turtlego/src/ast"
	"turtlego/src/pcode"
)

var typeSizeMap map[byte]int = map[byte]int{
	ast.INT: 8,
}

func (g *Generator) genLetInit(n ast.Node) Register {
	node := n.(*ast.LetInit)
	var_type := node.TypeGenerated()
	var_size, _ := typeSizeMap[var_type]
	result_reg := g.appendCodeFor(node.Expr)

	location_reg := g.GetRegister()
	ins := &pcode.Instruction{pcode.MOV_REG_REG, []int{location_reg.RegisterNumber, pcode.STACK_FRAME_POINTER_REG}}
	g.Program.WriteInstruction(ins)

	ins = &pcode.Instruction{pcode.SUB_REG_INT_INT,
		[]int{location_reg.RegisterNumber, node.LocationInfo.LocationOnStack * var_size}}
	g.Program.WriteInstruction(ins)

	ins = &pcode.Instruction{pcode.MOV_REG_ADDRESS_REG, []int{location_reg.RegisterNumber, result_reg.RegisterNumber}}
	g.Program.WriteInstruction(ins)

	g.ReleaseRegister(location_reg)

	g.SymTab[node.Ident] = node.LocationInfo

	return result_reg
}
