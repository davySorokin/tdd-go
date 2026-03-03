package main

import "fmt"

func whileRun(x int) (counter int) {
	counter = 0
	num := x
	for num <= 10 {
		fmt.Printf("While run: %d\n", num)
		num++
		counter++
	}
	return
}

func main() {
	whileRun(1)
}
