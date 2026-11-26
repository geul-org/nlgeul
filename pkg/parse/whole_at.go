//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what wholeAt — every candidate that ends at j with the letters from b as a word of no markers: those written wrong (wholeTypos), and a known word or a token sequence that stands there (a new word only where its graph lets it stand, and where no change of script or sign shows the boundary, only as a cue shows it, or where a path ends in a slash before Hangul that does not read as markers on it: sc.cut, sc.onSlash), grown from each candidate kept at b once for each way its stem stands; letters no table lists stand there in no way when the new word's graph turns them down for what unlistedPiece tells, and when it turns them down only for the place where the script changes inside them (crossesChange) the whole new word is not grown while the other ways their stem stands are (an atom's names: 제1-1 as 차례 and 범위); nothing where b is not a place a word may begin (starts)
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// Past the gate, a way of the letters with a verdict of 0 or less is
// the new word that the place where the script changes turns down: the
// other ways (a word of the tables, an atom, a bound stem, 사이시옷)
// stand only with a verdict above 0.
func (sc *scan) wholeAt(cands []*path, best [][]*path, b, j int, shown bool) []*path {
	if sc.starts != nil && !sc.starts[b] {
		return cands
	}
	cands = sc.wholeTypos(cands, best, b, j)
	v := sc.ix.known(sc.letters(b, j), rule.TailBare, false)
	kept := sc.keptBefore(best, b, j, tail{start: j})
	if len(kept) == 0 || (v <= 0 && !shown && j < len(sc.s) && !sc.change(j) && (j != sc.cut || sc.onSlash) && !sc.stood(b, j)) || (v <= 0 && sc.ix.judged(sc.ix.stem, sc.unlistedPiece(b, j, sc.letters(b, j), tail{start: j}), "") <= 0) {
		return cands
	}
	for _, e := range sc.stepsOf(b, j, tail{start: j}) {
		if v <= 0 && e.verdict <= 0 {
			continue
		}
		cands = sc.growAlone(cands, kept, e, b, j)
	}
	return cands
}
