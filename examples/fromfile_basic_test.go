package tac_test

import (
	"fmt"
	"os"

	"github.com/gloo-foo/testable"
	"github.com/gloo-foo/testable/run"

	command "github.com/gloo-foo/cmd-tac"
)

// This example demonstrates reading from a file instead of inline input.
func ExampleTac_fromFile_basic() {
	// tac testdata/text.txt
	data, _ := os.ReadFile("testdata/text.txt")
	output, _ := testable.Test(command.Tac(), run.Input(string(data)))
	fmt.Print(output)
	// Output:
	// Third
	// Second
	// First
}
