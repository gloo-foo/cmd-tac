package tac_test

import (
	"fmt"

	"github.com/gloo-foo/testable"

	command "github.com/gloo-foo/cmd-tac"
)

func ExampleTac_basic() {
	// echo "First\nSecond\nThird" | tac
	output, _ := testable.Test(command.Tac(), "First\nSecond\nThird")
	fmt.Print(output)
	// Output:
	// Third
	// Second
	// First
}
