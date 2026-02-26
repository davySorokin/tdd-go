package main

import (
	"fmt"
	"math/rand"
)

func randNum() int {
	randomising := rand.Intn(10)
	fmt.Println("My favourite number is: ", randomising)
	return randomising
}
func main() {
	randNum()
}
