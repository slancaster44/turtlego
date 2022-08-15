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

// Map retrive and print stack depth id
const OFFSET_FOR_SCOPE_ID int = 0x08

func (g *Generator) genIdentCode(stmt ast.Node) Register {
	ident := stmt.(*ast.Identifier).Value
	location := g.SymTab[ident]

	location_reg := g.GetRegister()
	g.WriteInstruction(pcode.MOV_REG_REG, location_reg.RegisterNumber, pcode.STACK_FRAME_POINTER_REG)

	cur_scope_reg := g.GetRegister()
	g.WriteInstruction(pcode.MOV_REG_REG, cur_scope_reg.RegisterNumber, location_reg.RegisterNumber)
	g.WriteInstruction(pcode.ADD_REG_INT_INT, cur_scope_reg.RegisterNumber, OFFSET_FOR_SCOPE_ID)
	g.WriteInstruction(pcode.MOV_REG_REG_ADDRESS, cur_scope_reg.RegisterNumber, cur_scope_reg.RegisterNumber)
	g.WriteInstruction(pcode.CMP_REG_INT, cur_scope_reg.RegisterNumber, location.ScopeDepth)
	jnz, _ := g.WriteInstruction(pcode.JNZ, 0x0)

	_, start_of_loop :=
		g.WriteInstruction(pcode.MOV_REG_REG, cur_scope_reg.RegisterNumber, location_reg.RegisterNumber)
	g.WriteInstruction(pcode.ADD_REG_INT_INT, cur_scope_reg.RegisterNumber, 0x00)
	g.WriteInstruction(pcode.MOV_REG_REG_ADDRESS, location_reg.RegisterNumber, cur_scope_reg.RegisterNumber)
	g.WriteInstruction(pcode.MOV_REG_REG_ADDRESS, cur_scope_reg.RegisterNumber, cur_scope_reg.RegisterNumber)

	g.WriteInstruction(pcode.ADD_REG_INT_INT, cur_scope_reg.RegisterNumber, OFFSET_FOR_SCOPE_ID)
	g.WriteInstruction(pcode.MOV_REG_REG_ADDRESS, cur_scope_reg.RegisterNumber, cur_scope_reg.RegisterNumber)
	g.WriteInstruction(pcode.CMP_REG_INT, cur_scope_reg.RegisterNumber, location.ScopeDepth)
	jnz2, _ := g.WriteInstruction(pcode.JNZ, 0x0)
	g.WriteInstruction(pcode.JMP, start_of_loop)

	_, end_of_loop :=
		g.WriteInstruction(pcode.ADD_REG_INT_INT, location_reg.RegisterNumber, (location.LocationOnStack*STACK_VAR_SIZE)+SIZE_OF_STACK_METADATA)

	g.ReleaseRegister(cur_scope_reg)

	reg := g.GetRegister()
	g.WriteInstruction(pcode.MOV_REG_REG_ADDRESS, reg.RegisterNumber, location_reg.RegisterNumber)

	g.ReleaseRegister(location_reg)

	jnz.Arguments[0] = end_of_loop
	jnz2.Arguments[0] = end_of_loop

	return reg
}
