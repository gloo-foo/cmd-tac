// Package alias provides unprefixed names for the tac command and its flags.
//
//	import tac "github.com/gloo-foo/cmd-tac/alias"
//	tac.Tac(tac.Sep(":"))
package alias

import command "github.com/gloo-foo/cmd-tac"

// Tac re-exports the constructor: it reverses the order of input lines.
var Tac = command.Tac

// Sep re-exports the custom-record-separator flag (-s): input is split on the
// given separator, the records are reversed, and output is rejoined.
var Sep = command.TacSep
