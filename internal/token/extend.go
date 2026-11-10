//ff:func feature=token type=reader control=iteration dimension=2
//ff:what extend — every reading that ends with tail t at j: each reading kept at a place before the tail's stem, grown by the word from there
package token

func (sc *scan) extend(cands []*path, best [][]*path, t tail, j int) []*path {
	for b := 0; b < t.start; b++ {
		for _, p := range best[b] {
			cands = append(cands, grow(p, sc.wordOf(b, j, t)))
		}
	}
	return cands
}
