package encode

func EncodeUType(opcode, rd, imm uint32) uint32 {

	opcode = 0b1111111 & opcode
	rd = 0b11111 & rd

	return opcode | (rd << 7) | (imm << 12)
}

func Lui(regs, imms []uint32) uint32 {
	guard(regs, 1, imms, 1)
	return EncodeUType(0b0110111, regs[0], imms[0])
}

func Auipc(regs, imms []uint32) uint32 {
	guard(regs, 1, imms, 1)
	return EncodeUType(0b0010111, regs[0], imms[0])
}
