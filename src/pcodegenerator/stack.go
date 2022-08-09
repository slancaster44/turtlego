package pcodegenerator

import "turtlego/src/pcode"

func (g *Generator) genPushRegToStack(r int) {
	g.WriteInstruction(pcode.PUSH_REG, r)
}

func (g *Generator) genPopToReg(r int) {
	g.WriteInstruction(pcode.POP, r)
}

func (g *Generator) pushStackFrame(numVars int) {

	reg := g.GetRegister()

	g.WriteInstruction(pcode.LOADINT, reg.RegisterNumber, numVars*STACK_VAR_SIZE)
	g.WriteInstruction(pcode.SUB_REG_REG_INT, pcode.STACK_POINTER, reg.RegisterNumber)
	g.WriteInstruction(pcode.MOV_REG_REG, pcode.STACK_FRAME_POINTER_REG, pcode.STACK_POINTER)

	g.ReleaseRegister(reg)
}
