package main

import "fmt"

func subtraction(x, y int) int {
	z := x - y
	fmt.Printf("Remaining value from %d and %d is %d", x, y, z)
	return z
}

func main() {
	subtraction(10, 5)
}
