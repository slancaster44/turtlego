/* This file maps various symbols and keywords to the relevent "tokent aliases"
 * The lexer uses these maps to determine which token aliasess go with what
 * symbols
 */
package lexer

import (
	"turtlego/src/tokens"
)

var singleChar = map[string]byte{
	string(0): tokens.EOF,
	"\n":      tokens.EOL,
	";":       tokens.EOL,
	"!":       tokens.POP,
	"=":       tokens.EQ,
	"{":       tokens.LCURL,
	"}":       tokens.RCURL,
	"(":       tokens.LPAREN,
	")":       tokens.RPAREN,
	"+":       tokens.OP1,
	"-":       tokens.OP1,
	"*":       tokens.OP2,
	"/":       tokens.OP2,
	"<":       tokens.OP4,
	">":       tokens.OP4,
	",":       tokens.COMMA,
}

var doubleChar = map[string]byte{
	"==": tokens.OP4,
	"!=": tokens.OP4,
	">=": tokens.OP4,
	"<=": tokens.OP4,
	"||": tokens.OP5,
	"&&": tokens.OP5,
}

var keywords = map[string]byte{
	"let":   tokens.LET,
	"fn":    tokens.FN,
	"str":   tokens.STR_T,
	"int":   tokens.INT_T,
	"flt":   tokens.FLT_T,
	"true":  tokens.BOOL,
	"false": tokens.BOOL,
	"print": tokens.BUILTIN,
	"if":    tokens.IF,
	"else":  tokens.ELSE,
}
