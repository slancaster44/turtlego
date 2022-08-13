package x86_64

import (
	"strconv"
	"turtlego/src/assembler/backpatch"
	"turtlego/src/pcode"
)

type assemblerFn func(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch)

var builtinsMap map[int]assemblerFn = map[int]assemblerFn{
	pcode.BUILTIN_PRINT: printBuiltin,
}

func Builtin(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	fn, ok := builtinsMap[ins.Arguments[0]]

	if !ok {
		panic("No builtin assembler fn for: " + strconv.Itoa(ins.Arguments[0]))
	}

	return fn(ins)
}

var syscall = []byte{0x0F, 0x05}

func printBuiltin(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data := []byte{}, []byte{}

	//Save Registers
	code = append(code, codeToPushPrimaryRegisters()...)

	//Save value to be printed to top of stack
	reg_to_printed := ins.Arguments[1]
	i, _, _ := genAuxInstruction(PushReg, reg_to_printed)
	code = append(code, i...)

	//Set System Call
	//mov rsi, rsp
	i, _, _ = genAuxInstruction(MovRegReg, pcode.REG4+1, pcode.STACK_POINTER)
	code = append(code, i...)

	//mov rax, 4
	i, _, _ = genAuxInstruction(MovRegImm, pcode.REG1, 0x01)
	code = append(code, i...)

	//mov rbx, 1
	i, _, _ = genAuxInstruction(MovRegImm, pcode.REG2, 0x01)
	code = append(code, i...)

	//mov rdx, item length
	i, _, _ = genAuxInstruction(MovRegImm, pcode.REG4, STACK_VAR_SIZE)
	code = append(code, i...)

	//make syscall
	code = append(code, syscall...)

	//Pop value printed from stack
	i, _, _ = genAuxInstruction(PopReg, reg_to_printed)
	code = append(code, i...)

	//Restore Registers
	code = append(code, codeToPopPrimaryRegisters()...)

	return code, data, []backpatch.BackPatch{}
}
