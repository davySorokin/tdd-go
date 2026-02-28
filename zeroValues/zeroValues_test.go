package main

import "testing"

func TestZeroValues(t *testing.T) {
	var (
		gotA int
		gotB bool
		gotC float64
		gotD string
	)
	wantA, wantB, wantC, wantD := valZero(gotA, gotB, gotC, gotD) // PASSES
	// wantA, wantB, wantC, wantD := valZeroTwo(gotA, gotB, gotC, gotD) // FAILS

	if gotA != wantA || gotB != wantB || gotC != wantC || gotD != wantD {
		t.Errorf("\nNo Matches!")
	} else {
		t.Logf("\nMatches Confirmed")
	}
}
