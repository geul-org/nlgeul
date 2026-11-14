//ff:func feature=token type=reader control=iteration dimension=1
//ff:what whole — every candidate that ends at j with a word of no markers: a known word that may stand alone anywhere, any token sequence at the unit's end or where the script changes (without a marker nothing else shows a boundary inside a unit), once for each way its stem stands (a kind, a set, a name); never inside an atom, and a token sequence no word table lists only as the new word's graph lets it stand (not after a boundary when its letters read as markers there); a word of the word table that only stands alone grows none of the candidates whose last word ends in a modifying closing or is a determiner (afterAdn), and where a determiner ends is kept; shown tells that a sign right after the place shows the boundary, so that a word no table lists stands without a change of script there (so does a cue of the spacing rules that makes the letters a word that stands, stood); a word that a cue shows after a word no table lists grows the candidates kept where that word begins (keptBefore); and the same word read as the right form of a word of the word table it is written wrong for (wholeTypos)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func (sc *scan) whole(cands []*path, best [][]*path, j int, shown bool) []*path {
	for b := 0; b < j && !sc.inside[j]; b++ {
		cands = sc.wholeTypos(cands, best, b, j)
		v := sc.ix.known(string(sc.s[b:j]), rule.TailBare, false)
		kept := sc.keptBefore(best, b, j, tail{start: j})
		if len(kept) == 0 || (v <= 0 && !shown && j < len(sc.s) && !sc.change(j) && !sc.stood(b, j)) || (v <= 0 && sc.unlisted(b, j, string(sc.s[b:j]), tail{start: j}) <= 0) {
			continue
		}
		for _, e := range sc.stepsOf(b, j, tail{start: j}) {
			cands = sc.growAlone(cands, kept, e, b, j)
		}
	}
	return cands
}
