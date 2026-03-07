package main

import "testing"

func TestPointers(t *testing.T) {
	want, wantA := pointers()

	/*if want != 100 || wantA != 200 {
		t.Errorf("expected 100 and 200, got %d and %d", want, wantA)
	} else {
		t.Logf("\nPASS\nFound Address of 'a' & 'b' -> %d & %d", want, wantA)
	}*/

	if want != 100 {
		t.Errorf("FAIL: expected 100, got %d", want)
	} else {
		t.Logf("PASS: got expected value 100")
	}

	if wantA != 200 {
		t.Errorf("FAIL: expected 200, got %d", wantA)
	} else {
		t.Logf("PASS: got expected value 200")
	}
}
