package main

import "fmt"

func loopCon(x int) (counter int) {
	sum := x
	counter = 0

	for ; sum < 100; sum++ {
		fmt.Printf("Round #%d of Loop.\n", sum)
		counter++
	}
	return
}

func main() {
	loopCon(10)
}
