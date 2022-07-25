package assembler

import "turtlego/src/pcode"

const (
	MOV            byte = 0x48
	MOV_INT_PREFIX      = 0xC7
)

var registerMap = map[int]byte{}

func (a *Assembler) assembleMovInt(ins pcode.Instruction) {
	a.code = append(a.code, MOV)
	a.code = append(a.code, MOV_INT_PREFIX)
}
