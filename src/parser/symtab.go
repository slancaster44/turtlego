package parser

import "turtlego/src/ast"

type TableEntry struct {
	LocationOnStack int
	Type            ast.TypeInfo
}

type SymTab struct {
	CurStackOffset int
	Entries        map[string]TableEntry
}

func NewSymtab() SymTab {
	return SymTab{0, map[string]TableEntry{}}
}
