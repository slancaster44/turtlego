package main

import (
	"bufio"
	"encoding/binary"
	"os"
	"turtlego/lexer"
	"turtlego/parser"
	"turtlego/riscv"
	"turtlego/source"
	"turtlego/tokens"
)

func main() {
	s := source.New("doc/test.s")
	l := lexer.New(s)
	l.SetMaps(lexer.Asm_singleChar, map[string]byte{}, lexer.Asm_keyword)
	p := parser.New(l)
	p.SetSkipables([]byte{tokens.EOL})
	p.SetAsmMaps()

	program := p.ParsePogram()
	a := riscv.NewAssembler()
	a.Assemble_ASM_Program(program)
	res := a.GenFlatBinary(0)

	f, _ := os.Create("output.bin")
	w := bufio.NewWriter(f)
	binary.Write(w, binary.LittleEndian, res)
	w.Flush()
}
