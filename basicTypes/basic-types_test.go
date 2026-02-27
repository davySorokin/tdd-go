package main

import (
	"testing"
)

func TestBasicTypes(t *testing.T) {
	wantA := false
	var wantB uint64 = 1<<64 - 1
	wantC := -5 + 12i
	gotA, gotB, gotC := basicTyping(wantA, wantB, wantC)

	if wantA != gotA || wantB != gotB || wantC != gotC {
		t.Errorf("No Match!\n Wanted:\n%t\n%d\n%2f\nGot:\n%t\n%d\n%2f",
			wantA, wantB, wantC, gotA, gotB, gotC)
	} else {
		t.Logf("\nBoth Want & Got are the Same!\n%t\n%d\n%2fi", wantA, wantB, wantC)
	}

}
