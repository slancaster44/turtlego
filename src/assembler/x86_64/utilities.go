package x86_64

import (
	"encoding/binary"
	"fmt"
	"turtlego/src/assembler/backpatch"
	"turtlego/src/pcode"
)

// TODO: Ensure this is producing correct values
func mkIntByteArray(num int) []byte {
	out := make([]byte, 4)
	binary.LittleEndian.PutUint32(out, uint32(num))
	return out
}

const (
	REGISTER_BASE byte = 0xC0

	RAX = 0x00
	RCX = 0x01
	RDX = 0x02
	RBX = 0x03

	RSI = 0x06

	SP = 0x04
	BP = 0x05
)

var registerMap = map[int]byte{
	pcode.STACK_POINTER:           SP,
	pcode.STACK_FRAME_POINTER_REG: BP,
	pcode.REG1:                    RAX,
	pcode.REG2:                    RBX,
	pcode.REG3:                    RCX,
	pcode.REG4:                    RDX,
	pcode.REG4 + 1:                RSI,
}

const STACK_VAR_SIZE = 0x08

// Warning: At this point, the encoding assumes that
// the selected instruction encodes its registers
// as 11XX XXXX C0 + Register Encoding
func dualRegisterEncoding(d, s int) byte {

	source, ok := registerMap[s]
	if !ok {
		panic(fmt.Sprintf("Failed to lookup register '%d' for dual register encoding", s))
	}

	dest, ok := registerMap[d]
	if !ok {
		panic(fmt.Sprintf("Failed to look up register '%d' for dual register encoding", d))
	}

	return REGISTER_BASE + (8 * source) + dest
}

func singleRegisterEncoding(r int) byte {
	rval, ok := registerMap[r]
	if !ok {
		fmt.Println(registerMap)
		panic(fmt.Sprintf("Failed to lookup reigster '%d' for single reigster encoding", r))
	}
	return REGISTER_BASE + rval
}

func genAuxInstruction(fn assemblerFn, args ...int) ([]byte, []byte, []backpatch.BackPatch) {
	return fn(pcode.Instruction{0, args})
}

func MkTrueBackPatchAddress(targetAddress int, addressOfCode int) []byte {
	return mkIntByteArray(targetAddress - addressOfCode - 0x06) //TODO: 0x06 WTF??????!?!!!?!?
}
