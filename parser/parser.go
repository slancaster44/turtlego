package parser

import (
	"turtlego/ast"
	"turtlego/lexer"
	"turtlego/message"
	"turtlego/tokens"
)

type (
	prefixFn func() ast.Node
	infixFn  func(ast.Node) ast.Node
)

type Parser struct {
	lxr *lexer.Lexer

	infixParseFns  map[byte]infixFn
	prefixParseFns map[byte]prefixFn
	precedenceMap  map[byte]int

	skipables []byte
	nodeCount func() int
}

func New(l *lexer.Lexer) Parser {
	p := Parser{lxr: l}
	p.nodeCount = ast.GetStaticCounter()
	return p
}

func (p *Parser) SetMaps(ifns map[byte]infixFn, pfns map[byte]prefixFn, prec map[byte]int) {
	p.infixParseFns = ifns
	p.prefixParseFns = pfns
	p.precedenceMap = prec
}

func (p *Parser) SetSkipables(s []byte) {
	p.skipables = s
}

func (p *Parser) raiseError(n, m string) {
	message.Error(p.lxr.CurTok.Filename, n, m,
		p.lxr.Src().LineNo, p.lxr.Src().ColumnNo)
}

func (p *Parser) Skip() {
	shouldSkip := func() bool {
		for _, i := range p.skipables {
			if i == p.lxr.CurTok.Alias {
				return true
			}
		}
		return false
	}

	for shouldSkip() {
		p.lxr.MoveUp()
	}
}

func (p *Parser) ParsePogram() ast.Program {
	var retVal ast.Program
	for !p.lxr.IsDone() {
		retVal = append(retVal, p.ParseExpression(0))
		p.Skip()
	}

	return retVal
}

func (p *Parser) ParseExpression(precedence int) ast.Node {
	p.Skip()

	prefixFunction := p.prefixParseFns[p.lxr.CurTok.Alias]
	if prefixFunction == nil {
		p.raiseError("Syntax", "Unexpected prefix '"+p.lxr.CurTok.Value+"'")
	}

	expression_ast := prefixFunction()
	expression_ast.SetId(p.nodeCount())

	for p.noPrecedenceOverride(precedence) {
		infixFunction := p.infixParseFns[p.lxr.CurTok.Alias]
		if infixFunction == nil {
			return expression_ast
		}

		expression_ast = infixFunction(expression_ast)
		expression_ast.SetId(p.nodeCount())
	}

	return expression_ast
}

func (p *Parser) noPrecedenceOverride(prec int) bool {
	return p.lxr.CurTok.Alias != tokens.EOL &&
		p.lxr.CurTok.Alias != tokens.EOF &&
		prec < p.precedenceMap[p.lxr.CurTok.Alias]
}

func (p *Parser) expectToken(expected string, tokTypes ...byte) {
	has_tt := false
	for _, v := range tokTypes {
		if v == p.lxr.CurTok.Alias {
			has_tt = true
			break
		}
	}

	if !has_tt {
		p.raiseError("Syntax", "Expected "+expected+" got '"+p.lxr.CurTok.Value+"'")
	}
	p.lxr.MoveUp()
}

func (p *Parser) expectNode(expected string, nodeTypes ...int) ast.Node {
	val := p.ParseExpression(0)
	has_nt := false
	for _, v := range nodeTypes {
		if v == val.Type() {
			has_nt = true
			break
		}
	}
	if !has_nt {
		p.raiseError("Syntax", "Expected "+expected+" got "+val.Token().Value)
	}
	return val
}
