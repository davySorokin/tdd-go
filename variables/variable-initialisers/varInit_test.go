package main

import "testing"

// Same as 'variable_test.go' but with values initialised.
func TestVarInit(t *testing.T) {
	wantA := 0
	wantB, wantC, wantD := false, false, false
	gotA, gotB, gotC, gotD := variableInit(wantA, wantB, wantC, wantD)

	if wantA != gotA || wantB != gotB || wantC != gotC || wantD != gotD {
		t.Errorf("\n'WANT's and 'GOT's don't match!\nWANT: %d, %t, %t, %t\nGOT: %d, %t, %t, %t", wantA, wantB, wantC, wantD, gotA, gotB, gotC, gotD)
	} else {
		t.Logf("\nBOTH WANT & GOT MATCH!\n%d, %t, %t, %t", wantA, wantB, wantC, wantD)
	}
}
