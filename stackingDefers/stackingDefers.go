package main

import "fmt"

// utilising the LIFO approach with defer.
func stackDef() int {
	fmt.Println("counting...")

	a := make([]int, 0, 9)
	for i := 1; i <= 10; i++ {
		a = append(a, i)
		defer fmt.Print(i) // runs after the 'Done!' line.
		fmt.Print(a[i-1])  // print slice/array with 'Done!' before defer function executes.
	}
	fmt.Println() // adds new line between them.
	return len(a)
}

func main() {
	stackDef()
}
