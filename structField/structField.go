package main

type testOne struct {
	x, y int
}

func sFields() testOne {
	v := testOne{2, 6}
	v.x = 8
	return v
}

func main() {
	sFields()
}
