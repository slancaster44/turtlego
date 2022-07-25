package assembler

import "turtlego/src/pcode"

type Assembler struct {
	outputFile string
	code       []byte
	data       []byte

	output []byte

	pcode *pcode.Program
}

const (
	X86_64 byte = iota
)

const (
	ELF byte = iota
)

func NewAssembler(pcode *pcode.Program) *Assembler {
	a := &Assembler{}

	a.pcode = pcode

	a.code = []byte{}
	a.data = []byte{}
	a.output = []byte{}

	return a
}
