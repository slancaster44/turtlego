package pcodegenerator

import (
	"turtlego/src/ast"
	"turtlego/src/message"
	"turtlego/src/pcode"
	"turtlego/src/tokens"
)

const STACK_VAR_SIZE int = 8
const SIZE_OF_STACK_METADATA int = (STACK_VAR_SIZE * 2)
const UNKOWN_ADDRESS int = 0

type OpTypePair struct {
	Op  string
	Typ ast.TypeInfo
}

type Generator struct {
	SyntaxTree                ast.Block
	CodeGenFnMap              map[byte]func(ast.Node) Register
	infixGenFnMap             map[OpTypePair]func(ast.Node) Register
	builtinsGenMap            map[string]func(*ast.Builtin) Register
	numberOfActiveAllocations map[int]int //Maps a given register to the number of times it has been allocated
	numRegisterPushesToStack  int         //Number of total saved register values on stack
	Program                   pcode.Program
	NumberOfRegisters         int
	SymTab                    map[string]ast.Location //Maps identifiers to their location on the stack
}

func NewGenerator(st ast.Block, registerCountInTargetMachine int) *Generator {
	ret_val := &Generator{}
	ret_val.SyntaxTree = st
	ret_val.numRegisterPushesToStack = 0
	ret_val.NumberOfRegisters = registerCountInTargetMachine

	ret_val.CodeGenFnMap = map[byte]func(ast.Node) Register{
		ast.INT_NT:     ret_val.genIntCode,
		ast.INFIX_NT:   ret_val.genInfixCode,
		ast.LETINIT_NT: ret_val.genLetInit,
		ast.IDENT_NT:   ret_val.genIdentCode,
		ast.BUILTIN_NT: ret_val.genBuiltinCode,
		ast.BLOCK_NT:   ret_val.genBlockCode,
		ast.BOOLEAN_NT: ret_val.genBoolCode,
		ast.IFEL_NT:    ret_val.genIfElse,
		ast.CHR_NT:     ret_val.genIntCode,
		ast.LIST_NT:    ret_val.genListInit,
	}

	ret_val.numberOfActiveAllocations = make(map[int]int)
	for i := 0; i < registerCountInTargetMachine; i++ {
		ret_val.numberOfActiveAllocations[i] = 0
	}

	ret_val.infixGenFnMap = map[OpTypePair]func(ast.Node) Register{
		{"+", ast.INT}:   ret_val.mkInfixOpGenFunc(pcode.ADD_REG_INT_INT, pcode.ADD_REG_REG_INT),
		{"-", ast.INT}:   ret_val.mkInfixOpGenFunc(pcode.SUB_REG_INT_INT, pcode.SUB_REG_REG_INT),
		{"*", ast.INT}:   ret_val.mkInfixOpGenFunc(pcode.MUL_REG_INT_INT, pcode.MUL_REG_REG_INT),
		{"/", ast.INT}:   ret_val.mkInfixOpGenFunc(pcode.DIV_REG_INT_INT, pcode.DIV_REG_REG_INT),
		{"||", ast.BOOL}: ret_val.mkInfixOpGenFunc(pcode.BOOL_OR_REG_IMM, pcode.BOOL_OR_REG_REG),
		{"&&", ast.BOOL}: ret_val.mkInfixOpGenFunc(pcode.BOOL_AND_REG_IMM, pcode.BOOL_AND_REG_REG),
		{"==", ast.BOOL}: ret_val.mkInfixOpGenFunc(pcode.EQ_REG_IMM, pcode.EQ_REG_REG),
		{"!=", ast.BOOL}: ret_val.mkInfixOpGenFunc(pcode.NE_REG_IMM, pcode.NE_REG_REG),
		{"<", ast.BOOL}:  ret_val.mkInfixOpGenFunc(pcode.LT_REG_IMM, pcode.LT_REG_REG),
		{"<=", ast.BOOL}: ret_val.mkInfixOpGenFunc(pcode.LE_REG_IMM, pcode.LE_REG_REG),
		{">", ast.BOOL}:  ret_val.mkInfixOpGenFunc(pcode.GT_REG_IMM, pcode.GT_REG_REG),
		{">=", ast.BOOL}: ret_val.mkInfixOpGenFunc(pcode.GE_REG_IMM, pcode.GE_REG_REG),
	}

	ret_val.builtinsGenMap = map[string]func(*ast.Builtin) Register{
		"print": ret_val.genPrintCode,
	}

	ret_val.SymTab = make(map[string]ast.Location)

	return ret_val
}

func (g *Generator) raiseError(n, m string, tok tokens.Token) {
	message.Error(tok.Filename, n, m, tok.LineNo, tok.ColumnNo)
}

func (g *Generator) GenPCode() {
	g.WriteInstruction(pcode.MK_HEAP)
	g.pushStackFrame(g.SyntaxTree.NumStackVars, g.SyntaxTree.ScopeDepth)
	for _, stmt := range g.SyntaxTree.Exprs {
		reg := g.appendCodeFor(stmt)
		g.ReleaseRegister(reg)
	}

	if g.numRegisterPushesToStack != 0 {
		g.raiseError("Generatation",
			"This is a bug, not all registers where properly deallocated during code generation",
			g.SyntaxTree.Tok)
	}
}

func (g *Generator) appendCodeFor(stmt ast.Node) Register {
	fn, ok := g.CodeGenFnMap[stmt.NodeType()]

	if !ok {
		g.raiseError("Generation", "This is a bug, unimplemented node type", stmt.GetTok())
	}

	return fn(stmt)
}

// Writes an instruction to code, returns a pointer to that instruction
// as well as the location of that instruction in code
func (g *Generator) WriteInstruction(opcode byte, args ...int) (*pcode.Instruction, int) {
	locOfIns := len(g.Program.Instructions)

	ins := pcode.MkInstruction(opcode, args...)
	g.Program.WriteInstruction(ins)

	return ins, locOfIns
}
