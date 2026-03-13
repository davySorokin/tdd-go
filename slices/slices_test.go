package main

import (
	"reflect"
	"testing"
)

func TestSlices(t *testing.T) {
	want := [6]int{2, 3, 5, 7, 11, 13}
	var wanted []int = want[1:4]
	got, got1 := slices()

	if !reflect.DeepEqual(wanted, got) || !reflect.DeepEqual(wanted, got1) {
		t.Errorf("\nNo Match!\nWanted: %v\nGot - 1: %v\nGot - 2: %v", wanted, got, got1)
	} else {
		t.Logf("\nMatch Found!\nWanted: %v\nGot - 1: %v\nGot - 2: %v", wanted, got, got1)
	}
}
