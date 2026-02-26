package main

import "fmt"

// Input 'total',
// Returns both 'x' and 'y'.
func naked(total int) (x, y int) {
	x = total + 10
	y = x * 2
	return
}

func main() {
	fmt.Println(naked(20))
}
