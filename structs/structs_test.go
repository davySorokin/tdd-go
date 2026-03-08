package main

import "testing"

func TestStruct(t *testing.T) {
	a := vertex{1, 2}
	b := struEx()

	if a != b {
		t.Errorf("\nStruct Values Don't Match!\nWANT: %v\nGOT: %v", a, b)
	} else {
		t.Logf("\nBoth Structs Match!\nWANT: %v\nGOT: %v", a, b)
	}
}
