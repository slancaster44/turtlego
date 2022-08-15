package x86_64

import (
	"turtlego/src/assembler/backpatch"
	"turtlego/src/pcode"
)

var cmp_reg_int []byte = []byte{0x48, 0x81}
var reg_offset byte = 0xF8

func CmpRegInt(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

	code = append(code, cmp_reg_int...)

	reg_encoded := reg_offset | registerMap[ins.Arguments[0]]
	code = append(code, reg_encoded)

	imm_encoded := mkIntByteArray(ins.Arguments[1])
	code = append(code, imm_encoded...)

	return code, data, patches
}
