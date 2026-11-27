//ff:func feature=parse type=helper control=sequence
//ff:what shortStem — is the stem before a tail too short for the tail's suffix: the tail's first piece is a suffix of longSuffixes (용, 제) and the stem is one syllable (군+용: a word of a syllable and such a suffix is a word of its own, the word table's business)
package parse

import "slices"

func shortStem(t tail, stem string) bool {
	return len(t.pieces) > 0 && slices.Contains(longSuffixes, t.pieces[0].text) && len([]rune(stem)) < 2
}
