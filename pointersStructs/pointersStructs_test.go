package main

import (
	"fmt"
	"testing"
)

func TestPointStruct(t *testing.T) {
	want := "x: 5, y: 1000000000"
	p := pointStructs()
	got := fmt.Sprintf("x: %d, y: %d", p.x, p.y)

	if want != got {
		t.Errorf("\nNO MATCH!\nWANT: %v\nGOT: %v", want, got)
	} else {
		t.Logf("MATCH FOUND!\nWANT: %v\nGOT: %v", want, got)
	}
}
