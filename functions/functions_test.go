package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := addition(5, 1)
	want := 6

	if got != want {
		t.Errorf("wanted %d but got %d", want, got)
	}
}
