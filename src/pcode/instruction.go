package pcode

import "fmt"

type Instruction struct {
	Opcode    byte
	Arguments []int
}

const (
	HLT byte = iota

	REG1
	REG2
	REG3
	REG4

	LOADINT
)

var InstructionMnemonicMap = map[byte]string{
	HLT:     "halt",
	LOADINT: "loadint",
}

func (i *Instruction) Stringify() string {
	v := InstructionMnemonicMap[i.Opcode]
	v += ", "
	for _, arg := range i.Arguments {
		v += fmt.Sprintf("%d ", arg)
	}
	return v
}

func (i *Instruction) Print() {
	fmt.Println(i.Stringify())
}
