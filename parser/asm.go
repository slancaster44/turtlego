package parser

import (
	"strconv"
	"turtlego/ast"
	"turtlego/tokens"
)

func (p *Parser) parseRegister() ast.Node {
	retVal := ast.Register{}
	retVal.Tok = p.lxr.CurTok
	retVal.Value = int(p.lxr.CurTok.Alias - tokens.ZERO_REG)
	p.lxr.MoveUp()

	return &retVal
}

func (p *Parser) parseLabel() ast.Node {
	retVal := ast.Label{}
	retVal.Tok = p.lxr.CurTok
	retVal.Name = p.lxr.CurTok.Value
	p.lxr.MoveUp()

	if p.lxr.CurTok.Alias == tokens.COLON {
		p.lxr.MoveUp()
	}

	return &retVal
}

func (p *Parser) parseDotLabel() ast.Node {
	retVal := ast.DotLabel{}
	retVal.Tok = p.lxr.CurTok
	p.lxr.MoveUp()

	retVal.Name = p.lxr.CurTok.Value
	p.expectToken("label", tokens.IDENT)

	if p.lxr.CurTok.Alias == tokens.COLON {
		p.lxr.MoveUp()
	}

	return &retVal
}

func (p *Parser) parseInteger() ast.Node {
	retVal := ast.Integer{}
	retVal.Value, _ = strconv.Atoi(p.lxr.CurTok.Value)
	retVal.Tok = p.lxr.CurTok
	p.lxr.MoveUp()
	return &retVal
}

func (p *Parser) parseSingleInstruction() ast.Node {
	retVal := ast.Other_Type{}
	retVal.Mnemonic = p.lxr.CurTok.Value
	retVal.Tok = p.lxr.CurTok
	p.lxr.MoveUp()
	return &retVal
}

func (p *Parser) parseRInstuction() ast.Node {
	retVal := ast.R_Type{}
	retVal.Mnemonic = p.lxr.CurTok.Value
	retVal.Tok = p.lxr.CurTok
	p.lxr.MoveUp()

	retVal.Rd = p.expectNode("register", ast.REGISTER_NT).(*ast.Register)
	p.expectToken(",", tokens.COMMA)

	retVal.Rs1 = p.expectNode("register", ast.REGISTER_NT).(*ast.Register)
	p.expectToken(",", tokens.COMMA)

	retVal.Rs2 = p.expectNode("register", ast.REGISTER_NT).(*ast.Register)

	return &retVal
}

func (p *Parser) parseUJInstruction() ast.Node {
	retVal := ast.UJ_Type{}
	retVal.Tok = p.lxr.CurTok
	retVal.Mnemonic = p.lxr.CurTok.Value
	p.lxr.MoveUp()

	retVal.Rd = p.expectNode("register", ast.REGISTER_NT).(*ast.Register)
	p.expectToken(",", tokens.COMMA)

	retVal.Imm = p.expectNode("label or integer", ast.LABEL_NT, ast.DOT_LABEL_NT, ast.INTEGER_NT)

	return &retVal
}

func (p *Parser) parseBInstruction() ast.Node {
	retVal := ast.SB_Type{}
	retVal.Tok = p.lxr.CurTok
	retVal.Mnemonic = p.lxr.CurTok.Value
	p.lxr.MoveUp()

	retVal.Rs1 = p.expectNode("register", ast.REGISTER_NT).(*ast.Register)
	p.expectToken(",", tokens.COMMA)

	retVal.Rs2 = p.expectNode("regster", ast.REGISTER_NT).(*ast.Register)
	p.expectToken(",", tokens.COMMA)

	retVal.Imm = p.expectNode("label", ast.LABEL_NT, ast.DOT_LABEL_NT, ast.INTEGER_NT)

	return &retVal
}

func (p *Parser) parseIWithOffset() ast.Node {
	retVal := ast.I_Type{}
	retVal.Tok = p.lxr.CurTok
	retVal.Mnemonic = p.lxr.CurTok.Value
	p.lxr.MoveUp()

	retVal.Rd = p.expectNode("register", ast.REGISTER_NT).(*ast.Register)
	p.expectToken(",", tokens.COMMA)

	retVal.Rs1 = p.expectNode("register", ast.REGISTER_NT).(*ast.Register)
	retVal.Operator = p.lxr.CurTok.Value
	p.expectToken("+ or -", tokens.PLUS, tokens.MINUS)

	retVal.Imm = p.expectNode("integer or label", ast.INTEGER_NT, ast.LABEL_NT, ast.DOT_LABEL_NT)

	return &retVal
}

func (p *Parser) parseIWithImm() ast.Node {
	retVal := ast.I_Type{}
	retVal.Tok = p.lxr.CurTok
	retVal.Mnemonic = p.lxr.CurTok.Value
	p.lxr.MoveUp()

	retVal.Rd = p.expectNode("register", ast.REGISTER_NT).(*ast.Register)
	p.expectToken(",", tokens.COMMA)

	retVal.Rs1 = p.expectNode("register", ast.REGISTER_NT).(*ast.Register)
	p.expectToken(",", tokens.COMMA)

	retVal.Imm = p.expectNode("integer or label", ast.INTEGER_NT, ast.LABEL_NT, ast.DOT_LABEL_NT)

	return &retVal
}

func (p *Parser) parseSType() ast.Node {
	retVal := ast.SB_Type{}
	retVal.Tok = p.lxr.CurTok
	retVal.Mnemonic = p.lxr.CurTok.Value
	p.lxr.MoveUp()

	retVal.Rs1 = p.expectNode("register", ast.REGISTER_NT).(*ast.Register)

	retVal.Operator = p.lxr.CurTok.Value
	p.expectToken("+ or -", tokens.PLUS, tokens.MINUS)

	retVal.Imm = p.expectNode("integer or label", ast.INTEGER_NT, ast.LABEL_NT, ast.DOT_LABEL_NT)
	p.expectToken(",", tokens.COMMA)

	retVal.Rs2 = p.expectNode("register", ast.REGISTER_NT).(*ast.Register)

	return &retVal
}
