package parser

import (
	"strconv"
	"turtlego/src/ast"
)

func (p *Parser) parseFlt() ast.Node {
	v, ok := strconv.ParseFloat(p.lxr.CurTok.Value, 64)

	if ok != nil {
		p.raiseError("Syntax", "Could not parse floating point '"+p.lxr.CurTok.Value+"'")
	}

	n := &ast.Flt{v, p.lxr.CurTok}
	p.lxr.MoveUp()
	return n
}

func (p *Parser) parseInt() ast.Node {
	v, ok := strconv.Atoi(p.lxr.CurTok.Value)

	if ok != nil {
		p.raiseError("Syntax", "Could not parse integer '"+p.lxr.CurTok.Value+"'")
	}

	n := &ast.Int{v, p.lxr.CurTok}
	p.lxr.MoveUp()
	return n
}

/////////////////////////////////////////////////
// "<text>" or '<text>
func (p *Parser) parseStr() ast.Node {
	v := p.lxr.CurTok.Value
	n := &ast.String{v, p.lxr.CurTok}

	p.lxr.MoveUp()
	return n
}

/////////////////////////////////////////////////
//Syntax: true OR false
func (p *Parser) parseBool() ast.Node {
	v, err := strconv.ParseBool(p.lxr.CurTok.Value)

	if err != nil {
		p.raiseError("Syntax", "Could not parse boolean '"+p.lxr.CurTok.Value+"'")
	}

	n := &ast.Boolean{v, p.lxr.CurTok}
	p.lxr.MoveUp()
	return n
}

/////////////////////////////////////////////////
//Syntax: <ident>
func (p *Parser) parseIdent() ast.Node {
	v := p.lxr.CurTok.Value

	t, depth := p.searchTypeTable(v)
	if t == ast.NO_TYPE {
		p.raiseError("Identifier", "Could not resolve identifier '"+v+"'")
	}

	n := &ast.Identifier{v, p.lxr.CurTok, depth, t}
	p.lxr.MoveUp()

	return n
}
