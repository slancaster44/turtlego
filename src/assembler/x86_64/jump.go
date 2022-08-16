package x86_64

import (
	"turtlego/src/assembler/backpatch"
	"turtlego/src/pcode"
)

var jmz []byte = []byte{0x0F, 0x84}
var addressPlaceholder []byte = mkIntByteArray(0)

func JumpIfZero(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

	code = append(code, jmz...)
	code = append(code, addressPlaceholder...)

	locOfAddressInIns := len(jmz)
	addressLoc := len(jmz) + len(addressPlaceholder)
	bp := backpatch.BackPatch{locOfAddressInIns, addressLoc, ins.Arguments[0]}
	patches = append(patches, bp)

	return code, data, patches
}

var jz []byte = []byte{0x0F, 0x85}

func JumpIfNotZero(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

	code = append(code, jz...)
	code = append(code, addressPlaceholder...)

	locOfAddressInIns := len(jz)
	addressLoc := len(jz) + len(addressPlaceholder)
	bp := backpatch.BackPatch{locOfAddressInIns, addressLoc, ins.Arguments[0]}
	patches = append(patches, bp)

	return code, data, patches
}

var jmp []byte = []byte{0xE9}

func Jump(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

	code = append(code, jmp...)
	code = append(code, addressPlaceholder...)

	locOfAddresInIns := len(jmp)
	bp := backpatch.BackPatch{locOfAddresInIns, len(jmp) + len(addressPlaceholder), ins.Arguments[0]}
	patches = append(patches, bp)

	return code, data, patches
}
