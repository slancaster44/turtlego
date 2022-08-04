package parser

type TypeTable struct {
	CurStackOffset int
	Entries        map[string]byte
}

func NewTypeTable() TypeTable {
	return TypeTable{0, map[string]byte{}}
}
