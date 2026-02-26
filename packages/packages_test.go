package main

import (
	"testing"
)

// Maunally testing (clicking)
/*func TestInt(t *testing.T) {
	want := 5
	got := randNum()

	// Use 'go test -v' <- To see if it passes.
	if want == got {
		fmt.Printf("Yay! Both Want and Got are the Same %d", got)
	}
	if want != got {
		t.Errorf("\nWant and Got are not the Same: 'Want': %d, 'Got': %d", got, want)
	}
}*/

// Requires 'go test -v'
// Testing random ints between 0-10 to get 5.
func TestInt(t *testing.T) {
	want := 5
	failScore := 0
	passScore := 0

	for i := 1; i <= 30; i++ {
		got := randNum()
		if got != want {
			failScore++
			//t.Logf("FAILED - 'Want': %d, 'Got: %d", want, got)
		} else {
			passScore++
			//t.Logf("PASSED - Both 'Want' and 'Got' are %d", got)
		}
	}
	t.Logf("\nPassed: %d\nFailed: %d", passScore, failScore)

	if passScore == 0 {
		t.Errorf("rand() never returned %d in 30 attempts", want)
	}
}
