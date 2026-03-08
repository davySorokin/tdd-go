package main

import "fmt"

type vertex struct {
	x, y int
}

func struEx() vertex {
	return vertex{1, 2}
}

func main() {
	fmt.Println(struEx())
}
