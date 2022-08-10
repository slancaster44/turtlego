package pcodegenerator

import "turtlego/src/ast"

func (g *Generator) genBlockCode(n ast.Node) Register {
	block := n.(*ast.Block)

	var reg Register
	g.pushStackFrame(block.NumStackVars)

	for iter, stmt := range block.Exprs {
		reg = g.appendCodeFor(stmt)
		if iter != len(block.Exprs)-1 {
			g.ReleaseRegister(reg)
		}
	}

	g.popStackFrame(block.NumStackVars)

	return reg
}
