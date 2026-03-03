package main

import (
	"fmt"
	"math"
)

func ifShorts(x, y, lim float64) (v float64) {
	if v = math.Pow(x, y); v < lim {
		fmt.Printf("%.4f", v)
	}
	return
}

func main() {
	ifShorts(2, 3, 10)
}
