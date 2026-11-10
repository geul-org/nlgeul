//ff:func feature=token type=reader control=sequence
//ff:what stemPref — a stem's preference: log10 of its share of the word table's counts when the table has it with a count (the count of the kinds that take what follows it: 작 the noun is not as common as 작 the adjective); for words bound by a slash or a hyphen the sum of theirs; else log10(1/syllables) for each of its tokens
package token

import "math"

func (ix *index) stemPref(stem string, tail int, isKnown bool, tokens int) float64 {
	if count := ix.taken(stem, tail); isKnown && count > 0 {
		return math.Log10(float64(count) / ix.lexTotal)
	}
	if parts := partsBound(stem); isKnown && len(parts) > 1 {
		return ix.partsPref(parts)
	}
	return float64(tokens) * math.Log10(1.0/syllables)
}
