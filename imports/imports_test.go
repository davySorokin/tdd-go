package main

import (
	"fmt"
	"testing"
)

func TestSqrt(t *testing.T) {
	got := squareValue()
	want := 2.6457513110645907

	if got != want {
		t.Errorf("Wanted: %g, Got: %g", got, want)
	}
	// run using go test -v
	if got == want {
		fmt.Printf("Both 'Want' & 'Got' are: %g", got)
	}
}
