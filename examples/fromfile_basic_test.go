package tac_test

import (
	"fmt"
	"os"

	"github.com/gloo-foo/testable"

	command "github.com/gloo-foo/cmd-tac"
)

// This example demonstrates reading from a file instead of inline input.
func ExampleTac_fromFile_basic() {
	// tac testdata/text.txt
	data, _ := os.ReadFile("testdata/text.txt")
	output, _ := testable.Test(command.Tac(), string(data))
	fmt.Print(output)
	// Output:
	// Third
	// Second
	// First
}
