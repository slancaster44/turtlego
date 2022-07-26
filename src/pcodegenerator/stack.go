package pcodegenerator

import "turtlego/src/pcode"

func (g *Generator) genPushRegToStack(r int) {
	i := pcode.Instruction{pcode.PUSH_REG, []int{r}}
	g.Output.WriteInstruction(&i)
}

func (g *Generator) genPopToReg(r int) {
	i := pcode.Instruction{pcode.POP, []int{r}}
	g.Output.WriteInstruction(&i)
}
