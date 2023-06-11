package encode

func Ecall(regs, imms []uint32) uint32 {
	guard(regs, 0, imms, 0)
	return 0b1110011
}

func Ebreak(regs, imms []uint32) uint32 {
	guard(regs, 0, imms, 0)
	return 0x00100073
}
