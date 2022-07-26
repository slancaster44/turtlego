package pcode

import "fmt"

type Instruction struct {
	Opcode    byte
	Arguments []int
}

const (
	LOADINT byte = iota
	ADD

	PUSH_REG
	POP

	ADD_REG_REG_INT
)

const (
	REG1 byte = iota
	REG2
	REG3
	REG4
)

var InstructionMnemonicMap = map[byte]string{
	LOADINT:         "loadint",
	ADD:             "add",
	PUSH_REG:        "push_reg",
	POP:             "pop",
	ADD_REG_REG_INT: "add_reg_reg_int",
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
