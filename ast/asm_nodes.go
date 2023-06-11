package ast

import "turtlego/tokens"

type Label struct {
	Name string
	id   int
	Tok  tokens.Token
}

func (l *Label) Id() int {
	return l.id
}

func (l *Label) SetId(i int) {
	l.id = i
}

func (l *Label) Type() int {
	return LABEL_NT
}

func (l *Label) Token() tokens.Token {
	return l.Tok
}

type DotLabel struct {
	Name string
	id   int
	Tok  tokens.Token
}

func (l *DotLabel) Id() int {
	return l.id
}

func (l *DotLabel) SetId(i int) {
	l.id = i
}

func (l *DotLabel) Type() int {
	return DOT_LABEL_NT
}

func (l *DotLabel) Token() tokens.Token {
	return l.Tok
}

type Directive struct {
	Name    string
	Options []Node
	id      int
	Tok     tokens.Token
}

func (l *Directive) Id() int {
	return l.id
}

func (l *Directive) Type() int {
	return DIRECTIVE_NT
}

func (l *Directive) SetId(i int) {
	l.id = i
}

func (l *Directive) Token() tokens.Token {
	return l.Tok
}

type Integer struct {
	Value int
	id    int
	Tok   tokens.Token
}

func (l *Integer) Id() int {
	return l.id
}

func (l *Integer) Type() int {
	return INTEGER_NT
}

func (l *Integer) SetId(i int) {
	l.id = i
}

func (l *Integer) Token() tokens.Token {
	return l.Tok
}

type String struct {
	Value string
	id    int
	Tok   tokens.Token
}

func (l *String) Type() int {
	return STRING_NT
}

func (l *String) SetId(i int) {
	l.id = i
}

func (l *String) Id() int {
	return l.id
}

func (l *String) Token() tokens.Token {
	return l.Tok
}

type Register struct {
	Value int
	id    int
	Tok   tokens.Token
}

func (l *Register) Id() int {
	return l.id
}

func (l *Register) SetId(i int) {
	l.id = i
}

func (l *Register) Type() int {
	return REGISTER_NT
}

func (l *Register) Token() tokens.Token {
	return l.Tok
}

type R_Type struct {
	Mnemonic string
	Rd       *Register
	Rs1      *Register
	Rs2      *Register
	id       int
	Tok      tokens.Token
}

func (l *R_Type) Id() int {
	return l.id
}

func (l *R_Type) SetId(i int) {
	l.id = i
}

func (l *R_Type) Type() int {
	return R_TYPE_NT
}

func (l *R_Type) Token() tokens.Token {
	return l.Tok
}

type I_Type struct {
	Mnemonic string
	Rd       *Register
	Rs1      *Register
	Operator string
	Imm      Node
	id       int
	Tok      tokens.Token
}

func (l *I_Type) Id() int {
	return l.id
}

func (l *I_Type) SetId(i int) {
	l.id = i
}

func (l *I_Type) Type() int {
	return I_TYPE_NT
}

func (l *I_Type) Token() tokens.Token {
	return l.Tok
}

type SB_Type struct {
	Mnemonic string
	Rs1      *Register
	Rs2      *Register
	Operator string
	Imm      Node
	id       int
	Tok      tokens.Token
}

func (l *SB_Type) Id() int {
	return l.id
}

func (l *SB_Type) SetId(i int) {
	l.id = i
}

func (l *SB_Type) Type() int {
	return SB_TYPE_NT
}

func (l *SB_Type) Token() tokens.Token {
	return l.Tok
}

type UJ_Type struct {
	Mnemonic string
	Rd       *Register
	Imm      Node
	id       int
	Tok      tokens.Token
}

func (l *UJ_Type) Id() int {
	return l.id
}

func (l *UJ_Type) SetId(i int) {
	l.id = i
}

func (l *UJ_Type) Type() int {
	return UJ_TYPE_NT
}

func (l *UJ_Type) Token() tokens.Token {
	return l.Tok
}

type Other_Type struct {
	Mnemonic string
	Regs     []*Register
	Imms     []Node
	id       int
	Tok      tokens.Token
}

func (l *Other_Type) Id() int {
	return l.id
}

func (l *Other_Type) SetId(i int) {
	l.id = i
}

func (l *Other_Type) Type() int {
	return OTHER_TYPE_NT
}

func (l *Other_Type) Token() tokens.Token {
	return l.Tok
}
