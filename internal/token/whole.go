//ff:func feature=token type=reader control=iteration dimension=2
//ff:what whole — every reading that ends at j with a word of no markers: a known word that may stand alone anywhere, any token sequence at the unit's end or where the script changes (without a marker nothing else shows a boundary inside a unit); never inside an atom, and a new word after a boundary not when its letters read as markers there (marked: the places such markers start)
package token

import "github.com/geul-org/koparse/pkg/rule"

func (sc *scan) whole(cands []*path, best [][]*path, j int, marked map[int]bool) []*path {
	for b := 0; b < j && !sc.inside[j]; b++ {
		_, v := sc.ix.known(string(sc.s[b:j]), rule.TailBare)
		if len(best[b]) == 0 || (v <= 0 && j < len(sc.s) && !sc.change(j)) || (v <= 0 && b > 0 && marked[b] && sc.ix.newWord(true) <= 0) {
			continue
		}
		e := sc.wordOf(b, j, tail{start: j})
		for _, p := range best[b] {
			cands = append(cands, grow(p, e))
		}
	}
	return cands
}
