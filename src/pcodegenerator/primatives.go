package pcodegenerator

import (
	"turtlego/src/ast"
	"turtlego/src/pcode"
)

func (p *Generator) genIntCode(stmt ast.Node) {
	number := stmt.(*ast.Int).Value
	ins := &pcode.Instruction{
		pcode.LOADINT,
		[]int{
			int(pcode.REG1),
			number,
		},
	}

	p.Output.WriteInstruction(ins)
}
