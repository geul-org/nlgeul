//ff:func feature=token type=reader control=iteration dimension=1
//ff:what extend — every candidate that ends with tail t at j: each candidate kept at a place before the tail's stem (or, where a cue of the spacing rules shows the stem as a word that follows a word no table lists, kept there by keptBefore), grown by the word from there, once for each way the tail is read after those letters (as they stand, or as a changed form of a stem of the word table, and when the stem is words bound by a slash or a hyphen with a word written wrong, as the right form of that word); no word begins or has its stem end inside an atom; the places the stem begins at are, in order, the whole unit (j at the unit's end), the places farther than a new word's stem may be where a known word may begin, and the places within newStemMax tokens (beginnings), or the one place a tail that names its stem allows
package token

func (sc *scan) extend(cands []*path, best [][]*path, t tail, j int) []*path {
	if t.only != "" {
		if b := t.start - len([]rune(t.only)); b >= 0 {
			cands = sc.extendAt(cands, best, t, b, j)
		}
		return cands
	}
	begins := sc.beginsOf(t.start)
	if j == len(sc.s) && t.start > 0 && (len(begins) == 0 || begins[0] != 0) {
		cands = sc.extendAt(cands, best, t, 0, j)
	}
	for _, b := range begins {
		cands = sc.extendAt(cands, best, t, b, j)
	}
	return cands
}
