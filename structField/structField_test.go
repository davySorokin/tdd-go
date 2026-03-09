package main

import "testing"

func TestStructField(t *testing.T) {
	want := 8
	got := sFields().x

	if want != got {
		t.Errorf("\nStruct Field Don't Match!\nWANT: %v\nGOT: %v", want, got)
	} else {
		t.Logf("Struct Field Match!\nWANT: %v\nGOT: %v", want, got)
	}
}
