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

	RDI = 0x07
	RSI = 0x06

	R8  = 0x00
	R9  = 0x01
	R10 = 0x02

	HP = 0x03 //Heap pointer register (r11)
	HS = 0x04 //Heap size register (r12)

	SP = 0x04
	BP = 0x05
)

var primaryRegisters []int = []int{pcode.REG1, pcode.REG2, pcode.REG3, pcode.REG4}

var registerMap = map[int]byte{
	pcode.STACK_POINTER:           SP,
	pcode.STACK_FRAME_POINTER_REG: BP,
	pcode.REG1:                    RAX,
	pcode.REG2:                    RBX,
	pcode.REG3:                    RCX,
	pcode.REG4:                    RDX,
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

func codeToLoadImmDirectly(r byte, n int) []byte {
	code := []byte{}

	code = append(code, mov_reg_imm...)
	code = append(code, REGISTER_BASE+r)
	code = append(code, mkIntByteArray(n)...)

	return code
}

var mov_ereg_imm []byte = []byte{0x49, 0xC7}

func codeMovExtendedRegImm(r byte, n int) []byte {
	code := []byte{}

	code = append(code, mov_ereg_imm...)
	code = append(code, REGISTER_BASE|r)

	code = append(code, mkIntByteArray(n)...)

	return code
}

var mov_ereg_reg []byte = []byte{0x49, 0x89}

func codeMovExtendedRegNormalReg(er1, r2 byte) []byte {

	reg_encoding := REGISTER_BASE + (8 * r2) + er1
	code := append(mov_ereg_reg, reg_encoding)

	return code
}

func codeToCopyRegDirectly(r1, r2 byte) []byte {
	code := []byte{}

	code = append(code, mov_reg_reg...)
	code = append(code, REGISTER_BASE+(8*r2)+r1)

	return code
}

func genAuxInstruction(fn assemblerFn, args ...int) ([]byte, []byte, []backpatch.BackPatch) {
	return fn(pcode.Instruction{0, args})
}

func MkTrueBackPatchAddress(targetAddress int, addressOfCode int) []byte {
	return mkIntByteArray(targetAddress - addressOfCode)
}
