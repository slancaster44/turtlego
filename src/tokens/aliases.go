/* This file holds the various token aliases as constants.
 */
package tokens

const (
	EOF byte = iota
	EOL

	IDENT
	FLT
	INT
	STR
	CHR
	BOOL

	LCURL
	RCURL
	LPAREN
	RPAREN
	LBRACK
	RBRACK

	IF
	ELSE

	POP

	OP1
	OP2
	OP3
	OP4
	OP5

	EQ

	COMMA

	LET
	FN

	FLT_T
	INT_T
	STR_T

	BUILTIN
)
