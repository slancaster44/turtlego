package encode

const STORE_OPCODE uint32 = 0x23

func EncodeSType(opcode, funct3, rs1, rs2, imm uint32) uint32 {

	//Clean out extra bits
	opcode = 0b1111111 & opcode
	funct3 = 0b111 & funct3
	rs1 = 0b11111 & rs1
	imm = 0x7FF & imm //11-bits
	imm_40 := 0b11111 & imm
	imm_115 := imm >> 5

	return opcode | (imm_40 << 7) | (funct3 << 7) | (rs1 << 15) | (rs2 << 20) | (imm_115 << 25)
}

func Sb(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeSType(STORE_OPCODE, 0, regs[0], regs[1], imms[0])
}

func Sh(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeSType(STORE_OPCODE, 0b001, regs[0], regs[1], imms[0])
}

func Sw(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeSType(STORE_OPCODE, 0b010, regs[0], regs[1], imms[0])
}
