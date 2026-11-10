//ff:func feature=token type=reader control=iteration dimension=2
//ff:what whole — every reading that ends at j with a word of no markers: a known word that may stand alone anywhere, any token sequence only at the unit's end (without a marker nothing shows a boundary inside a unit)
package token

import "github.com/geul-org/koparse/pkg/rule"

func (sc *scan) whole(cands []*path, best [][]*path, j int) []*path {
	for b := 0; b < j; b++ {
		_, v := sc.ix.known(string(sc.s[b:j]), rule.TailBare)
		if v <= 0 && j < len(sc.s) {
			continue
		}
		e := sc.wordOf(b, j, tail{start: j})
		for _, p := range best[b] {
			cands = append(cands, grow(p, e))
		}
	}
	return cands
}
