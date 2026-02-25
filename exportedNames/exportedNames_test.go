package main

import "testing"

func TestPi(t *testing.T) {
	got := pies()
	want := 3.141592653589793

	if got != want {
		t.Errorf("Needed %f, but Got %f", got, want)
	}
}
