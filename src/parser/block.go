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
	p.symtabs = append([]SymTab{NewSymtab()}, p.symtabs...)

	var Exprs []ast.Node
	for p.lxr.CurTok.Alias != tokens.RCURL {
		Exprs = append(Exprs, p.parseExpr(0))
		p.skipWhitespace()
	}
	p.lxr.MoveUp()
	if len(Exprs) == 0 {
		p.raiseError("Expression", "Cannot have a block expression that returns no value")
	}
	scopeDepth := len(p.symtabs)

	typeGenerated := Exprs[len(Exprs)-1].TypeGenerated()
	numOfStackVars := len(p.symtabs[0].Entries)
	p.symtabs = p.symtabs[1:]

	return &ast.Block{Exprs, typeGenerated, numOfStackVars, scopeDepth, tok}
}
