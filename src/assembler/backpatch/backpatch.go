package backpatch

/* If an instruction needs to write an address
 * that is not yet known, then it will return a
 * backpatch. This will contain information about
 * where in the returned code to patch in the address
 * once it is known, as well is where in the generated pcode
 * the address is
 */

type BackPatch struct {
	LocationOfAddressToPatch     int
	LocationOfInstructionPatched int
	PcodeAddressToPatchTo        int
}
