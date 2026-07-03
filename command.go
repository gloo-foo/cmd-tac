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
	var sep TacSep
	rest := make([]any, 0, len(opts))
	for _, o := range opts {
		if s, ok := o.(TacSep); ok {
			sep = s
			continue
		}
		rest = append(rest, o)
	}
	// Classify the remaining arguments as before; tac reads from the pipeline,
	// so positionals are not consumed, but unsupported kinds still warn.
	gloo.NewParameters[gloo.File, struct{}](rest...)

	if sep != "" {
		return tacWithSeparator(sep)
	}

	return patterns.Accumulate(func(lines [][]byte) ([][]byte, error) {
		slices.Reverse(lines)
		return lines, nil
	})
}

// tacWithSeparator joins all input lines, splits on the custom separator,
// reverses the records, and emits each as a separate output line.
func tacWithSeparator(sep TacSep) gloo.Command[[]byte, []byte] {
	return patterns.Accumulate(func(lines [][]byte) ([][]byte, error) {
		// Reconstruct original input by joining lines with newlines
		parts := make([]string, len(lines))
		for i, l := range lines {
			parts[i] = string(l)
		}
		joined := strings.Join(parts, "\n")

		// Split on the custom separator
		records := strings.Split(joined, string(sep))
		slices.Reverse(records)

		result := make([][]byte, len(records))
		for i, r := range records {
			result[i] = []byte(r)
		}
		return result, nil
	})
}
