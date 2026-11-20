//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what unmodifiable — may no adnominal stand right before a word: it has a known candidate and each of its known candidates (a new word's, a stem no table names, aside) is one no adnominal modifies (takesNoAdnominal), or its first candidate is a mark that ends a sentence (마침표, 물음표, 느낌표: stopMarks)
package parse

import "slices"

// A new word is the reading of last resort of letters no table knows
// as they stand; it tells nothing of what the grammar lets stand around
// them, and a word with a known reading is judged by that.
func unmodifiable(n Node) bool {
	first := nearOf(n.Contest.Branches[0])
	if first.Stem == "" && slices.Contains(stopMarks, first.Last) {
		return true
	}
	known := false
	for _, br := range n.Contest.Branches {
		near := nearOf(br)
		if near.Stem != "" && near.Kind == "" {
			continue
		}
		if !takesNoAdnominal(near) {
			return false
		}
		known = true
	}
	return known
}
