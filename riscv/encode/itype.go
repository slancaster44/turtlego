package encode

const MATH_ITYPE_OPCODE = 0x13
const LOAD_ITYPE_OPCODE = 0x03
const JUMP_ITYPE_OPCODE = 0x67

func EncodeIType(opcode, rd, funct3, rs1, imm uint32) uint32 {

	//Clean out extra bits
	opcode = 0b1111111 & opcode
	rd = 0b11111 & rd
	funct3 = 0b111 & funct3
	rs1 = 0b11111 & rs1
	imm = 0x7FF & imm //11-bits

	return opcode | (rd << 7) | (funct3 << 12) | (rs1 << 15) | (imm << 20)
}

func Addi(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeIType(MATH_ITYPE_OPCODE, regs[0], 0, regs[1], imms[0])
}

func Addiw(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeIType(MATH_ITYPE_OPCODE+0x8, regs[0], 0, regs[1], imms[0])
}

func Slti(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeIType(MATH_ITYPE_OPCODE, regs[0], 0b10, regs[1], imms[0])
}

func Sltiu(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeIType(MATH_ITYPE_OPCODE, regs[0], 0b11, regs[1], imms[0])
}

func Xori(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeIType(MATH_ITYPE_OPCODE, regs[0], 0b100, regs[1], imms[0])
}

func Ori(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeIType(MATH_ITYPE_OPCODE, regs[0], 0b110, regs[1], imms[0])
}

func Andi(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeIType(MATH_ITYPE_OPCODE, regs[0], 0b111, regs[1], imms[0])
}

func Slli(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeIType(MATH_ITYPE_OPCODE, regs[0], 0b001, regs[1], imms[0])
}

func Srli(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeIType(MATH_ITYPE_OPCODE, regs[0], 0b101, regs[1], imms[0])
}

func Srai(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeIType(MATH_ITYPE_OPCODE, regs[0], 0b101, regs[1], 0x400+imms[0])
}

func Lb(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeIType(LOAD_ITYPE_OPCODE, regs[0], 0, regs[1], imms[0])
}

func Lh(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeIType(LOAD_ITYPE_OPCODE, regs[0], 0b001, regs[1], imms[0])
}

func Lw(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeIType(LOAD_ITYPE_OPCODE, regs[0], 0b010, regs[1], imms[0])
}

func Lbu(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeIType(LOAD_ITYPE_OPCODE, regs[0], 0b100, regs[1], imms[0])
}

func Lhu(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeIType(LOAD_ITYPE_OPCODE, regs[0], 0b101, regs[1], imms[0])
}

func Lwu(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeIType(LOAD_ITYPE_OPCODE, regs[0], 0b110, regs[1], imms[0])
}

func Ld(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeIType(LOAD_ITYPE_OPCODE, regs[0], 0b011, regs[1], imms[0])
}

func Jalr(regs, imms []uint32) uint32 {
	guard(regs, 2, imms, 1)
	return EncodeIType(JUMP_ITYPE_OPCODE, regs[0], 0, regs[1], imms[0])
}
