package assembler

import (
	"encoding/binary"
	"io/ioutil"
)

func (a *Assembler) WriteElf(filename string) {
	a.WriteExitToCode()
	a.BuildElfHeader()
	a.WriteOutput(a.code...)
	a.WriteOutput(a.data...)
	ioutil.WriteFile(filename, a.output, 0775)
}

const (
	SIZE_OF_ELF_HEADER  byte = 0x40
	SIZE_OF_PROG_HEADER byte = 0x38

	ELF_START_ADDR      uint64 = 0x400000
	ELF_DATA_START_ADDR uint64 = 0x600000
	ELF_ALIGNMENT       uint64 = 0x200000
)

/* Elf format:
 * ELF header
 * Program Header
 * Code
 * TODO: Refractor to use labeled constants
 */
func (a *Assembler) BuildElfHeader() {
	codeSize := len(a.code)
	codeOffset := SIZE_OF_ELF_HEADER + (2 * SIZE_OF_PROG_HEADER)

	//Write ELF magic number
	a.WriteOutput(0x7f, 0x45, 0x4c, 0x46)

	//Indicate that his is a 64-bit executable
	a.WriteOutput(0x02)

	//Indicate the bytes are little endian (b/c this is x86)
	a.WriteOutput(0x01)

	//Indicate ELF version
	a.WriteOutput(0x01)

	//Indicate Target OS ABI
	a.WriteOutput(0x00)

	//Indicate ABI Version
	a.WriteOutput(0x00)

	//Write Unused header bytes
	a.WriteOutput(0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00)

	//Write executable type and target architecture
	a.WriteOutput(0x02, 0x00)
	a.WriteOutput(0x3e, 0x00)

	//Elf version, again???
	a.WriteOutput(0x01, 0x00, 0x00, 0x00)

	//Write load offset
	a.WriteValOutput(8, ELF_START_ADDR+uint64(codeOffset))

	//Write offset from file start to program header
	a.WriteOutput(0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00)

	//Start section header table
	a.WriteOutput(0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00)

	//Write Flags
	a.WriteOutput(0x00, 0x00, 0x00, 0x00)

	//Size of this header
	a.WriteOutput(0x40, 0x00)

	//Size of program header table entry
	a.WriteOutput(0x38, 0x00)

	//Number of sections (just data & text)
	a.WriteOutput(0x02, 0x00)

	//Size of section headers (not used in this implementation)
	a.WriteOutput(0x00, 0x00)

	//Number of entries in section headers (again, not used)
	a.WriteOutput(0x00, 0x00)

	//Index of Section Header table Entry (again not used)
	a.WriteOutput(0x00, 0x00)

	//Build Program Header. One for text segment, one for data

	//Text
	//PT_LOAD, loadable segment
	a.WriteOutput(0x01, 0x00, 0x00, 0x00)

	//Flags
	a.WriteOutput(0x07, 0x00, 0x00, 0x00)

	//Text Offset
	a.WriteValOutput(8, 0)
	a.WriteValOutput(8, ELF_START_ADDR)

	//Pysical Address. Same as text offset on linux
	a.WriteValOutput(8, ELF_START_ADDR)

	//Size of text segment (when would they be different?)
	a.WriteValOutput(8, uint64(codeSize)) //Size in file
	a.WriteValOutput(8, uint64(codeSize)) //Size in memory

	a.WriteValOutput(8, ELF_ALIGNMENT)

	//And data section
	dataSize := len(a.data)
	dataOffset := int(codeOffset) + codeSize
	dataAddress := ELF_DATA_START_ADDR + uint64(dataOffset)

	a.WriteOutput(0x01, 0x00, 0x00, 0x00) //PT_LOAD
	a.WriteOutput(0x07, 0x00, 0x00, 0x00)

	//Offset of data
	a.WriteValOutput(8, uint64(dataOffset))

	//Address that data is loaded to
	a.WriteValOutput(8, dataAddress) //Virtual
	a.WriteValOutput(8, dataAddress) //Same for physicall addr

	//Size of section
	a.WriteValOutput(8, uint64(dataSize)) //Size in file
	a.WriteValOutput(8, uint64(dataSize)) //Size in memory

	a.WriteValOutput(8, ELF_ALIGNMENT)
}

func (a *Assembler) WriteOutput(b ...byte) {
	a.output = append(a.output, b...)
}

func (a *Assembler) WriteValOutput(size byte, val uint64) {
	buffr := make([]byte, size)
	binary.LittleEndian.PutUint64(buffr, val)
	a.WriteOutput(buffr...)
}

func (a *Assembler) WriteExitToCode() {
	//Three instructions must be written to the
	//code in order for the program to exit properly with
	//an exit code of 0

	mov_rax_0x01 := []byte{0x48, 0xC7, 0xC0, 0x01, 0x00, 0x00, 0x00} //Set rax to exit interrupt
	mov_rbx_0x00 := []byte{0x48, 0xC7, 0xC3, 0x01, 0x00, 0x00, 0x00} //Set exit code
	int_0x80 := []byte{0xCD, 0x80}                                   //Call interrupt set in rax

	a.code = append(a.code, mov_rax_0x01...)
	a.code = append(a.code, mov_rbx_0x00...)
	a.code = append(a.code, int_0x80...)
}
