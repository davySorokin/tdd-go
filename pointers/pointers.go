package main

func pointers() (int, int) {
	a, b := 100, 200

	z := &a
	//fmt.Println(z)  // show address of 'a'
	//fmt.Println(*z) // show value at address of 'a'
	x := &b
	//fmt.Println(x)  // address of 'b'
	//fmt.Println(*x) // value at address of 'b'

	return *z, *x
}

func main() {
	pointers()
}
