package pcode

import "fmt"

type Instruction struct {
	Opcode    byte
	Arguments []int
}

const (
	LOADINT byte = iota

	PUSH_REG
	PUSH_INT
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

	CMP_REG_INT
	JMZ //Jump if flags register says so
	JNZ //Jump if flags register says not to
	JMP

	BUILTIN_CALL //bc <builtin_id> <reg_with_arg>
	NOP

	BOOL_OR_REG_REG
	BOOL_OR_REG_IMM
	BOOL_AND_REG_REG
	BOOL_AND_REG_IMM

	EQ_REG_REG
	EQ_REG_IMM

	NE_REG_REG
	NE_REG_IMM

	LT_REG_REG
	LT_REG_IMM

	LE_REG_REG
	LE_REG_IMM

	GT_REG_REG
	GT_REG_IMM
	GE_REG_REG
	GE_REG_IMM

	MK_HEAP
	ALLOC
	DEALLOC
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
	PUSH_INT:        "push_int",
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

	NOP: "nop",
	JMP: "jmp",
	JMZ: "jmz",

	CMP_REG_INT: "cmp_reg_int",

	BUILTIN_CALL: "builtin_call",

	BOOL_AND_REG_IMM: "bool_and_reg_imm",
	BOOL_AND_REG_REG: "bool_and_reg_reg",
	BOOL_OR_REG_IMM:  "bool_or_reg_imm",
	BOOL_OR_REG_REG:  "bool_or_reg_reg",

	EQ_REG_IMM: "eq_reg_imm",
	EQ_REG_REG: "eq_reg_reg",
	NE_REG_REG: "ne_reg_reg",
	NE_REG_IMM: "ne_reg_imm",

	LT_REG_IMM: "lt_reg_imm",
	LT_REG_REG: "lt_reg_reg",

	LE_REG_REG: "le_reg_reg",
	LE_REG_IMM: "le_reg_imm",
	GT_REG_IMM: "gt_reg_imm",
	GT_REG_REG: "gt_reg_reg",
	GE_REG_IMM: "ge_reg_imm",
	GE_REG_REG: "ge_reg_reg",

	MK_HEAP: "mk_heap",
	ALLOC:   "alloc", //alloc <reg> <size-in-qwords>
	DEALLOC: "dealloc",
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

func MkInstruction(opcode byte, args ...int) *Instruction {
	return &Instruction{opcode, args}
}
