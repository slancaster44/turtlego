package encode

import (
	"fmt"
	"os"
	"runtime/debug"
)

func guard(regs []uint32, num_regs int, imms []uint32, num_imms int) {
	if len(regs) != num_regs {
		debug.PrintStack()
		fmt.Println("Error: Invalid number of registers")
		os.Exit(1)
	}

	if len(imms) != num_imms {
		debug.PrintStack()
		fmt.Println("Error: Invalid number of immediates")
		os.Exit(1)
	}
}
