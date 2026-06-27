package tac_test

import (
	"fmt"

	command "github.com/gloo-foo/cmd-tac"
	"github.com/gloo-foo/testable"
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
