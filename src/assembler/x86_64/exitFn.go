package x86_64

func ExitX86() []byte {
	//Three instructions must be written to the
	//code in order for the program to exit properly with
	//an exit code of 0
	outCode := []byte{}

	mov_rax_0x01 := []byte{0x48, 0xC7, 0xC0, 0x01, 0x00, 0x00, 0x00} //Set rax to exit interrupt
	mov_rbx_0x00 := []byte{0x48, 0xC7, 0xC3, 0x00, 0x00, 0x00, 0x00} //Set exit code to 0
	int_0x80 := []byte{0xCD, 0x80}                                   //Call interrupt set in rax

	outCode = append(outCode, mov_rax_0x01...)
	outCode = append(outCode, mov_rbx_0x00...)
	outCode = append(outCode, int_0x80...)

	return outCode
}
