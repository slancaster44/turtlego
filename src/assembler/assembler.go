package assembler

import (
	"fmt"
	"os"
	"turtlego/src/pcode"
)

type Assembler struct {
	outputFile string
	code       []byte
	data       []byte

	output []byte

	pcodeX86map map[byte]func(pcode.Instruction)

	pcode *pcode.Program
}

const (
	X86_64 byte = iota
)

const (
	ELF byte = iota
)

func NewAssembler(pc *pcode.Program) *Assembler {
	a := &Assembler{}

	a.pcode = pc

	a.code = []byte{}
	a.data = []byte{}
	a.output = []byte{}

	a.pcodeX86map = map[byte]func(pcode.Instruction){
		pcode.LOADINT:         a.assembleMovInt,
		pcode.PUSH_REG:        a.assemblePushReg,
		pcode.POP:             a.assemblePop,
		pcode.ADD_REG_REG_INT: a.assembleAddRegRegInt,
		pcode.ADD_REG_INT_INT: a.assembleAddRegIntInt,
		pcode.MOV_REG_REG:     a.assembleMovRegReg,
		pcode.SUB_REG_REG_INT: a.assembleSubRegRegInt,
		pcode.SUB_REG_INT_INT: a.assembleSubRegIntInt,
		pcode.MUL_REG_REG_INT: a.assembleMulRegRegInt,
		pcode.MUL_REG_INT_INT: a.assembleMulRegIntInt,
		pcode.DIV_REG_REG_INT: a.assembleDivRegRegInt,
		pcode.DIV_REG_INT_INT: a.assembleDivRegIntInt,
	}

	return a
}

func (a *Assembler) Assemble() {
	for _, instruction := range a.pcode.Instructions {
		fn, ok := a.pcodeX86map[instruction.Opcode]
		if !ok {
			a.raiseError("Generation", "x86 code generation "+pcode.InstructionMnemonicMap[instruction.Opcode])
		}
		fn(*instruction)
	}
}

func (a *Assembler) raiseError(name, msg string) {
	fmt.Printf("%sError: %s\n", name, msg)
	os.Exit(2)
}
