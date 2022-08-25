/* This interface will allow us to pass around all the ast nodes
 * in the evaluator as though they are all the same type
 */
package ast

import (
	"turtlego/src/tokens"
)

type Node interface {
	PrintAll(string)
	Stringify(string) string
	GetTok() tokens.Token
	TypeGenerated() TypeInfo
	NodeType() byte
}

const (
	BLOCK_NT byte = iota
	LETINIT_NT
	INFIX_NT
	PREFIX_NT

	IFEL_NT
	NOP_NT

	BOOLEAN_NT
	IDENT_NT
	STRING_NT
	FLOAT_NT
	INT_NT

	BUILTIN_NT
	CHR_NT

	LIST_NT
)
