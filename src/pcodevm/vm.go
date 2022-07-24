package pcodevm

type VM struct {
	R1 int
	R2 int
	R3 int
	R4 int

	Sp int
	Bs int

	Memory []byte
}

//Instructions
//LD <register> <integer>
//LD <register> <address>
//PUSH <register>
//POP <register>
//ADD
//SUB
