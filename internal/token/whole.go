//ff:func feature=token type=reader control=iteration dimension=2
//ff:what whole — every reading that ends at j with a word of no markers: a known word that may stand alone anywhere, any token sequence at the unit's end or where the script changes (without a marker nothing else shows a boundary inside a unit); never inside an atom
package token

import "github.com/geul-org/koparse/pkg/rule"

func (sc *scan) whole(cands []*path, best [][]*path, j int) []*path {
	for b := 0; b < j && !sc.inside[j]; b++ {
		_, v := sc.ix.known(string(sc.s[b:j]), rule.TailBare)
		if len(best[b]) == 0 || (v <= 0 && j < len(sc.s) && !sc.change(j)) {
			continue
		}
		e := sc.wordOf(b, j, tail{start: j})
		for _, p := range best[b] {
			cands = append(cands, grow(p, e))
		}
	}
	return cands
}
