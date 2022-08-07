package pcodegenerator

import "turtlego/src/pcode"

func (g *Generator) genPushRegToStack(r int) {
	i := pcode.Instruction{pcode.PUSH_REG, []int{r}}
	g.Program.WriteInstruction(&i)
}

func (g *Generator) genPopToReg(r int) {
	i := pcode.Instruction{pcode.POP, []int{r}}
	g.Program.WriteInstruction(&i)
}

func (g *Generator) pushStackFrame(numVars int) {

	reg := g.GetRegister()
	ins := &pcode.Instruction{pcode.LOADINT, []int{reg.RegisterNumber, numVars * 8}} //TODO: Generalize
	g.Program.WriteInstruction(ins)

	ins = &pcode.Instruction{pcode.ADD_REG_REG_INT, []int{pcode.STACK_POINTER, reg.RegisterNumber}}
	g.Program.WriteInstruction(ins)

	ins = &pcode.Instruction{pcode.MOV_REG_REG, []int{pcode.STACK_FRAME_POINTER_REG, pcode.STACK_POINTER}}
	g.Program.WriteInstruction(ins)

	g.ReleaseRegister(reg)
}
