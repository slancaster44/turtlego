package assembler

import (
	"encoding/binary"
	"fmt"
	"turtlego/src/pcode"
)

const (
	MOV            byte = 0x48
	MOV_INT_PREFIX      = 0xC7
	REGISTER_BASE       = 0xC0
	PUSH_REG_BASE       = 0x50
	POP_BASE            = 0x58
	ADD                 = 0x48
	ADD_REG_PREFIX      = 0x01

	RAX = 0x00
	RCX = 0x01
	RDX = 0x02
	RBX = 0x03

	//When encoding two registers into a single
	//argument, then the register of one of those
	//operands is different than normal
	SRC_RAX = 0x00
	SRC_RBX = 0x01
	SRC_RCX = 0x02
	SRC_RDX = 0x03
)

var registerMap = map[int]byte{
	int(pcode.REG1): RAX,
	int(pcode.REG2): RBX,
	int(pcode.REG3): RCX,
	int(pcode.REG4): RDX,
}

var srcRegisterMap = map[int]byte{
	int(pcode.REG1): SRC_RAX,
	int(pcode.REG2): SRC_RBX,
	int(pcode.REG3): SRC_RCX,
	int(pcode.REG4): SRC_RDX,
}

func dualRegisterEncoding(d, s int) byte {

	source, ok := srcRegisterMap[s]
	if !ok {
		panic(fmt.Sprintf("Failed to lookup register '%d'", s))
	}

	dest, ok := registerMap[d]
	if !ok {
		panic(fmt.Sprintf("Failed to look up register '%d'", d))
	}

	return REGISTER_BASE + (8 * source) + dest
}

func singleRegisterEncoding(r int) byte {
	rval, ok := registerMap[r]
	if !ok {
		panic(fmt.Sprintf("Failed to lookup reigster '%d'", r))
	}
	return REGISTER_BASE + rval
}

func convertNumToByteArray(num int) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(num))
	return buf
}

func (a *Assembler) assembleMovInt(ins pcode.Instruction) {
	a.code = append(a.code, MOV)
	a.code = append(a.code, MOV_INT_PREFIX)
	a.code = append(a.code, singleRegisterEncoding(ins.Arguments[0]))
	a.code = append(a.code, convertNumToByteArray(ins.Arguments[1])...)
}

func (a *Assembler) assemblePushReg(ins pcode.Instruction) {
	r := ins.Arguments[0]
	instruction := PUSH_REG_BASE + registerMap[r]
	a.code = append(a.code, instruction)
}

func (a *Assembler) assemblePop(ins pcode.Instruction) {
	r := ins.Arguments[0]
	instruction := POP_BASE + registerMap[r]
	a.code = append(a.code, instruction)
}

func (a *Assembler) assembleAddRegRegInt(ins pcode.Instruction) {
	a.code = append(a.code, ADD)
	a.code = append(a.code, ADD_REG_PREFIX)

	dest := int(registerMap[ins.Arguments[0]])
	src := int(registerMap[ins.Arguments[1]])
	reg := dualRegisterEncoding(dest, src)
	a.code = append(a.code, reg)
}
