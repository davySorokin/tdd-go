package main

import "testing"

// go test -v

// Check amount of times looped test.
func TestLooping(t *testing.T) {
	wanted := 10
	got := looping()

	if wanted != got {
		t.Errorf("\nIncorrect Looped Amount: %d\nNeeded: %d\n", got, wanted)
	} else {
		t.Logf("\nCorrect Amount Looped: %d", wanted)
	}
}
