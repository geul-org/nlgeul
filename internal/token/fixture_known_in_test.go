//ff:func feature=token type=helper control=sequence
//ff:what knownIn — what knownReads gives for a stem as a test reads it: the names of its candidates joined with "," and the best verdict (an empty string and 0 for none)
package token

func knownIn(ix *index, stem string, tail int, sole bool) (string, float64) {
	return readsIn(ix.knownReads(stem, tail, sole))
}
