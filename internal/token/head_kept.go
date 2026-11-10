//ff:func feature=token type=helper control=selection
//ff:what headKept — where a unit's core begins when the signs right before it belong to it: the head of a path (/, ~/, ./, ../, with the dot of a hidden name), the dashes of an option before a Latin letter, or one of @ # $ _ . before a Latin letter or digit; else where it began
package token

import (
	"slices"
	"strings"
	"unicode"
)

// pathHeads are what a path begins with before its first name.
var pathHeads = []string{"/", "~/", "./", "../", "/.", "~/.", "./.", "../."}

func headKept(all []rune, from int) int {
	stop := signsBefore(all, from)
	head, next := string(all[stop:from]), atOf(all, from)
	latin := next < unicode.MaxASCII && (unicode.IsLetter(next) || unicode.IsDigit(next) || next == '_')
	switch {
	case head == "":
		return from
	case slices.Contains(pathHeads, head):
		return stop
	case (head == "-" || head == "--") && latin && !unicode.IsDigit(next):
		return stop
	case latin && strings.ContainsRune("@#$_.", all[from-1]):
		return from - 1
	}
	return from
}
