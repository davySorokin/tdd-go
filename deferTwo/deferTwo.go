package main

import (
	"bytes"
	"fmt"
)

func deferTwo() string {
	var buf bytes.Buffer

	// redirecting fmt.Print to buffer.
	fmt.Fprint(&buf, "Hello ")
	defer fmt.Fprint(&buf, "David")

	// returning the captured string.
	return buf.String()
}

func main() {
	fmt.Println(deferTwo())
}
