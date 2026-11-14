//ff:func feature=token type=reader control=iteration dimension=2
//ff:what extend — every candidate that ends with tail t at j: each candidate kept at a place before the tail's stem (or, where a cue of the spacing rules shows the stem as a word that follows a word no table lists, kept there by keptBefore), grown by the word from there, once for each way the tail is read after those letters (as they stand, or as a changed form of a stem of the word table, and when the stem is words bound by a slash or a hyphen with a word written wrong, as the right form of that word); no word begins or has its stem end inside an atom
package token

func (sc *scan) extend(cands []*path, best [][]*path, t tail, j int) []*path {
	for b := 0; b < t.start; b++ {
		if (len(best[b]) == 0 && !sc.shows[b][t.start]) || sc.inside[b] || sc.inside[t.start] || (t.only != "" && string(sc.s[b:t.start]) != t.only) {
			continue
		}
		for _, read := range sc.variants(b, t) {
			if kept := sc.keptBefore(best, b, t.start, read); len(kept) > 0 {
				cands = sc.grown(cands, kept, read, b, j)
			}
		}
		if len(best[b]) > 0 {
			cands = sc.boundTails(cands, best[b], t, b, j)
		}
	}
	return cands
}
