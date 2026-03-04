package main

import (
	"testing"
)

func TestLoopingCon(t *testing.T) {
	want := 10
	got := loopCon(10)

	switch {
	case want <= 10:
		t.Logf("\nLoop Count Too Low! - %d\n", got)
	case want >= 11 && want <= 30:
		t.Logf("\nLoop Count Still Low! - %d\n", got)
	case want >= 31 && want <= 60:
		t.Logf("\nLoop Count Mid-Way! - %d\n", got)
	case want >= 61 && want <= 80:
		t.Logf("\nLoop Count Relatively High! - %d\n", got)
	default:
		t.Logf("\nLoop Count High! - %d", got)
	}
}
