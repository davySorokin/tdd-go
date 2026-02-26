package main

import (
	"testing"
)

func TestSub(t *testing.T) {
	want := 5
	got := subtraction(10, 5)

	if want != got {
		t.Errorf("\nFAILED: Value doesn't match what is needed.\nWANT: %d\nGOT: %d", want, got)
	} else {
		t.Logf("WANT and GOT values match %d!", want)
	}
}
