package main

import "testing"

func TestArrays(t *testing.T) {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	want := primes
	got := arraying()

	if want != got {
		t.Errorf("\nNo Match!\nWant: %v\nGot: %v", want, got)
	} else {
		t.Logf("\nMatch Found!\nWant: %v\nGot: %v", want, got)
	}
}
