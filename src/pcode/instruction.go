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
	MOV_REG_ADDRESS_REG //Move from a register to the address in a register
	MOV_REG_REG_ADDRESS //Move from an address in a register to a register

	BUILTIN_CALL //bc <builtin_id> <reg_with_arg>
)

const (
	STACK_FRAME_POINTER_REG int = -2
	STACK_POINTER           int = -1
	REG1                    int = 0
	REG2                    int = 1
	REG3                    int = 2
	REG4                    int = 3
)

const (
	BUILTIN_PRINT int = iota
)

var InstructionMnemonicMap = map[byte]string{
	LOADINT:         "loadint",
	ADD_REG_INT_INT: "add_reg_int_int",
	PUSH_REG:        "push_reg",
	POP:             "pop",
	ADD_REG_REG_INT: "add_reg_reg_int",

	MOV_REG_REG:         "mov_reg_reg",
	MOV_REG_ADDRESS_REG: "mov_reg_address_reg",
	MOV_REG_REG_ADDRESS: "mov_reg_reg_address",

	SUB_REG_INT_INT: "sub_reg_int_int",
	SUB_REG_REG_INT: "sub_reg_reg_int",

	DIV_REG_INT_INT: "div_reg_int_int",
	DIV_REG_REG_INT: "div_reg_reg_int",

	MUL_REG_INT_INT: "mul_reg_int_int",
	MUL_REG_REG_INT: "mul_reg_reg_int",

	BUILTIN_CALL: "builtin_call",
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
