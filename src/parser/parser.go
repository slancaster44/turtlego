/*  This file contains the parser object which is the
 * object that actually does the dirty work of parsing
 * tokens into an ast
 */
package parser

import (
	"turtlego/src/ast"
	"turtlego/src/lexer"
	"turtlego/src/message"
	"turtlego/src/tokens"
)

/* This map associates various operators with the
 * proper precedence. The lower the number, the lower the
 * operator will be placed in the tree
 */
var precedence = map[byte]int{
	tokens.OP1:    3,
	tokens.OP2:    4,
	tokens.OP3:    5,
	tokens.OP4:    2,
	tokens.LPAREN: 7,
	//tokens.LBRACK: 7,
	//tokens.DOT:    8,
	tokens.OP5: 1,
}

type (
	prefixFn func() ast.Node
	infixFn  func(ast.Node) ast.Node
)

type operatorInputOutputKey struct {
	op        string
	inputType byte
}
type Parser struct {
	lxr  *lexer.Lexer
	Tree ast.Block

	prefixParseFns map[byte]prefixFn
	infixParseFns  map[byte]infixFn

	infixOperatorTypeMap   map[string][]byte
	infixOperatorOutputMap map[operatorInputOutputKey]byte

	prefixOperatorTypeMap   map[string][]byte
	prefixOperatorOutputMap map[operatorInputOutputKey]byte

	symtabs []SymTab
}

func New(lexer *lexer.Lexer) *Parser {
	p := &Parser{}
	p.lxr = lexer

	/* If you can't tell, I love a good map
	 */
	p.prefixParseFns = map[byte]prefixFn{
		//tokens.STR:         p.parseStr,
		tokens.FLT:     p.parseFlt,
		tokens.INT:     p.parseInt,
		tokens.BOOL:    p.parseBool,
		tokens.IDENT:   p.parseIdent,
		tokens.POP:     p.parsePrefix,
		tokens.OP1:     p.parsePrefix, //'+' and '-'
		tokens.LPAREN:  p.parseLParen,
		tokens.LET:     p.parseLet,
		tokens.LCURL:   p.parseBlock,
		tokens.BUILTIN: p.parseBuiltin,
		tokens.IF:      p.parseIfEl,
	}

	p.infixParseFns = map[byte]infixFn{
		tokens.OP1: p.parseInfix,
		tokens.OP2: p.parseInfix,
		tokens.OP3: p.parseInfix,
		tokens.OP4: p.parseInfix,
		tokens.OP5: p.parseInfix,
	}

	//Map of operators to their accepted input types
	p.infixOperatorTypeMap = map[string][]byte{
		"==": {ast.INT, ast.STR, ast.BOOL, ast.FLT}, //Ex: "==" will accept int, str, bool, and flt inputs
		"!=": {ast.INT, ast.STR, ast.BOOL, ast.FLT},
		"<=": {ast.FLT, ast.INT},
		">=": {ast.FLT, ast.INT},
		"||": {ast.BOOL},
		"&&": {ast.BOOL},
		"+":  {ast.STR, ast.INT, ast.FLT},
		"-":  {ast.FLT, ast.INT},
		"/":  {ast.FLT, ast.INT},
		"*":  {ast.FLT, ast.INT},
		"**": {ast.FLT, ast.INT},
		">":  {ast.FLT, ast.INT},
		"<":  {ast.FLT, ast.INT},
	}

	//Maps operators and their input types with the coresponding output type
	p.infixOperatorOutputMap = map[operatorInputOutputKey]byte{
		{"==", ast.INT}:  ast.BOOL, //Ex: when "==" as integer input, it has boolean output
		{"==", ast.FLT}:  ast.BOOL,
		{"==", ast.STR}:  ast.BOOL,
		{"==", ast.BOOL}: ast.BOOL,
		{"!=", ast.INT}:  ast.BOOL,
		{"!=", ast.FLT}:  ast.BOOL,
		{"!=", ast.STR}:  ast.BOOL,
		{"!=", ast.BOOL}: ast.BOOL,
		{"<=", ast.INT}:  ast.BOOL,
		{"<=", ast.FLT}:  ast.BOOL,
		{">=", ast.INT}:  ast.BOOL,
		{">=", ast.FLT}:  ast.BOOL,
		{"<", ast.INT}:   ast.BOOL,
		{"<", ast.FLT}:   ast.BOOL,
		{">", ast.INT}:   ast.BOOL,
		{">", ast.FLT}:   ast.BOOL,
		{"||", ast.BOOL}: ast.BOOL,
		{"&&", ast.BOOL}: ast.BOOL,
		{"+", ast.FLT}:   ast.FLT,
		{"+", ast.INT}:   ast.INT,
		{"+", ast.STR}:   ast.STR,
		{"-", ast.FLT}:   ast.FLT,
		{"-", ast.INT}:   ast.INT,
		{"/", ast.FLT}:   ast.FLT,
		{"/", ast.INT}:   ast.INT,
		{"*", ast.FLT}:   ast.FLT,
		{"*", ast.INT}:   ast.INT,
		{"**", ast.FLT}:  ast.FLT,
		{"**", ast.INT}:  ast.INT,
	}

	p.prefixOperatorTypeMap = map[string][]byte{
		"-": {ast.FLT, ast.INT},
		"+": {ast.FLT, ast.INT},
		"!": {ast.BOOL},
	}

	p.prefixOperatorOutputMap = map[operatorInputOutputKey]byte{
		{"-", ast.FLT}:  ast.FLT,
		{"-", ast.INT}:  ast.INT,
		{"+", ast.FLT}:  ast.FLT,
		{"+", ast.INT}:  ast.INT,
		{"!", ast.BOOL}: ast.BOOL,
	}

	p.symtabs = []SymTab{NewSymtab()}

	return p
}

func (p *Parser) addSymtabEntry(ident string, Type byte, location int) {

	p.symtabs[0].Entries[ident] = TableEntry{location, Type}
}

func (p *Parser) searchSymtab(ident string) (TableEntry, int) {
	for i, table := range p.symtabs {
		varInfo, ok := table.Entries[ident]
		if ok {
			return varInfo, len(p.symtabs) - i
		}
	}

	return TableEntry{0, ast.NO_TYPE}, 0
}

func (p *Parser) raiseError(n, m string) {
	message.Error(p.lxr.CurTok.Filename, n, m,
		p.lxr.Src().LineNo, p.lxr.Src().ColumnNo)
}

func (p *Parser) ParseProgram() {
	tree := []ast.Node{}
	for p.lxr.CurTok.Alias != tokens.EOF {

		val := p.parseExpr(0)
		tree = append(tree, val)
		p.skipWhitespace()

	}
	p.Tree = ast.Block{tree, ast.INT, len(p.symtabs[0].Entries), len(p.symtabs), tree[0].GetTok()}
}

func (p *Parser) parseExpr(prec int) ast.Node {
	p.skipWhitespace()

	preFn := p.prefixParseFns[p.lxr.CurTok.Alias]
	if preFn == nil {
		p.raiseError("Syntax",
			"Invalid prefix: '"+p.lxr.CurTok.Value+"'")
	}
	leftExpr := preFn()

	for p.notDoneParsingExpr(prec) {

		infFn := p.infixParseFns[p.lxr.CurTok.Alias]
		if infFn == nil {
			return leftExpr
		}

		leftExpr = infFn(leftExpr)
	}

	return leftExpr
}

func (p *Parser) parseLParen() ast.Node {
	p.lxr.MoveUp()
	expr := p.parseExpr(0)

	if p.lxr.CurTok.Alias != tokens.RPAREN {
		p.raiseError("Syntax", "Expected ')' got '"+p.lxr.CurTok.Value+"'")
	}
	p.lxr.MoveUp()

	return expr
}

func (p *Parser) notDoneParsingExpr(prec int) bool {
	return p.lxr.CurTok.Alias != tokens.EOL &&
		p.lxr.CurTok.Alias != tokens.EOF &&
		prec < precedence[p.lxr.CurTok.Alias]
}

func (p *Parser) skipWhitespace() {
	for p.lxr.CurTok.Alias == tokens.EOL {
		p.lxr.MoveUp()
	}
}
