package main

// Just a simple zero value - no explicit initial value with naked/bare return.
// Silly but fun test.
func valZero(a int, b bool, c float64, d string) (aa int, bb bool, cc float64, dd string) {
	return
}

func valZeroTwo(a int, b bool, c float64, d string) (aa int, bb bool, cc float64, dd string) {
	return 0, false, 1.0, "hello"
}

func main() {
	var (
		a int
		b bool
		c float64
		d string
	)

	valZero(a, b, c, d)
	valZeroTwo(a, b, c, d)
}
