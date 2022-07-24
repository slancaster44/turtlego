package pcodegnerator

import (
	"turtlego/src/ast"
	"turtlego/src/message"
	"turtlego/src/pcode"
	"turtlego/src/tokens"
)

type Generator struct {
	SyntaxTree   []ast.Node
	CodeGenFnMap map[byte]func(ast.Node)
	Output       pcode.Program
}

func NewGenerator(st []ast.Node) *Generator {
	ret_val := &Generator{}
	ret_val.SyntaxTree = st

	ret_val.CodeGenFnMap = map[byte]func(ast.Node){
		ast.INT_NT: ret_val.genIntCode,
	}

	return ret_val
}

func (g *Generator) raiseError(n, m string, tok tokens.Token) {
	message.Error(tok.Filename, n, m, tok.LineNo, tok.ColumnNo)
}

func (g *Generator) GenPCode() {
	for _, stmt := range g.SyntaxTree {
		g.appendCodeFor(stmt)
	}
}

func (g *Generator) appendCodeFor(stmt ast.Node) {
	fn, ok := g.CodeGenFnMap[stmt.NodeType()]

	if !ok {
		g.raiseError("Generation", "Unimplemented node type", stmt.GetTok())
	}

	fn(stmt)
}
