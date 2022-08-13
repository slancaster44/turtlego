package parser

type TableEntry struct {
	LocationOnStack int
	Type            byte
}

type SymTab struct {
	CurStackOffset int
	Entries        map[string]TableEntry
}

func NewSymtab() SymTab {
	return SymTab{0, map[string]TableEntry{}}
}
