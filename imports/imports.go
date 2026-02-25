package main

import (
	"fmt"
	"math"
)

func squareValue() float64 {
	result := math.Sqrt(7)
	fmt.Printf("Now you have %g problems.\n", result)
	return result
}

func main() {
	squareValue()
}
