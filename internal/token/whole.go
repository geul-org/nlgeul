//ff:func feature=token type=reader control=iteration dimension=1
//ff:what whole — every candidate that ends at j with a word of no markers: a known word that may stand alone anywhere, any token sequence at the unit's end or where the script changes (without a marker nothing else shows a boundary inside a unit), once for each way its stem stands (a kind, a set, a name); never inside an atom, and a token sequence no word table lists only as the new word's graph lets it stand (not after a boundary when its letters read as markers there); a word of the word table that only stands alone grows none of the candidates whose last word ends in a modifying closing or is a determiner (afterAdn), and where a determiner ends is kept; shown tells that a sign right after the place shows the boundary, so that a word no table lists stands without a change of script there (so does a cue of the spacing rules that makes the letters a word that stands, stood); a word that a cue shows after a word no table lists grows the candidates kept where that word begins (keptBefore); and the same word read as the right form of a word of the word table it is written wrong for (wholeTypos); the places it begins at are those of extend (beginnings)
package token

func (sc *scan) whole(cands []*path, best [][]*path, j int, shown bool) []*path {
	if sc.inside[j] {
		return cands
	}
	begins := sc.beginsOf(j)
	if j == len(sc.s) && j > 0 && (len(begins) == 0 || begins[0] != 0) {
		cands = sc.wholeAt(cands, best, 0, j, shown)
	}
	for _, b := range begins {
		cands = sc.wholeAt(cands, best, b, j, shown)
	}
	return cands
}
