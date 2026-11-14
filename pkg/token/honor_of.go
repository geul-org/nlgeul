//ff:func feature=token type=reader control=sequence
//ff:what honorOf — the graph that judges a letter as the honorific 시 with a piece under it, and the graph the word before must pass: 셨·셔·십 always and with no such graph; 신·실·심·세 only with a word table, after a predicate's stem
package token

import "slices"

// The first value is nil for any other letter.
func (ix *index) honorOf(r rune) (*entry, *entry) {
	if slices.Contains(honorLetters, r) {
		return ix.honor, nil
	}
	return ix.wide[r], ix.pred
}
