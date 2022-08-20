package assembler

import (
	"fmt"
	"os"
	"turtlego/src/assembler/backpatch"
	"turtlego/src/assembler/x86_64"
	"turtlego/src/pcode"
)

// Functions that will create code for different instruction sets.
// Returns code, data, and patches
type AssemblerFn func(pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch)

type Assembler struct {
	outputFile string
	code       []byte
	data       []byte

	pcodeAddressToRealAddress map[int]int
	backPatches               []backpatch.BackPatch

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

	a.pcodeAddressToRealAddress = make(map[int]int)

	a.pcodeX86map = map[byte]AssemblerFn{
		pcode.LOADINT:             x86_64.MovRegImm,
		pcode.MOV_REG_REG:         x86_64.MovRegReg,
		pcode.MOV_REG_ADDRESS_REG: x86_64.MovRegInAddrFromReg,
		pcode.MOV_REG_REG_ADDRESS: x86_64.MovRegFromRegInAddr,
		pcode.ADD_REG_INT_INT:     x86_64.AddImmReg,
		pcode.ADD_REG_REG_INT:     x86_64.AddRegReg,
		pcode.SUB_REG_INT_INT:     x86_64.SubImmReg,
		pcode.SUB_REG_REG_INT:     x86_64.SubRegReg,
		pcode.MUL_REG_INT_INT:     x86_64.MulImmReg,
		pcode.PUSH_REG:            x86_64.PushReg,
		pcode.POP:                 x86_64.PopReg,
		pcode.BUILTIN_CALL:        x86_64.Builtin,
		pcode.PUSH_INT:            x86_64.PushInt,
		pcode.CMP_REG_INT:         x86_64.CmpRegInt,
		pcode.JMP:                 x86_64.Jump,
		pcode.JNZ:                 x86_64.JumpIfNotZero,
		pcode.JMZ:                 x86_64.JumpIfZero,
		pcode.NOP:                 x86_64.Nop,
		pcode.MUL_REG_REG_INT:     x86_64.MulRegReg,
		pcode.DIV_REG_REG_INT:     x86_64.DivRegReg,
		pcode.DIV_REG_INT_INT:     x86_64.DivImmReg,
		pcode.BOOL_OR_REG_IMM:     x86_64.OrRegImm,
		pcode.BOOL_OR_REG_REG:     x86_64.OrRegReg,
		pcode.BOOL_AND_REG_REG:    x86_64.AndRegReg,
		pcode.BOOL_AND_REG_IMM:    x86_64.AndRegImm,

		//These functions are all generated using a higher-order function
		//The function that creates these functions returns a compatiable
		//type with AssemblerFn, but it must be converted here to be in the map
		pcode.EQ_REG_REG: AssemblerFn(x86_64.EqRegReg),
		pcode.EQ_REG_IMM: AssemblerFn(x86_64.EqRegImm),
		pcode.NE_REG_IMM: AssemblerFn(x86_64.NeRegImm),
		pcode.NE_REG_REG: AssemblerFn(x86_64.NeRegReg),
		pcode.LT_REG_IMM: AssemblerFn(x86_64.LtRegImm),
		pcode.LT_REG_REG: AssemblerFn(x86_64.LtRegReg),
		pcode.LE_REG_IMM: AssemblerFn(x86_64.LeRegImm),
		pcode.LE_REG_REG: AssemblerFn(x86_64.LeRegReg),
		pcode.GE_REG_IMM: AssemblerFn(x86_64.GeRegImm),
		pcode.GE_REG_REG: AssemblerFn(x86_64.GeRegReg),
		pcode.GT_REG_IMM: AssemblerFn(x86_64.GtRegImm),
		pcode.GT_REG_REG: AssemblerFn(x86_64.GtRegReg),
	}
	a.exitFnsMap = map[byte]func() []byte{
		X86_64: x86_64.ExitX86,
	}

	return a
}

func (a *Assembler) Assemble(instructionSet byte) {
	for pcodeAddress, instruction := range a.pcode.Instructions {
		fn, ok := a.pcodeX86map[instruction.Opcode] //TODO: Expand for other architectures
		if !ok {
			a.raiseError("Generation", "x86 code generation at: "+pcode.InstructionMnemonicMap[instruction.Opcode])
		}
		newCode, newData, backpatches := fn(*instruction)

		a.pcodeAddressToRealAddress[pcodeAddress] = len(a.code)
		//fmt.Printf("%d %x\n", pcodeAddress, 0x400000+len(a.code)+0x40+(2*0x38))

		for _, patch := range backpatches {
			patch.LocationOfAddressToPatch += len(a.code)
			patch.LocationOfInstructionPatched += len(a.code)
			a.backPatches = append(a.backPatches, patch)
		}

		a.code = append(a.code, newCode...)
		a.data = append(a.data, newData...)
	}

	a.applyBackpatches()
	a.code = append(a.code, a.exitFnsMap[instructionSet]()...)
}

func (a *Assembler) raiseError(name, msg string) {
	fmt.Printf("%sError: %s\n", name, msg)
	os.Exit(2)
}

func (a *Assembler) applyBackpatches() {
	for _, patch := range a.backPatches {
		addressInCode := a.pcodeAddressToRealAddress[patch.PcodeAddressToPatchTo]
		trueAddress := x86_64.MkTrueBackPatchAddress(addressInCode, patch.LocationOfInstructionPatched)

		for iter, byt := range trueAddress {
			a.code[patch.LocationOfAddressToPatch+iter] = byt
		}
	}
}
