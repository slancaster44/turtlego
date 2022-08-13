package x86_64

import (
	"turtlego/src/assembler/backpatch"
	"turtlego/src/pcode"
)

var push_reg_adjust byte = 0x50

func PushReg(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data := []byte{}, []byte{}

	r := ins.Arguments[0]
	code = append(code, push_reg_adjust|registerMap[r])

	return code, data, []backpatch.BackPatch{}
}

var push_int byte = 0x68

func PushInt(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data := []byte{}, []byte{}

	imm := ins.Arguments[0]
	code = append(code, push_int)
	code = append(code, mkIntByteArray(imm)...)

	return code, data, []backpatch.BackPatch{}
}

var pop_reg_adjust byte = 0x58

func PopReg(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data := []byte{}, []byte{}

	r := ins.Arguments[0]
	code = append(code, pop_reg_adjust|registerMap[r])

	return code, data, []backpatch.BackPatch{}
}

func codeToPushPrimaryRegisters() []byte {
	code := []byte{}

	push_rax, _, _ := genAuxInstruction(PushReg, pcode.REG1)
	push_rbx, _, _ := genAuxInstruction(PushReg, pcode.REG2)
	push_rcx, _, _ := genAuxInstruction(PushReg, pcode.REG3)
	push_rdx, _, _ := genAuxInstruction(PushReg, pcode.REG4)

	code = append(code, push_rax...)
	code = append(code, push_rbx...)
	code = append(code, push_rcx...)
	code = append(code, push_rdx...)

	return code
}

func codeToPopPrimaryRegisters() []byte {
	code := []byte{}

	pop_rdx, _, _ := genAuxInstruction(PopReg, pcode.REG4)
	pop_rcx, _, _ := genAuxInstruction(PopReg, pcode.REG3)
	pop_rbx, _, _ := genAuxInstruction(PopReg, pcode.REG2)
	pop_rax, _, _ := genAuxInstruction(PopReg, pcode.REG1)

	code = append(code, pop_rdx...)
	code = append(code, pop_rcx...)
	code = append(code, pop_rbx...)
	code = append(code, pop_rax...)

	return code
}
