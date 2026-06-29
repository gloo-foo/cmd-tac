package alias_test

import (
	"slices"
	"testing"

	"github.com/gloo-foo/testable"

	tac "github.com/gloo-foo/cmd-tac/alias"
)

// The alias package re-exports the constructor and the separator flag under
// unprefixed names. A mis-wired re-export (say, Sep bound to the wrong function,
// or Tac bound to the wrong constructor) compiles cleanly, so only behavior can
// prove the wiring. Each test exercises one re-export and asserts the GNU tac
// output it must produce.

func assertLines(t *testing.T, got, want []string) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestAlias_TacReversesLineOrder(t *testing.T) {
	// Bare Tac reverses the order of input lines.
	lines, err := testable.TestLines(tac.Tac(), "line1\nline2\nline3\n")
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"line3", "line2", "line1"})
}

func TestAlias_SepReversesRecords(t *testing.T) {
	// Sep(":") splits the single input line on ":" and reverses the records.
	lines, err := testable.TestLines(tac.Tac(tac.Sep(":")), "a:b:c\n")
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"c", "b", "a"})
}
