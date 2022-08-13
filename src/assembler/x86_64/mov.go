package x86_64

import (
	"turtlego/src/assembler/backpatch"
	"turtlego/src/pcode"
)

var mov_reg_imm []byte = []byte{0x48, 0xC7}

func MovRegImm(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code := []byte{}
	data := []byte{}

	code = append(code, mov_reg_imm...)
	code = append(code, singleRegisterEncoding(ins.Arguments[0]))
	code = append(code, mkIntByteArray(ins.Arguments[1])...)

	return code, data, []backpatch.BackPatch{}
}

var mov_reg_reg []byte = []byte{0x48, 0x89}

func MovRegReg(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code := []byte{}
	data := []byte{}

	code = append(code, mov_reg_reg...)
	code = append(code, dualRegisterEncoding(ins.Arguments[0], ins.Arguments[1]))

	return code, data, []backpatch.BackPatch{}
}

var mov_addr_reg []byte = []byte{0x48, 0x89}

func MovRegInAddrFromReg(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data := []byte{}, []byte{}
	src_reg := ins.Arguments[1]
	reg_with_dest_addr := ins.Arguments[0]

	reg := registerMap[src_reg] << 3
	r_m := registerMap[reg_with_dest_addr]

	out_byte := (reg | r_m)

	code = append(code, mov_addr_reg...)
	code = append(code, out_byte)

	return code, data, []backpatch.BackPatch{}
}

var mov_reg_addr []byte = []byte{0x48, 0x8b}

func MovRegFromRegInAddr(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data := []byte{}, []byte{}
	dest_reg := ins.Arguments[0]
	reg_with_src_addr := ins.Arguments[1]

	reg := registerMap[dest_reg] << 3
	r_m := registerMap[reg_with_src_addr]

	out_byte := (reg | r_m) //Mod bits are 00, so we don't have to OR in any extr values for it

	code = append(code, mov_reg_addr...)
	code = append(code, out_byte)

	return code, data, []backpatch.BackPatch{}
}
