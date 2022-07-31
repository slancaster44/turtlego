package assembler

import (
	"encoding/binary"
	"fmt"
	"turtlego/src/pcode"
)

const (
	REX_W_PREFIX  byte = 0x48
	MOV_INT            = 0xC7
	MOV_REG_REG        = 0x89
	REGISTER_BASE      = 0xC0
	PUSH_REG_BASE      = 0x50
	POP_BASE           = 0x58

	ADD_REGS    = 0x01
	ADD_RAX_INT = 0x05
	ADD_REG_INT = 0x81

	SUB_REGS      = 0x29
	SUB_RAX_INT   = 0x2D
	SUB_REG_INT   = 0x81
	SUB_INDICATOR = 0x28

	RAX = 0x00
	RCX = 0x01
	RDX = 0x02
	RBX = 0x03
)

var registerMap = map[int]byte{
	pcode.REG1: RAX,
	pcode.REG2: RBX,
	pcode.REG3: RCX,
	pcode.REG4: RDX,
}

func dualRegisterEncoding(d, s int) byte {

	source, ok := registerMap[s]
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

func (a *Assembler) assembleMovRegReg(ins pcode.Instruction) {
	a.code = append(a.code, REX_W_PREFIX)
	a.code = append(a.code, MOV_REG_REG)

	i := dualRegisterEncoding(ins.Arguments[0], ins.Arguments[1])
	a.code = append(a.code, i)
}

func (a *Assembler) assembleMovInt(ins pcode.Instruction) {
	a.code = append(a.code, REX_W_PREFIX)
	a.code = append(a.code, MOV_INT)
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
	a.code = append(a.code, REX_W_PREFIX)
	a.code = append(a.code, ADD_REGS)
	reg := dualRegisterEncoding(ins.Arguments[0], ins.Arguments[1])
	a.code = append(a.code, reg)
}

func (a *Assembler) assembleAddRegIntInt(ins pcode.Instruction) {
	a.code = append(a.code, REX_W_PREFIX)

	//For whatever reason, our x86 overlords chose this asinine encodeing
	//where ONLY rax is different
	if ins.Arguments[0] == pcode.REG1 {
		a.code = append(a.code, ADD_RAX_INT)
	} else {
		a.code = append(a.code, ADD_REG_INT)
		a.code = append(a.code, singleRegisterEncoding(ins.Arguments[0]))
	}

	n := convertNumToByteArray(ins.Arguments[1])
	a.code = append(a.code, n...)
}

func (a *Assembler) assembleSubRegRegInt(ins pcode.Instruction) {
	a.code = append(a.code, REX_W_PREFIX)
	a.code = append(a.code, SUB_REGS)
	regs := dualRegisterEncoding(ins.Arguments[0], ins.Arguments[1])
	a.code = append(a.code, regs)
}

func (a *Assembler) assembleSubRegIntInt(ins pcode.Instruction) {
	a.code = append(a.code, REX_W_PREFIX)

	if ins.Arguments[0] == pcode.REG1 {
		a.code = append(a.code, SUB_RAX_INT)
	} else {
		a.code = append(a.code, SUB_REG_INT)
		a.code = append(a.code, SUB_INDICATOR+singleRegisterEncoding(ins.Arguments[0]))
	}

	n := convertNumToByteArray(ins.Arguments[1])
	a.code = append(a.code, n...)
}

func (a *Assembler) assembleMulRegRegInt(ins pcode.Instruction) {}

func (a *Assembler) assembleMulRegIntInt(ins pcode.Instruction) {}

func (a *Assembler) assembleDivRegRegInt(ins pcode.Instruction) {}

func (a *Assembler) assembleDivRegIntInt(ins pcode.Instruction) {}
