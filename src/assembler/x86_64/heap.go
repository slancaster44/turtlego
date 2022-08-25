package x86_64

import (
	"syscall"
	"turtlego/src/assembler/backpatch"
	"turtlego/src/pcode"
)

const (
	MMAP_ANON_FLAG int = syscall.MAP_PRIVATE | syscall.MAP_ANONYMOUS
	MMAP_FD        int = -1
	MMAP_PROT      int = syscall.PROT_EXEC | syscall.PROT_READ | syscall.PROT_WRITE
	MMAP_ADDR      int = 0
	MMAP_OFFSET    int = 0x00

	PAGE_SIZE     int = 4096
	PAGE_VAR_SIZE     = 8 //(one qword)
)

// Warning, you must take proper precautions to preserve
// the contents of the registers before using this function
func codeForMmap(length int) []byte {
	code := []byte{}

	//Load syscall arguments
	set_id_rax := codeToLoadImmDirectly(RAX, 0x09)
	code = append(code, set_id_rax...)

	set_addr_rdi := codeToLoadImmDirectly(RDI, MMAP_ADDR)
	code = append(code, set_addr_rdi...)

	set_len_rsi := codeToLoadImmDirectly(RSI, length)
	code = append(code, set_len_rsi...)

	set_prot_rdx := codeToLoadImmDirectly(RDX, MMAP_PROT)
	code = append(code, set_prot_rdx...)

	set_flags_r10 := codeMovExtendedRegImm(R10, MMAP_ANON_FLAG)
	code = append(code, set_flags_r10...)

	set_fd_r8 := codeMovExtendedRegImm(R8, MMAP_FD)
	code = append(code, set_fd_r8...)

	set_offset_r9 := codeMovExtendedRegImm(R9, MMAP_OFFSET)
	code = append(code, set_offset_r9...)

	//make syscall
	code = append(code, syscall_code...)

	return code
}

/* The structure of the heap
 * qword1: is allocated?
 * qword2: size of chunk in qwords
 * qword3...end: Space for use
 * And so on and so forth
 */

func MkHeap(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

	code = codeForMmap(PAGE_SIZE)

	//Save the address to the heap pointer register
	code = append(code, codeMovExtendedRegNormalReg(HP, RAX)...)

	//Set Heap Size register to 512 qwords
	code = append(code, codeMovExtendedRegImm(HS, 0x200)...)

	//Mmap already zeros out the entire page, so the allocation flag is
	//0x00 (unallocated) but we need to mark the size of the chunk in qwords
	loc_of_size_flag_rax, _, _ := genAuxInstruction(AddImmReg, pcode.REG1, PAGE_VAR_SIZE)
	code = append(code, loc_of_size_flag_rax...)

	mov_512_to_reg2, _, _ := genAuxInstruction(MovRegImm, pcode.REG2, 512)
	code = append(code, mov_512_to_reg2...)

	//TODO: implement mov imm to mem64
	mov_size_to_size_flag, _, _ := genAuxInstruction(MovRegInAddrFromReg, pcode.REG1, pcode.REG2)
	code = append(code, mov_size_to_size_flag...)

	return code, data, patches
}

// alloc <size-in-qwords-placed-in-register>
func Alloc(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}
	code = append(code, codeToPushPrimaryRegisters()...)

	//reg_with_size_in_qwords := ins.Arguments[0]
	//mov rax, reg_with_size

	//Search heap for gap closest to that size
	//Load start of data segment

	//Loop over chunks until one is found that is large enough
	//Save it's location in rbx, and size in rcx
	//If another chunk size is smaller than rcx, save it

	//If none are large enough, allocate enough space at the end of the heap
	//Cut off any excess space, if there is any, and mark it as available
	//Return the starting address the object can live in

	code = append(code, codeToPopPrimaryRegisters()...)
	return code, data, patches
}

func Dealloc(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

	//Mark the space in memory the object was using as available for use

	return code, data, patches
}

func Garbage(ins pcode.Instruction) ([]byte, []byte, []backpatch.BackPatch) {
	code, data, patches := []byte{}, []byte{}, []backpatch.BackPatch{}

	//Find any empty pages and munmap() them
	//Merge adjacent empty chunks

	return code, data, patches
}
