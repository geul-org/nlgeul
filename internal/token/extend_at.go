//ff:func feature=token type=reader control=iteration dimension=1
//ff:what extendAt — every candidate that ends with tail t at j and has its stem begin at b: those kept at b (or kept there by keptBefore, where a cue of the spacing rules shows the stem as a word that follows a word no table lists), grown once for each way the tail is read after those letters, and for a stem of words bound by a slash or a hyphen once more; nothing where no candidate is kept at b, the stem is inside an atom, b is not a place a word may begin (starts), or the tail names a stem other than the letters
package token

func (sc *scan) extendAt(cands []*path, best [][]*path, t tail, b, j int) []*path {
	if (len(best[b]) == 0 && !sc.shows[b][t.start]) || sc.inside[b] || sc.inside[t.start] || (sc.starts != nil && !sc.starts[b]) || (t.only != "" && sc.letters(b, t.start) != t.only) {
		return cands
	}
	for _, read := range sc.variants(b, t) {
		if kept := sc.keptBefore(best, b, t.start, read); len(kept) > 0 {
			cands = sc.grown(cands, kept, read, b, j)
		}
	}
	if len(best[b]) > 0 {
		cands = sc.boundTails(cands, best[b], t, b, j)
	}
	return cands
}
