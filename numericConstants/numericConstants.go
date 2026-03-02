package main

import "fmt"

const (
	// no work - overflows
	// big = 1 << 100

	// works
	// Create a Hugh number by shifting 100 bits to the left.
	big = 1 << 100
	// Shift again the 'big' to the right 99 times <- ending up with 1<<1 or 2.
	small = big >> 99
)

func showConst() {
	fmt.Println(small)
}

// Result: ((2*10)+1) = 21
func needInt(x int) int { return x*10 + 1 }

// Result: (2*0.1) = 0.2
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	//showConst()
	needInt(small)
	needFloat(small)
}
