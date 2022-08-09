package x86_64

import (
	"strconv"
	"turtlego/src/pcode"
)

type assemblerFn func(ins pcode.Instruction) ([]byte, []byte)

var builtinsMap map[int]assemblerFn = map[int]assemblerFn{
	pcode.BUILTIN_PRINT: printBuiltin,
}

func Builtin(ins pcode.Instruction) ([]byte, []byte) {
	fn, ok := builtinsMap[ins.Arguments[0]]

	if !ok {
		panic("No builtin assembler fn for: " + strconv.Itoa(ins.Arguments[0]))
	}

	return fn(ins)
}

var syscall = []byte{0x0F, 0x05}

func printBuiltin(ins pcode.Instruction) ([]byte, []byte) {
	code, data := []byte{}, []byte{}

	//Save Registers
	code = append(code, codeToPushPrimaryRegisters()...)

	//Save value to be printed to top of stack
	reg_to_printed := ins.Arguments[1]
	i, _ := genAuxInstruction(PushReg, reg_to_printed)
	code = append(code, i...)

	//Set System Call
	//mov rsi, rsp
	i, _ = genAuxInstruction(MovRegReg, pcode.REG4+1, pcode.STACK_POINTER)
	code = append(code, i...)

	//mov rax, 4
	i, _ = genAuxInstruction(MovRegImm, pcode.REG1, 0x01)
	code = append(code, i...)

	//mov rbx, 1
	i, _ = genAuxInstruction(MovRegImm, pcode.REG2, 0x01)
	code = append(code, i...)

	//mov rdx, item length
	i, _ = genAuxInstruction(MovRegImm, pcode.REG4, STACK_VAR_SIZE)
	code = append(code, i...)

	//int 80h
	code = append(code, syscall...)

	//Restore Registers
	code = append(code, codeToPopPrimaryRegisters()...)

	return code, data
}
