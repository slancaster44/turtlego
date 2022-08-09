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

func printBuiltin(ins pcode.Instruction) ([]byte, []byte) {
	code, data := []byte{}, []byte{}

	return code, data
}
