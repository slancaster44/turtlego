package pcodegenerator

import (
	"turtlego/src/ast"
	"turtlego/src/pcode"
)

func (p *Generator) genIntCode(stmt ast.Node) Register {
	number := stmt.(*ast.Int).Value
	reg := p.GetRegister()

	ins := &pcode.Instruction{pcode.LOADINT, []int{reg.RegisterNumber, number}}
	p.Program.WriteInstruction(ins)

	return reg
}

func (g *Generator) genIdentCode(stmt ast.Node) Register {
	ident := stmt.(*ast.Identifier).Value
	location := g.SymTab[ident]

	location_reg := g.GetRegister()

	ins := &pcode.Instruction{pcode.MOV_REG_REG, []int{location_reg.RegisterNumber, pcode.STACK_FRAME_POINTER_REG}}
	g.Program.WriteInstruction(ins)

	ins = &pcode.Instruction{pcode.ADD_REG_INT_INT,
		[]int{location_reg.RegisterNumber, location.LocationOnStack * STACK_VAR_SIZE}}
	g.Program.WriteInstruction(ins)

	reg := g.GetRegister()

	ins = &pcode.Instruction{pcode.MOV_REG_REG_ADDRESS, []int{reg.RegisterNumber, location_reg.RegisterNumber}}
	g.Program.WriteInstruction(ins)

	g.ReleaseRegister(location_reg)

	return reg
}
