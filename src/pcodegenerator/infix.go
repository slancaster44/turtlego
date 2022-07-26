package pcodegenerator

import (
	"turtlego/src/ast"
	"turtlego/src/pcode"
)

func (p *Generator) genInfixCode(n ast.Node) {
	op := n.(*ast.InfixExpr).Op
	key := OpTypePair{n.(*ast.InfixExpr).Op, n.TypeGenerated()}
	fn, ok := p.infixGenFnMap[key]

	if !ok {
		p.raiseError("Generation", "Cannot generate code for operator '"+op+"'", n.GetTok())
	}
	fn(n)
}

func (p *Generator) genAddInt(n ast.Node) {
	//Calculate Left Side
	p.appendCodeFor(n.(*ast.InfixExpr).Left)
	//Push Left Side to Stack
	p.genPushRegToStack(int(pcode.REG1))
	//Calculate Right Side
	p.appendCodeFor(n.(*ast.InfixExpr).Right)
	//Pop Left Side to R2
	p.genPopToReg(int(pcode.REG2))
	//Add R1, and R2, leave results in R1
	i := pcode.Instruction{pcode.ADD_REG_REG_INT, []int{int(pcode.REG1), int(pcode.REG2)}}
	p.Output.WriteInstruction(&i)
}
