//ff:func feature=token type=helper control=sequence
//ff:what misspeltIn — what misspeltCandidates gives for a stem as a test reads it: the names of its candidates joined with "," and the best verdict (an empty string and 0 for none)
package token

func misspeltIn(ix *index, right string, kinds, tail int, sole bool) (string, float64) {
	return candidatesIn(ix.misspeltCandidates(right, kinds, tail, sole))
}
