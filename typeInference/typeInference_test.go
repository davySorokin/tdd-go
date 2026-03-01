package main

import (
	"fmt"
	"testing"
)

func TestIntTypeInf(t *testing.T) {

	// go test -v (for logs)

	// Had to make wanted interface{} - or switch case would consistently show errors.
	var wanted interface{}
	// Can change to whatever type like in the typeInference.go
	wanted = "42"
	got := typeInf()

	switch v := got.(type) {
	case int:
		if v != wanted {
			t.Errorf("\nWrong Type!\nNEEDED: %T\nGOT: %T", wanted, got)
		} else {
			t.Logf("\nMatch Found!\nNEEDED: %T\nGOT: %T", wanted, got)
		}
	case string:
		if fmt.Sprintf("%T", v) != fmt.Sprintf("%T", wanted) {
			t.Errorf("\nWrong Type!\nNEEDED: %T\nGOT: %T", wanted, got)
		} else {
			t.Logf("\nMatch Found!\nNEEDED: %T\nGOT: %T", wanted, got)
		}
	case float64:
		if v != wanted {
			t.Errorf("\nWrong Type!\nNEEDED: %T\nGOT: %T", wanted, got)
		} else {
			t.Logf("\nMatch Found!\nNEEDED: %T\nGOT:%T", wanted, got)
		}
	default:
		t.Errorf("\nWRONG TYPE FOUND: %T\nNEEDED: %T", got, wanted)
	}
}
