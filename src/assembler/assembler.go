package assembler

import (
	"fmt"
	"os"
	"turtlego/src/assembler/x86_64"
	"turtlego/src/pcode"
)

type AssemblerFn func(pcode.Instruction) ([]byte, []byte)

type Assembler struct {
	outputFile string
	code       []byte
	data       []byte

	output []byte

	pcodeX86map map[byte]AssemblerFn
	exitFnsMap  map[byte]func() []byte

	pcode *pcode.Program
}

const (
	X86_64 byte = iota
)

func NewAssembler(pc *pcode.Program) *Assembler {
	a := &Assembler{}

	a.pcode = pc

	a.code = []byte{}
	a.data = []byte{}
	a.output = []byte{}

	a.pcodeX86map = map[byte]AssemblerFn{
		pcode.LOADINT:         x86_64.MovRegImm,
		pcode.ADD_REG_INT_INT: x86_64.AddImmReg,
		pcode.ADD_REG_REG_INT: x86_64.AddRegReg,
		pcode.SUB_REG_INT_INT: x86_64.SubImmReg,
		pcode.PUSH_REG:        x86_64.PushReg,
		pcode.POP:             x86_64.PopReg, //TODO: Change to POP
	}
	a.exitFnsMap = map[byte]func() []byte{
		X86_64: x86_64.ExitX86,
	}

	return a
}

func (a *Assembler) Assemble(instructionSet byte) {
	for _, instruction := range a.pcode.Instructions {
		fn, ok := a.pcodeX86map[instruction.Opcode] //TODO: Expand for other architectures
		if !ok {
			a.raiseError("Generation", "x86 code generation at: "+pcode.InstructionMnemonicMap[instruction.Opcode])
		}
		newCode, newData := fn(*instruction)
		a.code = append(a.code, newCode...)
		a.data = append(a.data, newData...)
	}

	a.code = append(a.code, a.exitFnsMap[instructionSet]()...)
}

func (a *Assembler) raiseError(name, msg string) {
	fmt.Printf("%sError: %s\n", name, msg)
	os.Exit(2)
}
