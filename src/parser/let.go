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
	locationOnStack := len(p.typetables[0].Entries)

	p.addTypetableEntry(id, expr.TypeGenerated())
	scopeDepth := len(p.typetables)

	l := ast.Location{true, locationOnStack, scopeDepth}

	ret_val = &ast.LetInit{tok, id, l, expr, expr.TypeGenerated()}

	return ret_val
}
