package pcodegenerator

import (
	"turtlego/src/ast"
	"turtlego/src/pcode"
)

func (p *Generator) genIntCode(stmt ast.Node) Register {
	number := stmt.(*ast.Int).Value
	reg := p.GetRegister()

	p.WriteInstruction(pcode.LOADINT, reg.RegisterNumber, number)

	return reg
}

func (p *Generator) genBoolCode(stmt ast.Node) Register {
	bl := stmt.(*ast.Boolean).Value
	reg := p.GetRegister()

	if bl {
		p.WriteInstruction(pcode.LOADINT, reg.RegisterNumber, 1)
	} else {
		p.WriteInstruction(pcode.LOADINT, reg.RegisterNumber, 0)
	}

	return reg
}

func (g *Generator) genIdentCode(stmt ast.Node) Register {
	ident := stmt.(*ast.Identifier).Value
	location := g.SymTab[ident]

	location_reg := g.GetRegister()

	g.WriteInstruction(pcode.MOV_REG_REG, location_reg.RegisterNumber, pcode.STACK_FRAME_POINTER_REG)
	g.WriteInstruction(pcode.ADD_REG_INT_INT, location_reg.RegisterNumber, location.LocationOnStack*STACK_VAR_SIZE)

	reg := g.GetRegister()

	g.WriteInstruction(pcode.MOV_REG_REG_ADDRESS, reg.RegisterNumber, location_reg.RegisterNumber)

	g.ReleaseRegister(location_reg)

	return reg
}
