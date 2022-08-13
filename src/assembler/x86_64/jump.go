package x86_64

import (
	"turtlego/src/assembler/backpatch"
	"turtlego/src/pcode"
)

var cmpRegImm []byte = []byte{0x48, 0x81}
var registerAdjustment byte = 0x38
var zeroImm []byte = mkIntByteArray(0)

var jmz []byte = []byte{0x0F, 0x84}
var addressPlaceholder []byte = mkIntByteArray(0)

func JumpIfZero(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

	cmpRegister := singleRegisterEncoding(ins.Arguments[0])

	//Compare Register contents with zero
	code = append(code, cmpRegImm...)
	code = append(code, registerAdjustment+byte(cmpRegister))
	code = append(code, zeroImm...)

	//Execute Jmz instruction
	code = append(code, jmz...)
	code = append(code, addressPlaceholder...)

	//Apply backpatch to insert correct address once known
	locOfIns := len(cmpRegImm) + 1 + len(zeroImm) //the +1 is for the byte that incodes the register
	locOfAddressInIns := locOfIns + len(jmz)
	bp := backpatch.BackPatch{locOfAddressInIns, locOfIns, ins.Arguments[1]}
	patches = append(patches, bp)

	return code, data, patches
}

var jmp []byte = []byte{0xE9}

func Jump(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

	code = append(code, jmp...)
	code = append(code, addressPlaceholder...)

	locOfAddresInIns := len(jmp)
	bp := backpatch.BackPatch{locOfAddresInIns, -1, ins.Arguments[0]}
	patches = append(patches, bp)

	return code, data, patches
}
