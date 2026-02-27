package main

import (
	"math/cmplx"
)

var (
	toBe   bool       = false
	maxInt uint64     = 1<<64 - 1
	cmplxZ complex128 = cmplx.Sqrt(-5 + 12i)
)

func basicTyping(toBe bool, maxInt uint64, cmplxZ complex128) (bool, uint64, complex128) {
	/*fmt.Printf("\ntoBe: %t", toBe)
	fmt.Printf("\nmaxInt: %d", maxInt)
	fmt.Printf("\nz: %v", cmplxZ) short version
	fmt.Printf("\nz: %2fi", cmplxZ)*/
	/*if toBe == false {
		fmt.Printf("\ntoBe: %t", toBe)
	} else {
		fmt.Println("\ntoBe doesn't match!")
	}
	if maxInt == 1<<64-1 {
		fmt.Printf("\nmaxInt: %d", maxInt)
	} else {
		fmt.Printf("\nmaxInt doesn't match!")
	}
	if cmplxZ == cmplx.Sqrt(-5+12i) {
		fmt.Printf("\ncmplxZ: %2fi\n", cmplxZ)
	} else {
		fmt.Println("cmplxZ doesn't match!")
	}*/

	return false, 1<<64 - 1, -5 + 12i
}

func main() {
	basicTyping(toBe, maxInt, cmplxZ)
}
