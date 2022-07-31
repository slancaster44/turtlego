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

	//p.ReleaseRegister(reg)
	return reg
}
