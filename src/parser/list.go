package parser

import (
	"turtlego/src/ast"
	"turtlego/src/tokens"
)

func (p *Parser) parseList() ast.Node {
	tok := p.lxr.CurTok
	p.lxr.MoveUp()

	vals := []ast.Node{}
	if p.lxr.CurTok.Alias != tokens.RBRACK {
		vals = p.parseArgs()
	}

	if p.lxr.CurTok.Alias != tokens.RBRACK {
		p.raiseError("Syntax", "Expected closing ']' at end of list")
	}
	p.lxr.MoveUp()

	contentsType := ast.TypeInfo(ast.NO_TYPE)
	if len(vals) != 0 {
		contentsType = vals[0].TypeGenerated()

		for _, i := range vals {
			if !i.TypeGenerated().DoesThisMatch(contentsType) {
				p.raiseError("Type", "All values in a list must have a similar type")
			}
		}
	}

	return &ast.List{ast.ListType{contentsType}, vals, tok}
}
