package encode

const (
	MATH_RTYPE_OPCODE uint32 = 0x33
)

func EncodeRType(opcode, rd, funct3, rs1, rs2, funct7 uint32) uint32 {

	//Mask out any bits that aren't used
	opcode = 0b1111111 & opcode
	rd = 0b11111 & rd
	funct3 = 0b111 & funct3
	rs1 = 0b11111 & rs1
	rs2 = 0b11111 & rs2
	funct7 = 0b1111111 & funct7

	//Stich value together
	var retVal uint32 = opcode | (rd << 7) | (funct3 << 12) | (rs1 << 15) | (rs2 << 20) | (funct7 << 25)
	return retVal
}

func Add(regs, imms []uint32) uint32 {
	guard(regs, 3, imms, 0)
	return EncodeRType(MATH_RTYPE_OPCODE, regs[0], 0x0, regs[1], regs[2], 0x0)
}

func Sub(regs, imms []uint32) uint32 {
	guard(regs, 3, imms, 0)
	return EncodeRType(MATH_RTYPE_OPCODE, regs[0], 0x0, regs[1], regs[2], 0x20)
}

func Sll(regs, imms []uint32) uint32 {
	guard(regs, 3, imms, 0)
	return EncodeRType(MATH_RTYPE_OPCODE, regs[0], 0x1, regs[1], regs[2], 0x0)
}

func Slt(regs, imms []uint32) uint32 {
	guard(regs, 3, imms, 0)
	return EncodeRType(MATH_RTYPE_OPCODE, regs[0], 0x2, regs[1], regs[2], 0x0)
}

func Sltu(regs, imms []uint32) uint32 {
	guard(regs, 3, imms, 0)
	return EncodeRType(MATH_RTYPE_OPCODE, regs[0], 0x3, regs[1], regs[2], 0x0)
}

func Xor(regs, imms []uint32) uint32 {
	guard(regs, 3, imms, 0)
	return EncodeRType(MATH_RTYPE_OPCODE, regs[0], 0x4, regs[1], regs[2], 0x0)
}

func Srl(regs, imms []uint32) uint32 {
	guard(regs, 3, imms, 0)
	return EncodeRType(MATH_RTYPE_OPCODE, regs[0], 0x5, regs[1], regs[2], 0x0)
}

func Sra(regs, imms []uint32) uint32 {
	guard(regs, 3, imms, 0)
	return EncodeRType(MATH_RTYPE_OPCODE, regs[0], 0x5, regs[1], regs[2], 0x20)
}

func Or(regs, imms []uint32) uint32 {
	guard(regs, 3, imms, 0)
	return EncodeRType(MATH_RTYPE_OPCODE, regs[0], 0x6, regs[1], regs[2], 0x0)
}

func And(regs, imms []uint32) uint32 {
	guard(regs, 3, imms, 0)
	return EncodeRType(MATH_RTYPE_OPCODE, regs[0], 0x7, regs[1], regs[2], 0x0)
}
