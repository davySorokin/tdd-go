package main

import (
	"math"
	"testing"
)

func TestConversion(t *testing.T) {
	wantedA, wantedB := 3, 4
	wantedC := math.Sqrt(float64(wantedA*wantedA + wantedB*wantedB))
	wantedC_two := math.Sqrt(float64((wantedA*wantedA + wantedB*wantedB) - 1))
	wantedD := uint64(wantedC)
	gotA, gotB, gotC, gotD := converting(wantedA, wantedB)

	// PASSES
	if wantedA != gotA || wantedB != gotB || wantedC != gotC || wantedD != gotD {
		t.Errorf("\nNO MATCHES!")
	} else {
		t.Logf("\nALL MATCHES FOUND!")
	}

	// FAILS
	if wantedA != gotA || wantedB != gotB || wantedC_two != gotC || wantedD != gotD {
		t.Errorf("\nNO MATCHES!")
	} else {
		t.Logf("\nALL MATCHES FOUND!")
	}

}
