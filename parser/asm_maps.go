package parser

import (
	"turtlego/tokens"
)

func (p *Parser) SetAsmMaps() {
	p.prefixParseFns = map[byte]prefixFn{
		tokens.ZERO_REG: p.parseRegister,
		tokens.RA_REG:   p.parseRegister,
		tokens.SP_REG:   p.parseRegister,
		tokens.GP_REG:   p.parseRegister,
		tokens.TP_REG:   p.parseRegister,
		tokens.T0_REG:   p.parseRegister,
		tokens.T1_REG:   p.parseRegister,
		tokens.T2_REG:   p.parseRegister,
		tokens.T3_REG:   p.parseRegister,
		tokens.T4_REG:   p.parseRegister,
		tokens.T5_REG:   p.parseRegister,
		tokens.T6_REG:   p.parseRegister,
		tokens.S0_REG:   p.parseRegister,
		tokens.S1_REG:   p.parseRegister,
		tokens.S2_REG:   p.parseRegister,
		tokens.S3_REG:   p.parseRegister,
		tokens.S4_REG:   p.parseRegister,
		tokens.S5_REG:   p.parseRegister,
		tokens.S6_REG:   p.parseRegister,
		tokens.S7_REG:   p.parseRegister,
		tokens.S8_REG:   p.parseRegister,
		tokens.S9_REG:   p.parseRegister,
		tokens.S10_REG:  p.parseRegister,
		tokens.S11_REG:  p.parseRegister,
		tokens.A0_REG:   p.parseRegister,
		tokens.A2_REG:   p.parseRegister,
		tokens.A3_REG:   p.parseRegister,
		tokens.A4_REG:   p.parseRegister,
		tokens.A5_REG:   p.parseRegister,
		tokens.A6_REG:   p.parseRegister,
		tokens.A7_REG:   p.parseRegister,
		tokens.IDENT:    p.parseLabel,
		tokens.ADD:      p.parseRInstuction,
		tokens.JAL:      p.parseUJInstruction,
		tokens.BEQ:      p.parseBInstruction,
		tokens.SRAI:     p.parseIWithImm,
		tokens.JALR:     p.parseIWithOffset,
		tokens.INT:      p.parseInteger,
		tokens.SB:       p.parseSType,
		tokens.ADDI:     p.parseIWithImm,
		tokens.PERIOD:   p.parseDotLabel,
		tokens.AND:      p.parseRInstuction,
		tokens.XOR:      p.parseRInstuction,
		tokens.EBREAK:   p.parseSingleInstruction,
	}
}
