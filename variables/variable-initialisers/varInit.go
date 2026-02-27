package main

// Same as 'variable.go' but with values initialised.
func variableInit(d int, a bool, v bool, i bool) (int, bool, bool, bool) {
	return d, a, v, i
}

func main() {
	d := 1
	a, v, i := false, false, false
	variableInit(d, a, v, i)
}
