package riscv

import (
	"turtlego/ast"
	"turtlego/riscv/encode"
)

func (a *Assembler) CodeLabel(n ast.Node) {
	label := n.(*ast.Label)

	_, ok := a.curCu.CodeLabels[label.Name]
	if ok {
		a.Error("LabelError", "Label name already used", n)
	}

	a.curTopLevelLabel = label.Name
	a.curCu.CodeLabels[label.Name] = a.curPC
}

func (a *Assembler) DotLabel(n ast.Node) {
	label := n.(*ast.DotLabel)
	label_name := a.curTopLevelLabel + "." + label.Name

	_, ok := a.curCu.CodeLabels[label_name]
	if ok {
		a.Error("LabelError", "Label name already used", n)
	}

	a.curCu.CodeLabels[label_name] = a.curPC
}

func (a *Assembler) IType(n ast.Node) {
	i_node := n.(*ast.I_Type)
	ins := NewInstruction()
	ins.Regs = []uint32{uint32(i_node.Rd.Value), uint32(i_node.Rs1.Value)}

	if i_node.Imm.Type() == ast.INTEGER_NT {
		ins.Imms = []uint32{uint32(i_node.Imm.(*ast.Integer).Value)}
	} else { //TODO: Labels
		a.Error("ValueError", "Expected Integer or label", n)
	}

	switch i_node.Mnemonic {
	case "addi":
		ins.AsmFn = encode.Addi
	default:
		a.Error("InsError", "No known asm fn", n)
	}

	a.curPC += 4
	a.curCu.Code = append(a.curCu.Code, ins)
}

func (a *Assembler) SBType(n ast.Node) {
	//TODO: Handle negative operator?

	sb_node := n.(*ast.SB_Type)
	ins := NewInstruction()
	ins.Regs = []uint32{uint32(sb_node.Rs1.Value), uint32(sb_node.Rs2.Value)}
	ins.Imms = []uint32{0}

	if sb_node.Imm.Type() == ast.DOT_LABEL_NT {
		dot_label := sb_node.Imm.(*ast.DotLabel)
		label := a.curTopLevelLabel + "." + dot_label.Name
		a.curCu.References = append(a.curCu.References, a.makePCRelReference(label, 0))
	} else if sb_node.Imm.Type() == ast.LABEL_NT {
		label := sb_node.Imm.(*ast.Label)
		ref := a.makePCRelReference(label.Name, 0)
		a.curCu.References = append(a.curCu.References, ref)
	} else if sb_node.Imm.Type() == ast.INTEGER_NT {
		ins.Imms[0] = uint32(sb_node.Imm.(*ast.Integer).Value)
	}

	switch sb_node.Mnemonic {
	case "beq":
		ins.AsmFn = encode.Beq
	default:
		a.Error("InsError", "No known asm fn", n)
	}

	a.curPC += 4
	a.curCu.Code = append(a.curCu.Code, ins)
}

func (a *Assembler) RType(n ast.Node) {
	r_node := n.(*ast.R_Type)
	ins := NewInstruction()

	ins.Regs = []uint32{uint32(r_node.Rd.Value), uint32(r_node.Rs1.Value), uint32(r_node.Rs2.Value)}

	switch r_node.Mnemonic {
	case "or":
		ins.AsmFn = encode.Or
	case "xor":
		ins.AsmFn = encode.Xor
	case "and":
		ins.AsmFn = encode.And
	default:
		a.Error("InsError", "No known asm fn", n)
	}

	a.curPC += 4
	a.curCu.Code = append(a.curCu.Code, ins)
}

func (a *Assembler) Other(n ast.Node) {
	o_type := n.(*ast.Other_Type)
	ins := NewInstruction()
	switch o_type.Mnemonic {
	case "ebreak":
		ins.AsmFn = encode.Ebreak
	default:
		a.Error("InsError", "No known asm fn", n)
	}

	a.curPC += 4
	a.curCu.Code = append(a.curCu.Code, ins)
}
