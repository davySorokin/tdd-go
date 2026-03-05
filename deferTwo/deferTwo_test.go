package main

import "testing"

func TestDefTwo(t *testing.T) {
	want := "Hello "
	got := deferTwo()

	if want != got {
		t.Errorf("No Match!\nWANT: %s\nGOT: %s\n", want, got)
	}
}
