package main

// Letting variable be set to default 0s and false(s).
func checkingVars(o int, one, two, three bool) (int, bool, bool, bool) {
	//fmt.Printf("%d, %t, %t, %t", o, one, two, three)
	return o, one, two, three
}

func main() {
	var a int
	var b, c, d bool
	checkingVars(a, b, c, d)
}
