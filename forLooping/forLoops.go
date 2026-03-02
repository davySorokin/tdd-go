package main

import "fmt"

func looping() int {
	num1 := 1
	counter := 0

	for i := num1; i <= 10; i++ {
		fmt.Printf("Looped: %d Times.\n", i)
		counter++
	}
	fmt.Printf("\nLooped: %d Amount of Times.", counter)
	return counter
}

func main() {
	looping()
}
