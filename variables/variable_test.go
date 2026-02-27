package main

import "testing"

func TestVariable(t *testing.T) {
	var o int
	var one, two, three bool
	wantedA, wantedB, wantedC, wantedD := 0, false, false, false
	gotA, gotB, gotC, gotD := checkingVars(o, one, two, three)

	if gotA != wantedA || gotB != wantedB || gotC != wantedC || gotD != wantedD {
		t.Errorf("\nAll the 'WANT'ed Values: %d %t %t %t.\nBut we 'GOT': %d, %t, %t, %t",
			wantedA, wantedB, wantedC, wantedD, gotA, gotB, gotC, gotD)
	} else {
		t.Logf("\nAll 'WANT's and 'GOT's are equal to:\n%d %t %t %t", wantedA, wantedB, wantedC, wantedD)
	}
}
