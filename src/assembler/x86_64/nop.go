package x86_64

import (
	"turtlego/src/assembler/backpatch"
	"turtlego/src/pcode"
)

func Nop(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	return []byte{0x90}, []byte{}, []backpatch.BackPatch{}
}
