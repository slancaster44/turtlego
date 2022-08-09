package parser

import (
	"turtlego/src/ast"
	"turtlego/src/tokens"
)

//TODO: Proper print return value
//once strings are implemented

var BuiltinRetMap map[string]byte = map[string]byte{
	"print": ast.INT,
}

func (p *Parser) parseBuiltin() ast.Node {
	tok := p.lxr.CurTok
	name := p.lxr.CurTok.Value
	p.lxr.MoveUp()

	if p.lxr.CurTok.Alias != tokens.LPAREN {
		p.raiseError("Syntax", "Expected '('")
	}
	p.lxr.MoveUp()

	args := p.parseArgs()

	if p.lxr.CurTok.Alias != tokens.RPAREN {
		p.raiseError("Syntax", "Expected ')'")
	}
	p.lxr.MoveUp()

	retType, ok := BuiltinRetMap[name]
	if !ok {
		p.raiseError("Parser", "No valid return type for this builtin")
	}

	return &ast.Builtin{
		args,
		name,
		tok,
		retType,
	}
}
