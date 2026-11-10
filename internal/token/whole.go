//ff:func feature=token type=reader control=iteration dimension=2
//ff:what whole — every reading that ends at j with a word of no markers: a closed-set word anywhere, any token sequence only at the unit's end (without a marker nothing shows a boundary inside a unit)
package token

func (sc *scan) whole(cands []*path, best [][]*path, j int) []*path {
	for b := 0; b < j; b++ {
		_, key := sc.ix.words[string(sc.s[b:j])]
		if !key && j < len(sc.s) {
			continue
		}
		for _, p := range best[b] {
			cands = append(cands, grow(p, sc.wordOf(b, j, tail{start: j})))
		}
	}
	return cands
}
