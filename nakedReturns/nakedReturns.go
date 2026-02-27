package main

import "fmt"

// Input 'total',
// Returns both 'x' and 'y' without stating it in the return.
func naked(total int) (x, y int) {
	x = total + 10
	y = x * 2
	return // Will return x & y here.
}

func main() {
	fmt.Println(naked(20))
}
