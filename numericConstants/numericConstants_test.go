package main

import "testing"

func TestIntNumConst(t *testing.T) {

	want := 21
	got := needInt(2) // Since func needInt performs the calculation.
	//got2 := int(needFloat(float64(needInt(want)))) <- confirming different test.

	if want != got {
		t.Errorf("\nNo Match - NEED: %d\nGOT: %d", want, got)
	} else {
		t.Logf("\nMatch Found!\nWANT: %d\nGOT: %d", want, got)
	}
}

// Identical test to TestIntNumConst but for float64
func TestFloatNumConst(t *testing.T) {
	want := 0.2
	got := needFloat(2) // needFloat like needInt, performs the calculation

	if want != got {
		t.Errorf("\nNo Match - NEED: %2f\nGOT: %2f", want, got)
	} else {
		t.Logf("\nMatch Found!\nWANT: %2f\nGOT: %2f", want, got)
	}
}
