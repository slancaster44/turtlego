package assembler

import "turtlego/src/pcode"

type Assembler struct {
	outputFile string
	code       []byte
	data       []byte

	pcode pcode.Program

	instructionSet   byte
	targetFileformat byte
}

const (
	X86_64 byte = iota
)

const (
	ELF byte = iota
)

func NewAssembler(outFilename string,
	instructionSet, targetFileformat byte,
	pcode pcode.Program) *Assembler {
	a := &Assembler{}

	a.outputFile = outFilename
	a.instructionSet = instructionSet
	a.targetFileformat = targetFileformat
	a.pcode = pcode

	a.code = []byte{}
	a.data = []byte{}

	return a
}
