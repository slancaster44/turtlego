package parser

import (
	"turtlego/src/ast"
	"turtlego/src/tokens"
)

func (p *Parser) parseBlock() ast.Node {
	tok := p.lxr.CurTok
	p.lxr.MoveUp()

	//The 0th type table is assumed to be the current one. So we must push
	//The new type table for this scope on to the stack
	p.typetables = append([]TypeTable{NewTypeTable()}, p.typetables...)

	var Exprs []ast.Node
	for p.lxr.CurTok.Alias != tokens.RCURL {
		Exprs = append(Exprs, p.parseExpr(0))
		p.skipWhitespace()
	}
	p.lxr.MoveUp()

	typeGenerated := Exprs[len(Exprs)-1].TypeGenerated()
	numOfStackVars := len(p.typetables[0].Entries)
	p.typetables = p.typetables[1:]

	return &ast.Block{Exprs, typeGenerated, numOfStackVars, tok}
}
