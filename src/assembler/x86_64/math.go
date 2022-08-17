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

// x86 has no divide immediate by register, so first we will convert it to a div_reg_reg instruction
func DivImmReg(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

	imm := ins.Arguments[1]
	reg := ins.Arguments[0]

	var reg_for_imm int
	if reg == pcode.REG1 {
		reg_for_imm = pcode.REG2
	} else {
		reg_for_imm = pcode.REG1
	}

	mov, _, _ := genAuxInstruction(MovRegImm, reg_for_imm, imm)
	div, _, _ := genAuxInstruction(DivRegReg, reg, reg_for_imm)

	code = append(code, mov...)
	code = append(code, div...)

	return code, data, patches
}

var div_reg_reg []byte = []byte{0x48, 0xF7}
var div_reg_reg_registerAdjustment byte = 0xF8

/*
 * The challenge of this instruction is that arguments
 * must be in the right registers before division occurs
 * I hate that there are so many stack operations involved,
 * but it seems the only way to consitently ensure that things
 * get where they need to be.
 */

func DivRegReg(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

	divisorReg := ins.Arguments[0]
	dividendReg := ins.Arguments[1]

	//Push all registers except the result register, divisor register, and dividend register
	result_reg := ins.Arguments[0]
	for _, cur_reg := range primaryRegisters {
		if cur_reg != result_reg && cur_reg != dividendReg && cur_reg != divisorReg {
			push_this_reg, _, _ := genAuxInstruction(PushReg, cur_reg)
			code = append(code, push_this_reg...)
		}
	}

	//Push divisor and dividend register to stack.
	push_divisor, _, _ := genAuxInstruction(PushReg, divisorReg)
	push_dividend, _, _ := genAuxInstruction(PushReg, dividendReg)
	code = append(code, push_dividend...)
	code = append(code, push_divisor...)

	//Pop divisor and dividend into appropriate rax and rcx
	pop_dividend, _, _ := genAuxInstruction(PopReg, pcode.REG1)
	pop_divisor, _, _ := genAuxInstruction(PopReg, pcode.REG3)
	code = append(code, pop_dividend...)
	code = append(code, pop_divisor...)

	//Clear rdx
	push_rdx, _, _ := genAuxInstruction(PushReg, pcode.REG4)
	clr_rdx, _, _ := genAuxInstruction(MovRegImm, pcode.REG4, 0x0)
	code = append(code, push_rdx...)
	code = append(code, clr_rdx...)

	//Divide
	code = append(code, div_reg_reg...)
	code = append(code, div_reg_reg_registerAdjustment|RCX)

	//Pop rdx
	pop_rdx, _, _ := genAuxInstruction(PopReg, pcode.REG4)
	code = append(code, pop_rdx...)

	//Result is now in rax; Move rax to result register
	mov_result, _, _ := genAuxInstruction(MovRegReg, result_reg, pcode.REG1)
	code = append(code, mov_result...)

	//Pop registers except result register
	for _, cur_reg := range primaryRegisters {
		if cur_reg != result_reg && cur_reg != dividendReg && cur_reg != divisorReg {
			push_this_reg, _, _ := genAuxInstruction(PopReg, cur_reg)
			code = append(code, push_this_reg...)
		}
	}

	return code, data, patches
}
