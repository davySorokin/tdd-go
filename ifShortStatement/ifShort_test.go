package main

import "testing"

func TestIfShort(t *testing.T) {
	wanted := 8.0000          // Must equal %.4f
	got := ifShorts(2, 3, 10) // math.Pow(x, y) then limit checking value against '10'.

	if wanted != got {
		t.Errorf("\nTotal Value Exceeds Limit!\nLimit: 10\nWANTED: %.4f\nGOT: %.4f", wanted, got)
	} else {
		t.Logf("\nTotal Value is within Range of Limit: 10\nWANTED: %.4f\nGOT: %.4f", wanted, got)
	}
}
