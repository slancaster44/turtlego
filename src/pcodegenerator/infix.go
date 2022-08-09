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
func (p *Generator) mkInfixOpFuncInt(OP_REG_INT_INT, OP_REG_REG_INT byte) func(n ast.Node) Register {
	fn := func(n ast.Node) Register {
		//Calculate Left Side
		left := n.(*ast.InfixExpr).Left
		reg := p.appendCodeFor(left)

		//Calculate Right Side
		right := n.(*ast.InfixExpr).Right

		//If the right side is an integer, we can
		//add it to R1 as an immediate
		if right.NodeType() == ast.INT_NT {
			p.WriteInstruction(OP_REG_INT_INT, reg.RegisterNumber, right.(*ast.Int).Value)
		} else {
			reg2 := p.appendCodeFor(right)
			p.WriteInstruction(OP_REG_REG_INT, reg.RegisterNumber, reg2.RegisterNumber)
			p.ReleaseRegister(reg2) //The result has been moved to reg1, so we can now release reg2
		}

		return reg
	}

	return fn
}
