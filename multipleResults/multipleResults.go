package main

import "fmt"

func swapping(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swapping("Hello", "David")
	fmt.Println(a, b)
}
