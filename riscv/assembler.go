package riscv

import (
	"fmt"
	"os"
	"turtlego/ast"
	"turtlego/message"
)

type Ins_Tag uint16

const (
	ADDI Ins_Tag = iota
	ORI
)

type Instruction struct {
	Regs  []uint32
	Imms  []uint32
	AsmFn func([]uint32, []uint32) uint32
	Tag   Ins_Tag
}

func NewInstruction() Instruction {
	retVal := Instruction{}
	retVal.Regs = []uint32{}
	retVal.Imms = []uint32{}

	return retVal
}

type Reference struct {
	Label            string
	PC               uint64
	CodeOffset       uint64
	ProduceReference func(ins_addr uint64, label_addr uint64) uint32
	ImmIndex         int //The index into that instructions array of immeddiates that the result will be stored
}

type CompileUnit struct {
	Code []Instruction
	Data []byte

	DataLabels map[string]uint64
	CodeLabels map[string]uint64

	References []Reference
}

type Assembler struct {
	curCu            *CompileUnit
	allCus           []*CompileUnit
	curPC            uint64
	curTopLevelLabel string
	genMap           map[int]func(n ast.Node)
}

func NewAssembler() *Assembler {
	retVal := &Assembler{}
	retVal.curPC = 0
	retVal.genMap = map[int]func(n ast.Node){}
	retVal.allCus = []*CompileUnit{}

	return retVal
}

func (a *Assembler) Error(t, m string, n ast.Node) {
	message.RaiseError(t, m, n.Token())
}

func (a *Assembler) Assemble_ASM_Program(p ast.Program) *CompileUnit {
	a.curCu = &CompileUnit{}
	a.curCu.CodeLabels = map[string]uint64{}
	a.curCu.DataLabels = map[string]uint64{}

	a.genMap[ast.LABEL_NT] = a.CodeLabel
	a.genMap[ast.I_TYPE_NT] = a.IType
	a.genMap[ast.SB_TYPE_NT] = a.SBType
	a.genMap[ast.R_TYPE_NT] = a.RType
	a.genMap[ast.OTHER_TYPE_NT] = a.Other
	a.genMap[ast.DOT_LABEL_NT] = a.DotLabel

	for _, expr := range p {
		fn := a.genMap[expr.Type()]
		if fn == nil {
			fmt.Println("No assembler function for node: ", expr)
			os.Exit(1)
		}
		fn(expr)
	}

	a.allCus = append(a.allCus, a.curCu)
	a.genMap = map[int]func(ast.Node){}
	return a.curCu
}

func (a *Assembler) GenFlatBinary(codeBase uint64) []uint32 {
	output := []uint32{}

	for _, cu := range a.allCus {
		for _, ref := range cu.References {
			raw_value := cu.CodeLabels[ref.Label]
			ins_of_ref := cu.Code[ref.CodeOffset]
			val := ref.ProduceReference(ref.PC+codeBase, raw_value+codeBase)
			ins_of_ref.Imms[ref.ImmIndex] = val
		}

		for _, ins := range cu.Code {
			result := ins.AsmFn(ins.Regs, ins.Imms)
			output = append(output, result)
		}
	}

	return output
}
