package ast

const (
	INT PrimativeType = iota
	FLT
	STR
	BOOL
	CHR

	NO_TYPE
)

type TypeInfo interface {
	DoesThisMatch(TypeInfo) bool
}

type PrimativeType byte

func (p PrimativeType) DoesThisMatch(t TypeInfo) bool {
	switch t.(type) {
	case PrimativeType:
		return t == p
	default:
		return false
	}
}

type ListType struct {
	TypeOfElements TypeInfo
}

func (l ListType) DoesThisMatch(t TypeInfo) bool {
	switch t := t.(type) {
	case ListType:
		return t.TypeOfElements.DoesThisMatch(l.TypeOfElements)
	default:
		return false
	}
}
