package pcodegenerator

import "turtlego/src/pcode"

func (g *Generator) genPushRegToStack(r int) {
	g.WriteInstruction(pcode.PUSH_REG, r)
}

func (g *Generator) genPopToReg(r int) {
	g.WriteInstruction(pcode.POP, r)
}

func (g *Generator) pushStackFrame(numVars int, scopeDepth int) {

	reg := g.GetRegister()

	g.WriteInstruction(pcode.LOADINT, reg.RegisterNumber, numVars*STACK_VAR_SIZE)
	g.WriteInstruction(pcode.SUB_REG_REG_INT, pcode.STACK_POINTER, reg.RegisterNumber)

	//Write Stack Frame metadata
	g.WriteInstruction(pcode.PUSH_INT, scopeDepth)                    //Write scope of this stack frame
	g.WriteInstruction(pcode.PUSH_REG, pcode.STACK_FRAME_POINTER_REG) //Write location of last stack frame

	g.WriteInstruction(pcode.MOV_REG_REG, pcode.STACK_FRAME_POINTER_REG, pcode.STACK_POINTER)

	g.ReleaseRegister(reg)
}

func (g *Generator) popStackFrame(numVars int) {
	reg := g.GetRegister()

	g.WriteInstruction(pcode.LOADINT, reg.RegisterNumber, numVars*STACK_VAR_SIZE+SIZE_OF_STACK_METADATA)
	g.WriteInstruction(pcode.ADD_REG_REG_INT, pcode.STACK_POINTER, reg.RegisterNumber)
	g.WriteInstruction(pcode.MOV_REG_REG, pcode.STACK_FRAME_POINTER_REG, pcode.STACK_POINTER)

	g.ReleaseRegister(reg)
}
