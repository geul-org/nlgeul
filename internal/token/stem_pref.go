//ff:func feature=token type=reader control=sequence
//ff:what stemPref — a stem's preference: log10 of its share of the word table's counts when the table has it with a count; else log10(1/syllables) for each of its tokens
package token

import "math"

func (ix *index) stemPref(stem string, isKnown bool, tokens int) float64 {
	if w := ix.lex[stem]; isKnown && w.count > 0 {
		return math.Log10(float64(w.count) / ix.lexTotal)
	}
	return float64(tokens) * math.Log10(1.0/syllables)
}
