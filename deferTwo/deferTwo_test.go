package main

import "testing"

func TestDefTwo(t *testing.T) {

	// technically the right output.
	want := "Hello "

	// Cannot output - defer outputs after return is executed.
	//want := "Hello David"
	got := deferTwo()

	if want != got {
		t.Errorf("No Match!\nWANT: %s\nGOT: %s\n", want, got)
	}
}
