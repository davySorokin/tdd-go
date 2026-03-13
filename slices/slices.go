package main

func slices() ([]int, []int) {
	primeNum := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primeNum[1:4] //cuts the 'primeNum' into low and high bound range(s).
	//fmt.Print(s)
	return primeNum[1:4], s
}

func main() {
	slices()
}
