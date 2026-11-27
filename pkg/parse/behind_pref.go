//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what behindPref — the preference of a stem read as a candidate that stands behind every word of the word table with its letters (kindBehind: a number or a unit read with nothing in its word that tells it): the least of the preferences the word table's kinds give those letters (log10 of the kind's share of the table's counts; a kind with no count is valued as a new word is, by its tokens); a stem the word table lacks is valued as a new word, and no stem above a new word
package parse

import "math"

func (ix *index) behindPref(stem string, tokens int) float64 {
	fresh := float64(tokens) * math.Log10(1.0/syllables)
	w, ok := ix.lex[stem]
	if !ok {
		return fresh
	}
	least := fresh
	for i, k := range kindOrder {
		if w.kinds&k != 0 && w.by[i] > 0 {
			least = math.Min(least, math.Log10(float64(w.by[i])/ix.lexTotal))
		}
	}
	return least
}
