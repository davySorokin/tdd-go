package main

type structLit struct {
	x, y int
}

var (
	x = structLit{5, 10}  //has type structLit
	y = structLit{x: 5}   //assigning only 'x' explicitly
	c = structLit{}       //x:0 & y:0 <- empty
	v = &structLit{5, 10} //has type *structLit
)

func structL() structLit {
	return x
}

func main() {
	structL()
}
