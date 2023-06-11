package encode

const BRANCH_OPCODE uint32 = 0b1100011

func EncodeBType(opcode, funct3, rs1, rs2, imm uint32) uint32 {

	opcode = 0b1111111 & opcode
	funct3 = 0b111 & funct3
	rs1 = 0b11111 & rs1
	rs2 = 0b11111 & rs2
	imm = imm & 0xFFF
	imm_12 := imm >> 12
	imm_10_5 := (imm >> 5) & 0x1F
	imm_11 := (imm >> 11) & 1
	imm_4_1 := (imm >> 1) & 0xF

	return opcode | (imm_11 << 7) | (imm_4_1 << 8) | (funct3 << 12) | (rs1 << 15) | (rs2 << 20) | (imm_10_5 << 25) | (imm_12 << 31)
}

func Beq(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeBType(BRANCH_OPCODE, 0, regs[0], regs[1], imms[0])
}

func Bne(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeBType(BRANCH_OPCODE, 0b001, regs[0], regs[1], imms[0])
}

func Blt(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeBType(BRANCH_OPCODE, 0b100, regs[0], regs[1], imms[0])
}

func Bge(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeBType(BRANCH_OPCODE, 0b101, regs[0], regs[1], imms[0])
}

func Bltu(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeBType(BRANCH_OPCODE, 0b110, regs[0], regs[1], imms[0])
}

func Bgeu(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeBType(BRANCH_OPCODE, 0b111, regs[0], regs[1], imms[0])
}
