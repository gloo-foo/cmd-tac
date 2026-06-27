package command

// tacSepFlag sets a custom record separator (-s flag).
// When set, input is split on this separator instead of newlines,
// the resulting records are reversed, and output is rejoined with the separator.
type tacSepFlag string

// TacSep sets a custom record separator.
func TacSep(s string) tacSepFlag { return tacSepFlag(s) }

func (f tacSepFlag) Configure(flags *flags) { flags.separator = string(f) }

type flags struct {
	separator string
}
