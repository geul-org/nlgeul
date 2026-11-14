//ff:func feature=token type=reader control=iteration dimension=1
//ff:what wholeAt — every candidate that ends at j with the letters from b as a word of no markers: those written wrong (wholeTypos), and a known word or a token sequence that stands there (a new word only where its graph lets it stand, and where no change of script or sign shows the boundary, only as a cue shows it), grown from each candidate kept at b once for each way its stem stands; nothing where b is not a place a word may begin (starts)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func (sc *scan) wholeAt(cands []*path, best [][]*path, b, j int, shown bool) []*path {
	if sc.starts != nil && !sc.starts[b] {
		return cands
	}
	cands = sc.wholeTypos(cands, best, b, j)
	v := sc.ix.known(sc.letters(b, j), rule.TailBare, false)
	kept := sc.keptBefore(best, b, j, tail{start: j})
	if len(kept) == 0 || (v <= 0 && !shown && j < len(sc.s) && !sc.change(j) && !sc.stood(b, j)) || (v <= 0 && sc.unlisted(b, j, sc.letters(b, j), tail{start: j}) <= 0) {
		return cands
	}
	for _, e := range sc.stepsOf(b, j, tail{start: j}) {
		cands = sc.growAlone(cands, kept, e, b, j)
	}
	return cands
}
