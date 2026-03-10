package main

type pstruct struct {
	x int
	y int
}

func pointStructs() pstruct {
	a := pstruct{5, 10}
	b := &a
	b.y = 1e9
	//fmt.Println(a)
	return a
}

func main() {
	pointStructs()
}
