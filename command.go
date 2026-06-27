package command

import (
	"slices"
	"strings"

	gloo "github.com/gloo-foo/framework"
	"github.com/gloo-foo/framework/patterns"
)

// Tac returns a command that reverses the order of input lines.
// Use TacSep(s) (-s) to split on a custom record separator instead of newlines.
func Tac(opts ...any) gloo.Command[[]byte, []byte] {
	f := gloo.NewParameters[gloo.File, flags](opts...).Flags

	if f.separator != "" {
		return tacWithSeparator(f.separator)
	}

	return patterns.Accumulate(func(lines [][]byte) ([][]byte, error) {
		slices.Reverse(lines)
		return lines, nil
	})
}

// tacWithSeparator joins all input lines, splits on the custom separator,
// reverses the records, and emits each as a separate output line.
func tacWithSeparator(sep string) gloo.Command[[]byte, []byte] {
	return patterns.Accumulate(func(lines [][]byte) ([][]byte, error) {
		// Reconstruct original input by joining lines with newlines
		parts := make([]string, len(lines))
		for i, l := range lines {
			parts[i] = string(l)
		}
		joined := strings.Join(parts, "\n")

		// Split on the custom separator
		records := strings.Split(joined, sep)
		slices.Reverse(records)

		result := make([][]byte, len(records))
		for i, r := range records {
			result[i] = []byte(r)
		}
		return result, nil
	})
}
