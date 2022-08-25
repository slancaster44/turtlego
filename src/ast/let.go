package ast

import (
	"fmt"
	"strconv"
	"turtlego/src/tokens"
)

type Location struct {
	IsStack         bool
	LocationOnStack int
	ScopeDepth      int
}

// Let when the variable is intialized
type LetInit struct {
	Tok          tokens.Token
	Ident        string
	LocationInfo Location
	Expr         Node
	Type         TypeInfo
}

func (l *LetInit) GetTok() tokens.Token {
	return l.Tok
}

func (l *LetInit) Stringify(tab string) string {
	ret_val := tab + "<let var '" + l.Ident + "', depth='" + strconv.Itoa(l.LocationInfo.ScopeDepth)
	ret_val += "' stack_location='" + strconv.Itoa(l.LocationInfo.LocationOnStack) + "'>\n"
	ret_val += l.Expr.Stringify(tab + "\t")
	ret_val += tab + "<\\let>\n"
	return ret_val
}

func (l *LetInit) PrintAll(tab string) {
	fmt.Print(l.Stringify(tab))
}

func (l *LetInit) TypeGenerated() TypeInfo {
	return l.Type
}

func (l *LetInit) NodeType() byte {
	return LETINIT_NT
}

// Let when the variable value is not intialized
// Syntax: let <ident> : <type>
type LetType struct {
	Tok     tokens.Token
	Ident   string
	VarType Node
}
