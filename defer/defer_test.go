package main

import (
	"bytes"
	"os"
	"testing"
)

func TestDefFunc(t *testing.T) {
	// Defining test cases with simulated input/output.
	test := []struct {
		input  string
		output string
	}{
		{"Kirill", "Hello Kirill"},
		{"David", "Hello David"},
	}

	// loop through the 'test' cases
	for _, tt := range test {
		t.Run(tt.input, func(t *testing.T) {
			// create a pipe
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("Failed to create pipe: %v", err)
			}

			// save original input & output for later
			originalStdin := os.Stdin
			originalStdout := os.Stdout

			// redirect Stdin to read side of the pipe and os.Stdout to a buffer
			os.Stdin = r
			var buf bytes.Buffer
			os.Stdout = &buf

			// write the correct input into the pipe (simulate user input)
			_, err = w.Write([]byte(tt.input + "\n"))
			if err != nil {
				t.Fatalf("Failed to write into pipe: %v", err)
			}
			w.Close() // ensure pipe is closed after writing.

			// run function from main file for output.
			got := defFunc()

			// restore original input & output.
			os.Stdin = originalStdin
			os.Stdout = originalStdout

			// check if the return value matches the expected result
			if got != tt.input {
				t.Errorf("defFunc() = %v; want %v", got, tt.input)
			}

			// correct expected printed output, accounting for deferred println
			expectedOutput := "Input your name: \nHello " + tt.input + "\n"

			// check if expected values match actual printed output
			if buf.String() != expectedOutput {
				t.Errorf("Expected Printed Output: %v\nActual Printed Output: %v", expectedOutput, buf.String())
			}
		})
	}
}
