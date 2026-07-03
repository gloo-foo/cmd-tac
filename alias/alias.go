// Package alias provides unprefixed names for the tac command and its flags.
//
//	import tac "github.com/gloo-foo/cmd-tac/alias"
//	tac.Tac(tac.Sep(":"))
package alias

import (
	gloo "github.com/gloo-foo/framework"

	command "github.com/gloo-foo/cmd-tac"
)

// Tac re-exports the constructor: it reverses the order of input lines.
func Tac(opts ...any) gloo.Command[[]byte, []byte] { return command.Tac(opts...) }

// Sep re-exports the custom-record-separator option (-s): input is split on the
// given separator, the records are reversed, and each is emitted as its own line.
type Sep = command.TacSep
