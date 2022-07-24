package ast

import (
	"fmt"
	"turtlego/src/tokens"
)

type Int struct {
	Value int
	Tok   tokens.Token
}

func (s *Int) PrintAll(tab string) {
	fmt.Print(s.Stringify(tab))
}
func (s *Int) Stringify(tab string) string {
	return fmt.Sprintf("%s<Integer value='%d'>\n", tab, s.Value)
}

func (s *Int) GetTok() tokens.Token {
	return s.Tok
}

func (i *Int) TypeGenerated() byte {
	return INT
}

func (i *Int) NodeType() byte {
	return INT_NT
}

type Flt struct {
	Value float64
	Tok   tokens.Token
}

func (s *Flt) PrintAll(tab string) {
	fmt.Print(s.Stringify(tab))
}
func (s *Flt) Stringify(tab string) string {
	return fmt.Sprintf("%s<Floating Point value='%f'>\n", tab, s.Value)
}

func (s *Flt) GetTok() tokens.Token {
	return s.Tok
}

func (l *Flt) TypeGenerated() byte {
	return FLT
}

func (l *Flt) NodeType() byte {
	return FLOAT_NT
}

/////////////////////////////////////////////////

type String struct {
	Value string
	Tok   tokens.Token
}

func (s *String) PrintAll(tab string) {
	fmt.Print(s.Stringify(tab))
}
func (s *String) Stringify(tab string) string {
	return fmt.Sprintf("%s<String value='%s'>\n", tab, s.Value)
}

func (s *String) GetTok() tokens.Token {
	return s.Tok
}

func (s *String) TypeGenerated() byte {
	return STR
}

func (s *String) NodeType() byte {
	return STRING_NT
}

/////////////////////////////////////////////////

type Boolean struct {
	Value bool
	Tok   tokens.Token
}

func (b *Boolean) PrintAll(tab string) {
	fmt.Print(b.Stringify(tab))
}
func (b *Boolean) Stringify(tab string) string {
	return fmt.Sprintf("%s<Boolean value=%v>\n", tab, b.Value)
}

func (b *Boolean) GetTok() tokens.Token {
	return b.Tok
}

func (b *Boolean) TypeGenerated() byte {
	return BOOL
}

func (b *Boolean) NodeType() byte {
	return BOOLEAN_NT
}

/////////////////////////////////////////////////

type Identifier struct {
	Value           string
	Tok             tokens.Token
	ScopeDepthFound int
	Type            byte
}

func (i *Identifier) PrintAll(tab string) {
	fmt.Print(i.Stringify(tab))
}

func (i *Identifier) Stringify(tab string) string {
	return fmt.Sprintf("%s<Identifier value=%s, depth found=%d>\n", tab, i.Value, i.ScopeDepthFound)
}

func (i *Identifier) GetTok() tokens.Token {
	return i.Tok
}
func (i *Identifier) TypeGenerated() byte {
	return i.Type
}

func (i *Identifier) NodeType() byte {
	return IDENT_NT
}
