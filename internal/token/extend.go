//ff:func feature=token type=reader control=iteration dimension=2
//ff:what extend — every reading that ends with tail t at j: each reading kept at a place before the tail's stem, grown by the word from there; no word begins or has its stem end inside an atom
package token

func (sc *scan) extend(cands []*path, best [][]*path, t tail, j int) []*path {
	for b := 0; b < t.start; b++ {
		if len(best[b]) == 0 || sc.inside[b] || sc.inside[t.start] || (t.only != "" && string(sc.s[b:t.start]) != t.only) || !sc.gated(b, t) {
			continue
		}
		e := sc.wordOf(b, j, t)
		for _, p := range best[b] {
			cands = append(cands, grow(p, e))
		}
	}
	return cands
}
