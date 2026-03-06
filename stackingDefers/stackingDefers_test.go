package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestDefer(t *testing.T) {
	// Expected output
	want := "counting...\n12345678910\n10987654321"

	// redirecting the standard output (stdout) to a pipe
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// call the function that prints to standard output
	stackDef()

	// close the writer end of the pipe to stop capturing
	w.Close()

	// create a buffer to capture the output (this implements io.Writer)
	var got bytes.Buffer
	// read the entire captured output from the pipe
	io.Copy(&got, r) // then copy all the output from the pipe into the buffer

	// restore the original stdout
	os.Stdout = old

	// compare the expected output with the captured output (convert got to string)
	if got.String() != want {
		t.Errorf("\nMISMATCH!\nEXPECTED: %v\nGOT: %v", want, got.String())
	} else {
		t.Logf("\nMATCH FOUND!\nEXPECTED: %v\nGOT: %v", want, got.String())
	}
}
