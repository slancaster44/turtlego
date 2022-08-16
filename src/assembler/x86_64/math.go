package x86_64

import (
	"turtlego/src/assembler/backpatch"
	"turtlego/src/pcode"
)

var add_imm_reg []byte = []byte{0x48, 0x81}
var add_imm_rax []byte = []byte{0x48, 0x05}

func AddImmReg(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data := []byte{}, []byte{}

	if ins.Arguments[0] == pcode.REG1 {
		code = append(code, add_imm_rax...)
	} else {
		code = append(code, add_imm_reg...)
		code = append(code, singleRegisterEncoding(ins.Arguments[0]))
	}

	code = append(code, mkIntByteArray(ins.Arguments[1])...)

	return code, data, []backpatch.BackPatch{}
}

var add_reg_reg []byte = []byte{0x48, 0x01}

func AddRegReg(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data := []byte{}, []byte{}

	code = append(code, add_reg_reg...)
	code = append(code, dualRegisterEncoding(ins.Arguments[0], ins.Arguments[1]))
	return code, data, []backpatch.BackPatch{}
}

var sub_imm_reg []byte = []byte{0x48, 0x81}
var sub_imm_rax []byte = []byte{0x48, 0x2D}
var sub_reg_adjustment byte = 0xE8

func SubImmReg(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data := []byte{}, []byte{}

	if ins.Arguments[0] == pcode.REG1 {
		code = append(code, sub_imm_rax...)
	} else {
		code = append(code, sub_imm_reg...)
		code = append(code, sub_reg_adjustment|singleRegisterEncoding(ins.Arguments[0]))
	}

	code = append(code, mkIntByteArray(ins.Arguments[1])...)

	return code, data, []backpatch.BackPatch{}
}

var sub_reg_reg []byte = []byte{0x48, 0x29}

func SubRegReg(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data := []byte{}, []byte{}

	code = append(code, sub_reg_reg...)
	code = append(code, dualRegisterEncoding(ins.Arguments[0], ins.Arguments[1]))

	return code, data, []backpatch.BackPatch{}
}

var mul_imm_reg []byte = []byte{0x48, 0x69}

func MulImmReg(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

	code = append(code, mul_imm_reg...)
	code = append(code, dualRegisterEncoding(ins.Arguments[0], ins.Arguments[0]))
	code = append(code, mkIntByteArray(ins.Arguments[1])...)

	return code, data, patches
}

var mul_reg_reg []byte = []byte{0x48, 0x0F, 0xAF}

func MulRegReg(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

	code = append(code, mul_reg_reg...)
	code = append(code, dualRegisterEncoding(ins.Arguments[1], ins.Arguments[0]))

	return code, data, patches
}
