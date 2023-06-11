/* This file holds the various token aliases as constants.
 */
package tokens

const (
	EOF byte = iota
	EOL

	STR
	INT
	FLT
	CHR
	IDENT

	COLON
	PERIOD
	LPAREN
	RPAREN
	COMMA
	PLUS
	MINUS

	LD
	ST
	CALL
	TAIL
	RET

	EBREAK

	DDW
	DW
	DHW
	DB
	CDW
	CW
	CHW
	CB

	ADD
	ADDI
	AND
	ANDI
	AUPIC
	BEQ
	BGE
	BGEU
	BLT
	BLTU
	BNE
	JAL
	JALR
	LB
	LBU
	LH
	LHU
	LUI
	LW
	OR
	ORI
	SB
	SH
	SLL
	SLLI
	SLT
	SLTIU
	SLTU
	SRA
	SRAI
	SRL
	SRLI
	SUB
	SW
	XOR
	XORI

	/* These must stay in order !!!! */
	ZERO_REG
	RA_REG
	SP_REG
	GP_REG
	TP_REG
	T0_REG
	T1_REG
	T2_REG
	S0_REG
	S1_REG
	A0_REG
	A1_REG
	A2_REG
	A3_REG
	A4_REG
	A5_REG
	A6_REG
	A7_REG
	S2_REG
	S3_REG
	S4_REG
	S5_REG
	S6_REG
	S7_REG
	S8_REG
	S9_REG
	S10_REG
	S11_REG
	T3_REG
	T4_REG
	T5_REG
	T6_REG
	/* These must stay in order !!!! */

)
