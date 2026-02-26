package main

import "testing"

/*
Inside naked():
x = total + 10
y = x * 2
*/

// go test -v
func TestNaked(t *testing.T) {
	wantA, wantB := 30, 60
	gotA, gotB := naked(20)

	if wantA != gotA || wantB != gotB {
		t.Errorf("\n'WANT: %d & %d, 'GOT': %d & %d", wantA, wantB, gotA, gotB)
	} else {
		t.Logf("\n'WANT' equals %d & %d  and 'GOT' equals %d & %d", wantA, wantB, gotA, gotB)
	}
}
