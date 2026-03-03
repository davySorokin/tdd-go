package main

import "testing"

func TestWhile(t *testing.T) {
	wanted := 10
	got := whileRun(1)

	if wanted != got {
		t.Errorf("\nWanted While Runs - NO MATCH!\nWANTED: %d\nGOT: %d\n", wanted, got)
	} else {
		t.Logf("Wanted While Runs - MATCH!\nWANTED: %d\nGOT: %d\n", wanted, got)
	}
}
