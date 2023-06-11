package encode

func EncodeJType(opcode, rd, imm uint32) uint32 {

	opcode = 0b1111111 & opcode
	rd = 0b11111 & rd
	imm = imm & 0x1FFFFE
	imm_20 := imm >> 20
	imm_10_1 := (imm >> 1) & 0x3FF
	imm_11 := (imm >> 11) & 0b1
	imm_19_12 := (imm >> 12) & 0x7F

	return opcode | (rd << 7) | (imm_19_12 << 12) | (imm_11 << 20) | (imm_10_1 << 21) | (imm_20 << 31)
}

func Jal(regs, imms []uint32) uint32 {
	guard(regs, 1, imms, 1)
	return EncodeJType(0b1101111, regs[0], imms[0])
}
