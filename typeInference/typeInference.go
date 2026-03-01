package main

import "fmt"

func typeInf() interface{} {
	v := "42" // change this to what type you need, change return after.
	fmt.Printf("Stored variable is type %T", v)
	return v
}

func main() {
	typeInf()
}
