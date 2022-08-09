package x86_64

import "turtlego/src/pcode"

var push_reg_adjust byte = 0x50

func PushReg(ins pcode.Instruction) ([]byte, []byte) {
	code, data := []byte{}, []byte{}

	r := ins.Arguments[0]
	code = append(code, push_reg_adjust|registerMap[r])

	return code, data
}

var pop_reg_adjust byte = 0x58

func PopReg(ins pcode.Instruction) ([]byte, []byte) {
	code, data := []byte{}, []byte{}

	r := ins.Arguments[0]
	code = append(code, pop_reg_adjust|registerMap[r])

	return code, data
}

func codeToPushPrimaryRegisters() []byte {
	code := []byte{}

	push_rax, _ := genAuxInstruction(PushReg, pcode.REG1)
	push_rbx, _ := genAuxInstruction(PushReg, pcode.REG2)
	push_rcx, _ := genAuxInstruction(PushReg, pcode.REG3)
	push_rdx, _ := genAuxInstruction(PushReg, pcode.REG4)

	code = append(code, push_rax...)
	code = append(code, push_rbx...)
	code = append(code, push_rcx...)
	code = append(code, push_rdx...)

	return code
}

func codeToPopPrimaryRegisters() []byte {
	code := []byte{}

	pop_rdx, _ := genAuxInstruction(PopReg, pcode.REG4)
	pop_rcx, _ := genAuxInstruction(PopReg, pcode.REG3)
	pop_rbx, _ := genAuxInstruction(PopReg, pcode.REG2)
	pop_rax, _ := genAuxInstruction(PopReg, pcode.REG1)

	code = append(code, pop_rdx...)
	code = append(code, pop_rcx...)
	code = append(code, pop_rbx...)
	code = append(code, pop_rax...)

	return code
}
