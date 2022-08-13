package parser

import (
	"turtlego/src/ast"
	"turtlego/src/tokens"
)

func (p *Parser) parseLet() ast.Node {
	tok := p.lxr.CurTok
	p.lxr.MoveUp()

	if p.lxr.CurTok.Alias != tokens.IDENT {
		p.raiseError("Syntax", "Expected identifier after 'let', got '"+p.lxr.CurTok.Value+"'")
	}
	id := p.lxr.CurTok.Value
	p.lxr.MoveUp()

	var ret_val ast.Node
	if p.lxr.CurTok.Alias != tokens.EQ {
		p.raiseError("Syntax", "Expected '=' or ':', got '"+p.lxr.CurTok.Value+"'")
	}

	p.lxr.MoveUp()
	expr := p.parseExpr(0)

	var locationOnStack int

	//If the variable already exists in this scope then we need it's location to
	//be the already allocated space
	_, ok := p.symtabs[0].Entries[id]
	if !ok {
		locationOnStack = len(p.symtabs[0].Entries)
	} else {
		locationOnStack = p.symtabs[0].Entries[id].LocationOnStack
	}

	p.addSymtabEntry(id, expr.TypeGenerated(), locationOnStack)
	scopeDepth := len(p.symtabs)

	l := ast.Location{true, locationOnStack, scopeDepth}

	ret_val = &ast.LetInit{tok, id, l, expr, expr.TypeGenerated()}

	return ret_val
}
