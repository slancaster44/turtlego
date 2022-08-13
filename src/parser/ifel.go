package parser

import (
	"turtlego/src/ast"
	"turtlego/src/tokens"
)

func (p *Parser) parseIfEl() ast.Node {
	tok := p.lxr.CurTok
	p.lxr.MoveUp()

	cond := p.parseExpr(0)
	if cond.TypeGenerated() != ast.BOOL {
		p.raiseError("Type", "Expected boolean expression after 'if'")
	}

	trueExpr := p.parseExpr(0)
	var falseExpr ast.Node = &ast.Nop{p.lxr.CurTok}

	p.skipWhitespace()
	if p.lxr.CurTok.Alias == tokens.ELSE {
		p.lxr.MoveUp()
		falseExpr = p.parseExpr(0)
	}

	//TODO: Check trueexpr and false expr types
	if falseExpr.NodeType() != ast.NOP_NT && falseExpr.TypeGenerated() != trueExpr.TypeGenerated() {
		p.raiseError("Type", "Both the true and false branches of an if-else statement must generate the same type")
	}

	return &ast.IfElse{cond, trueExpr, falseExpr, tok}

}
