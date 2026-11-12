//ff:func feature=token type=reader control=iteration dimension=2
//ff:what whole — every reading that ends at j with a word of no markers: a known word that may stand alone anywhere, any token sequence at the unit's end or where the script changes (without a marker nothing else shows a boundary inside a unit); never inside an atom, and a token sequence no word table lists only as the new word's graph lets it stand (not after a boundary when its letters read as markers there); a word of the word table that only stands alone grows none of the readings whose last word ends in a modifying closing or is a determiner (afterAdn), and where a determiner ends is kept; shown tells that a sign right after the place shows the boundary, so that a word no table lists stands without a change of script there; and the same word read as the right form of a word of the word table it is written wrong for (wholeTypos)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func (sc *scan) whole(cands []*path, best [][]*path, j int, shown bool) []*path {
	for b := 0; b < j && !sc.inside[j]; b++ {
		cands = sc.wholeTypos(cands, best, b, j)
		_, v := sc.ix.known(string(sc.s[b:j]), rule.TailBare, false)
		if len(best[b]) == 0 || (v <= 0 && !shown && j < len(sc.s) && !sc.change(j)) || (v <= 0 && sc.unlisted(b, j, string(sc.s[b:j]), tail{start: j}) <= 0) {
			continue
		}
		e, adn, det := sc.wordOf(b, j, tail{start: j}), sc.afterAdn(b, tail{start: j}, false), sc.afterAdn(b, tail{start: j}, true)
		reads := sc.numberReads(e, b, j)
		for _, p := range best[b] {
			if (adn && p.last.adn) || (det && p.last.det) {
				continue
			}
			for _, r := range reads {
				cands = append(cands, grow(p, r))
			}
			sc.detEnds[j] = sc.detEnds[j] || e.det
		}
	}
	return cands
}
