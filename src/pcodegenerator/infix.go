package pcodegenerator

import (
	"turtlego/src/ast"
)

func (p *Generator) genInfixCode(n ast.Node) Register {
	op := n.(*ast.InfixExpr).Op
	key := OpTypePair{n.(*ast.InfixExpr).Op, n.TypeGenerated()}
	fn, ok := p.infixGenFnMap[key]

	if !ok {
		p.raiseError("Generation", "Cannot generate code for operator '"+op+"'", n.GetTok())
	}
	return fn(n)
}

// This function creates another function that can handle an infix operation
// on integers. For example, it can take the ADD_REG_INT_INT and the ADD_REG_REG_INT
// instructions, and generate a function that can convert an addition ast into
// the proper instructions
func (p *Generator) mkInfixOpGenFunc(OP_REG_IMM, OP_REG_REG byte) func(n ast.Node) Register {
	fn := func(n ast.Node) Register {
		//Calculate Left Side
		left := n.(*ast.InfixExpr).Left
		reg := p.appendCodeFor(left)

		//Calculate Right Side
		right := n.(*ast.InfixExpr).Right

		//If the right side is an integer, we can
		//add it to R1 as an immediate

		switch r := right.(type) {
		case *ast.Int:
			p.WriteInstruction(OP_REG_IMM, reg.RegisterNumber, r.Value)
		case *ast.Boolean:
			var valAsInt int
			if r.Value {
				valAsInt = 1
			} else {
				valAsInt = 0
			}
			p.WriteInstruction(OP_REG_IMM, reg.RegisterNumber, valAsInt)
		default:
			reg2 := p.appendCodeFor(right)
			p.WriteInstruction(OP_REG_REG, reg.RegisterNumber, reg2.RegisterNumber)
			p.ReleaseRegister(reg2) //The result has been moved to reg1, so we can now release reg2
		}

		return reg
	}

	return fn
}
