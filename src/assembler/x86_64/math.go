package x86_64

import (
	"turtlego/src/assembler/backpatch"
	"turtlego/src/pcode"
)

//Various functions that translate arithmetic & boolean pcode instructions into x86

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

	push_imm_reg, _, _ := genAuxInstruction(PushReg, reg_for_imm)
	code = append(code, push_imm_reg...)

	mov, _, _ := genAuxInstruction(MovRegImm, reg_for_imm, imm)
	div, _, _ := genAuxInstruction(DivRegReg, reg, reg_for_imm)

	code = append(code, mov...)
	code = append(code, div...)

	pop_imm_reg, _, _ := genAuxInstruction(PopReg, reg_for_imm)
	code = append(code, pop_imm_reg...)

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

var or_reg_reg []byte = []byte{0x48, 0x09}

func OrRegReg(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

	code = append(code, or_reg_reg...)
	code = append(code, dualRegisterEncoding(ins.Arguments[0], ins.Arguments[1]))

	return code, data, patches
}

var or_reg_imm []byte = []byte{0x40, 0x81}

func OrRegImm(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

	code = append(code, or_reg_imm...)

	reg := 0xC8 | registerMap[ins.Arguments[0]]
	code = append(code, reg)

	code = append(code, mkIntByteArray(ins.Arguments[1])...)

	return code, data, patches
}

var and_reg_reg []byte = []byte{0x48, 0x21}

func AndRegReg(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

	code = append(code, and_reg_reg...)
	code = append(code, dualRegisterEncoding(ins.Arguments[0], ins.Arguments[1]))

	return code, data, patches
}

var and_reg_imm []byte = []byte{0x48, 0x81}
var and_reg_imm_adjustment byte = 0xE0

func AndRegImm(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

	code = append(code, and_reg_imm...)

	reg := and_reg_imm_adjustment | registerMap[ins.Arguments[0]]
	code = append(code, reg)

	imm := mkIntByteArray(ins.Arguments[1])
	code = append(code, imm...)

	return code, data, patches
}

var sete_reg []byte = []byte{0x0F, 0x94}

func EqRegReg(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

	//Cmp reg1 to reg2
	cmp_regs, _, _ := genAuxInstruction(CmpRegReg, ins.Arguments...)
	code = append(code, cmp_regs...)

	//Clear reg1
	clr_reg1, _, _ := genAuxInstruction(MovRegImm, ins.Arguments[0], 0x0)
	code = append(code, clr_reg1...)

	//sete reg1 lowest byte
	code = append(code, sete_reg...)
	code = append(code, singleRegisterEncoding(ins.Arguments[0]))

	return code, data, patches
}

func EqRegImm(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

	//Cmp reg1 to reg2
	cmp_regs, _, _ := genAuxInstruction(CmpRegInt, ins.Arguments...)
	code = append(code, cmp_regs...)

	//Clear reg1
	clr_reg1, _, _ := genAuxInstruction(MovRegImm, ins.Arguments[0], 0x0)
	code = append(code, clr_reg1...)

	//sete reg1 lowest byte
	code = append(code, sete_reg...)
	code = append(code, singleRegisterEncoding(ins.Arguments[0]))

	return code, data, patches
}

var setne_reg []byte = []byte{0x0F, 0x95}

func NeRegImm(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

	cmp_regs, _, _ := genAuxInstruction(CmpRegInt, ins.Arguments...)
	code = append(code, cmp_regs...)

	clr_reg1, _, _ := genAuxInstruction(MovRegImm, ins.Arguments[0], 0x0)
	code = append(code, clr_reg1...)

	//setne reg1 lowest byte
	code = append(code, setne_reg...)
	code = append(code, singleRegisterEncoding(ins.Arguments[0]))

	return code, data, patches
}

func NeRegReg(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

	cmp_regs, _, _ := genAuxInstruction(CmpRegReg, ins.Arguments...)
	code = append(code, cmp_regs...)

	clr_reg1, _, _ := genAuxInstruction(MovRegImm, ins.Arguments[0], 0x0)
	code = append(code, clr_reg1...)

	//setne reg1 lowest byte
	code = append(code, setne_reg...)
	code = append(code, singleRegisterEncoding(ins.Arguments[0]))

	return code, data, patches
}

func mkComparisonAssemblerFns(setOpCodes []byte) (assemblerFn, assemblerFn) {
	var reg_regFn assemblerFn = func(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
		code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

		cmp_regs, _, _ := genAuxInstruction(CmpRegReg, ins.Arguments...)
		code = append(code, cmp_regs...)

		clr_reg1, _, _ := genAuxInstruction(MovRegImm, ins.Arguments[0], 0x0)
		code = append(code, clr_reg1...)

		//setne reg1 lowest byte
		code = append(code, setOpCodes...)
		code = append(code, singleRegisterEncoding(ins.Arguments[0]))

		return code, data, patches
	}

	var reg_immFn assemblerFn = func(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
		code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

		cmp_regs, _, _ := genAuxInstruction(CmpRegInt, ins.Arguments...)
		code = append(code, cmp_regs...)

		clr_reg1, _, _ := genAuxInstruction(MovRegImm, ins.Arguments[0], 0x0)
		code = append(code, clr_reg1...)

		//setne reg1 lowest byte
		code = append(code, setOpCodes...)
		code = append(code, singleRegisterEncoding(ins.Arguments[0]))

		return code, data, patches
	}

	return reg_regFn, reg_immFn
}

var setlt_reg []byte = []byte{0x0F, 0x9C} //<
var LtRegReg, LtRegImm assemblerFn = mkComparisonAssemblerFns(setlt_reg)

var setle_reg []byte = []byte{0x0F, 0x9E} //<=
var LeRegReg, LeRegImm assemblerFn = mkComparisonAssemblerFns(setle_reg)

var setgt_reg []byte = []byte{0x0F, 0x9F} //>
var GtRegReg, GtRegImm assemblerFn = mkComparisonAssemblerFns(setgt_reg)

var setge_reg []byte = []byte{0x0F, 0x9D} //>=
var GeRegReg, GeRegImm assemblerFn = mkComparisonAssemblerFns(setge_reg)
