package main

import (
	"fmt"
	"testing"
)

func TestStructL(t *testing.T) {
	want := "x: 5 & y: 10"
	p := structL()
	got := fmt.Sprintf("x: %d & y: %d", p.x, p.y)

	if want != got {
		t.Errorf("\nNo Match!\nWANT: %v\nGOT: %v", want, got)
	} else {
		t.Logf("\nMatch Found!\nWANT: %v\nGOT: %v", want, got)
	}
}
