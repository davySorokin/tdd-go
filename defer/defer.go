package main

import "fmt"

func defFunc() string {
	var x string
	fmt.Println("Input your name: ")
	_, err := fmt.Scanf("%s", &x)
	if err != nil {
		fmt.Println("Enter your name using alphabetical letters!\nNot '%s'", err)
		return ""
	}
	defer fmt.Println(x)
	fmt.Printf("Hello ")
	return x
}

func main() {
	defFunc()
}
