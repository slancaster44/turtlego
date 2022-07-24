package parser

import (
	"turtlego/src/ast"
)

//Syntax: <operator><expression>
func (p *Parser) parsePrefix() ast.Node {
	tok := p.lxr.CurTok
	op := p.lxr.CurTok.Value
	p.lxr.MoveUp()

	/* The precedence on prefix operators must fall in between
	 * LPAREN (object and function calls), and other types of binary
	 * operators (such as '+' or '/') so that prefix operators are evaluated
	 * before the other math, but after function calls
	 */
	expr := p.parseExpr(6)

	p.checkOpInputType(expr.TypeGenerated(), p.prefixOperatorTypeMap[op])
	generatedType := p.prefixOperatorOutputMap[operatorInputOutputKey{op, expr.TypeGenerated()}]

	return &ast.PrefixExpr{expr, op, tok, generatedType}
}

/////////////////////////////////////////////////
//Syntax: <expr> <operator> <expr>
func (p *Parser) parseInfix(left ast.Node) ast.Node {
	op := p.lxr.CurTok.Value
	tok := p.lxr.CurTok

	prec := precedence[p.lxr.CurTok.Alias]
	p.lxr.MoveUp()
	right := p.parseExpr(prec)

	if left.TypeGenerated() != right.TypeGenerated() {
		p.raiseError("Type", "Mismatched types on right and left")
	}
	p.checkOpInputType(right.TypeGenerated(), p.infixOperatorTypeMap[op])
	outputType := p.infixOperatorOutputMap[operatorInputOutputKey{op, right.TypeGenerated()}]

	return &ast.InfixExpr{left, right, op, tok, outputType}
}

/////////////////////////////////////////////////

func (p *Parser) checkOpInputType(inputType byte, acceptedTypes []byte) {
	for _, t := range acceptedTypes {
		if t == inputType {
			return
		}
	}
	p.raiseError("Type", "Invalid input type for operator")
}
