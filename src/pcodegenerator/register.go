package pcodegenerator

import (
	"fmt"
)

/* A note on register allocation:
 * Whenever a code generation function is run, it
 * will leave its results in a register. However,
 * it will not release that register, incase any
 * following functions need to hold on to the value
 * in that register.
 * As a result, if you need a register used by a
 * previous function to retain its value, do nothing,
 * the value is already locked. However, if you need
 * its value to change, you must release that register
 */

type Register struct {
	RegisterNumber               int
	LocationOfOldContentsOnStack int
}

// TODO: Make private
func (g *Generator) GetRegister() Register {
	selectedRegister := 0
	numberActiveAllocationsOnThisReg := 5 //TODO: Make pretty (no magic numbers)

	for reg, numberOfTimesAllocated := range g.numberOfActiveAllocations {
		if numberOfTimesAllocated <= numberActiveAllocationsOnThisReg {
			selectedRegister = reg
			numberActiveAllocationsOnThisReg = numberOfTimesAllocated
		}
	}

	if numberActiveAllocationsOnThisReg > 0 {
		g.genPushRegToStack(selectedRegister)
		g.numRegisterPushesToStack++
	}

	g.numberOfActiveAllocations[selectedRegister]++
	return Register{selectedRegister, g.numRegisterPushesToStack}
}

func (g *Generator) lockRegister(r Register) {
	if r.LocationOfOldContentsOnStack > 0 {
		g.genPushRegToStack(r.RegisterNumber)
		g.numRegisterPushesToStack++
	}

	g.numberOfActiveAllocations[r.RegisterNumber]++
}

// TODO: Make private
func (g *Generator) ReleaseRegister(r Register) {
	if r.LocationOfOldContentsOnStack != g.numRegisterPushesToStack {
		fmt.Println(r.LocationOfOldContentsOnStack, g.numRegisterPushesToStack)
		panic("Cannot deallocate a register. This is a bug. Please report")
	}

	g.numberOfActiveAllocations[r.RegisterNumber]--

	if r.LocationOfOldContentsOnStack > 0 {
		g.numRegisterPushesToStack--
		g.genPopToReg(r.RegisterNumber)
	}

}
