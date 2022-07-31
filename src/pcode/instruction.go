package pcode

import "fmt"

type Instruction struct {
	Opcode    byte
	Arguments []int
}

const (
	LOADINT byte = iota

	PUSH_REG
	POP

	//These instructions take the form:
	//operator_typeOfFirstArg_typeOfSecondArg_returnValue
	ADD_REG_INT_INT
	ADD_REG_REG_INT

	SUB_REG_INT_INT
	SUB_REG_REG_INT

	DIV_REG_INT_INT
	DIV_REG_REG_INT

	MUL_REG_INT_INT
	MUL_REG_REG_INT

	MOV_REG_REG
)

const (
	REG1 int = iota
	REG2
	REG3
	REG4
)

var InstructionMnemonicMap = map[byte]string{
	LOADINT:         "loadint",
	ADD_REG_INT_INT: "add_reg_int_int",
	PUSH_REG:        "push_reg",
	POP:             "pop",
	ADD_REG_REG_INT: "add_reg_reg_int",
	MOV_REG_REG:     "mov_reg_reg",
	SUB_REG_INT_INT: "sub_reg_int_int",
	SUB_REG_REG_INT: "sub_reg_reg_int",

	DIV_REG_INT_INT: "div_reg_int_int",
	DIV_REG_REG_INT: "div_reg_reg_int",

	MUL_REG_INT_INT: "mul_reg_int_int",
	MUL_REG_REG_INT: "mul_reg_reg_int",
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
