package pcode

type Program struct {
	Instructions []*Instruction
	Data         []byte
}

func (b *Program) WriteInstruction(v *Instruction) {
	b.Instructions = append(b.Instructions, v)
}

func (b *Program) WriteData(v byte) {
	b.Data = append(b.Data, v)
}
