package command

// TacSep is the custom record separator option (-s).
// When given, input is split on this separator instead of newlines,
// the resulting records are reversed, and each record is emitted as its
// own output line.
type TacSep string
