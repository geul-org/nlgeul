//ff:func feature=token type=reader control=sequence
//ff:what farKey — whether the letters from b to e are a key of the tables longer than a new word's stem may be, as they stand or with the last letter taken off (the stem of a word whose last syllable a piece was fused into)
package token

func (sc *scan) farKey(b, e int) bool {
	if sc.ix.farKeys[sc.letters(b, e)] {
		return true
	}
	return e > b+1 && sc.ix.farBases[sc.letters(b, e-1)]
}
