package command_test

import (
	"testing"

	"github.com/gloo-foo/testable"

	command "github.com/gloo-foo/cmd-tac"
)

func TestTac_MultipleLines(t *testing.T) {
	lines, err := testable.TestLines(command.Tac(), "line1\nline2\nline3\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 3 || lines[0] != "line3" || lines[1] != "line2" || lines[2] != "line1" {
		t.Fatalf("got %v", lines)
	}
}

func TestTac_SingleLine(t *testing.T) {
	lines, err := testable.TestLines(command.Tac(), "only\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "only" {
		t.Fatalf("got %v", lines)
	}
}

func TestTac_EmptyInput(t *testing.T) {
	lines, err := testable.TestLines(command.Tac(), "")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 0 {
		t.Fatalf("expected empty, got %v", lines)
	}
}

func TestTac_ManyLines(t *testing.T) {
	input := "a\nb\nc\nd\ne\nf\ng\nh\ni\nj\n"
	lines, err := testable.TestLines(command.Tac(), input)
	if err != nil {
		t.Fatal(err)
	}
	expected := []string{"j", "i", "h", "g", "f", "e", "d", "c", "b", "a"}
	if len(lines) != len(expected) {
		t.Fatalf("expected %d lines, got %d", len(expected), len(lines))
	}
	for i, want := range expected {
		if lines[i] != want {
			t.Fatalf("line %d: got %q, want %q", i, lines[i], want)
		}
	}
}

func TestTac_Separator(t *testing.T) {
	// Input: "a:b:c" (single line), separator ":"
	// Split on ":" => ["a", "b", "c"], reverse => ["c", "b", "a"]
	lines, err := testable.TestLines(command.Tac(command.TacSep(":")), "a:b:c\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 3 || lines[0] != "c" || lines[1] != "b" || lines[2] != "a" {
		t.Fatalf("got %v, want [c b a]", lines)
	}
}

func TestTac_SeparatorMultiChar(t *testing.T) {
	// Input: "one--two--three" with separator "--"
	lines, err := testable.TestLines(command.Tac(command.TacSep("--")), "one--two--three\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 3 || lines[0] != "three" || lines[1] != "two" || lines[2] != "one" {
		t.Fatalf("got %v, want [three two one]", lines)
	}
}

func TestTac_SeparatorNoMatch(t *testing.T) {
	// Separator not found: entire input is one record
	lines, err := testable.TestLines(command.Tac(command.TacSep("|")), "hello world\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || lines[0] != "hello world" {
		t.Fatalf("got %v, want [hello world]", lines)
	}
}

func TestTac_DefaultStillWorks(t *testing.T) {
	// Ensure default (no separator) still reverses by lines
	lines, err := testable.TestLines(command.Tac(), "x\ny\nz\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 3 || lines[0] != "z" || lines[1] != "y" || lines[2] != "x" {
		t.Fatalf("got %v, want [z y x]", lines)
	}
}

func TestTac_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{"three", "a\nb\nc\n", []string{"c", "b", "a"}},
		{"two", "first\nsecond\n", []string{"second", "first"}},
		{"single", "alone\n", []string{"alone"}},
		{"numbers", "1\n2\n3\n", []string{"3", "2", "1"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := testable.TestLines(command.Tac(), tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if len(lines) != len(tt.expected) {
				t.Fatalf("expected %d lines, got %d", len(tt.expected), len(lines))
			}
			for i, want := range tt.expected {
				if lines[i] != want {
					t.Fatalf("line %d: got %q, want %q", i, lines[i], want)
				}
			}
		})
	}
}
