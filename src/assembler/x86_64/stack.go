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
