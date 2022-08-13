package pcodegenerator

import (
	"turtlego/src/ast"
	"turtlego/src/pcode"
)

func (g *Generator) genIfElse(n ast.Node) Register {
	ifelnode := n.(*ast.IfElse)
	ret_reg := g.GetRegister()

	//Code for conditional
	cond_reg := g.appendCodeFor(ifelnode.Cond)

	//Jump if zero to false block
	jmz_reg_ins, _ := g.WriteInstruction(pcode.JMZ_REG, cond_reg.RegisterNumber, UNKOWN_ADDRESS)
	g.ReleaseRegister(cond_reg)

	//Code For True Block
	true_reg := g.appendCodeFor(ifelnode.TrueExpr)
	g.WriteInstruction(pcode.MOV_REG_REG, ret_reg.RegisterNumber, true_reg.RegisterNumber)
	g.ReleaseRegister(true_reg)

	//Jump to end of false block
	jmp_ins, _ := g.WriteInstruction(pcode.JMP, UNKOWN_ADDRESS)

	_, address := g.WriteInstruction(pcode.NOP)
	jmz_reg_ins.Arguments[1] = address

	//Code For False Block
	false_reg := g.appendCodeFor(ifelnode.FalseExpr)
	g.WriteInstruction(pcode.MOV_REG_REG, ret_reg.RegisterNumber, false_reg.RegisterNumber)
	g.ReleaseRegister(false_reg)

	_, address = g.WriteInstruction(pcode.NOP) //Target for those skipping over false block

	jmp_ins.Arguments[0] = address

	return ret_reg
}
