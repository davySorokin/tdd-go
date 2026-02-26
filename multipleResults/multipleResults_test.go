package main

import (
	"testing"
)

// go test -v
func TestSwap(t *testing.T) {
	wantA, wantB := "David", "Hello"

	// gotA -> "David" & gotB -> "Hello"
	gotA, gotB := swapping("Hello", "David")

	if gotA != wantA || gotB != wantB {
		t.Errorf("\nWANT doesn't equal GOT\nIt is still '%s %s' when it should be 'David Hello'.", wantB, wantA)
	} else {
		t.Logf("\nBoth WANT and GOT are SWAPPED = %s %s", wantB, wantA)
	}
}
