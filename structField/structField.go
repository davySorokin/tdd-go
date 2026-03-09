package main

type testOne struct {
	x, y int
}

func sFields() int {
	v := testOne{2, 6}
	v.x = 8
	return v.x
}

func main() {
	sFields()
}
