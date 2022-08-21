package pcodegenerator

import (
	"turtlego/src/ast"
	"turtlego/src/pcode"
)

func (g *Generator) genBuiltinCode(n ast.Node) Register {
	bn := n.(*ast.Builtin)

	fn, ok := g.builtinsGenMap[bn.Name]
	if !ok {
		g.raiseError("Generator", "No generation fn for builtin", n.GetTok())
	}

	return fn(bn)
}

func (g *Generator) genPrintCode(n *ast.Builtin) Register {

	for _, arg := range n.Args {
		if arg.TypeGenerated() != ast.CHR && arg.TypeGenerated() != ast.STR {
			g.raiseError("Type", "print() may only take characters or strings as inputs", arg.GetTok())
		}

		reg := g.appendCodeFor(arg)
		g.WriteInstruction(pcode.BUILTIN_CALL, pcode.BUILTIN_PRINT, reg.RegisterNumber)
		g.ReleaseRegister(reg)
	}

	return g.genIntCode(&ast.Int{0, n.GetTok()})
}
