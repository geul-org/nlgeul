//ff:func feature=token type=rule control=sequence
//ff:what unlisted — the verdict of the new word's graph for a stem no table lists, from b before tail t (which ends at j): the graph is told the stem as it was before a piece was fused into it, what follows it, the kinds the tables know its letters as, and whether its letters read as markers there
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func (sc *scan) unlisted(b, j int, stem string, t tail) float64 {
	kinds := sc.ix.lex[stem].kinds | sc.ix.closed[stem]
	return verdict(sc.ix.stem, rule.Piece{Form: stem, Tail: tailClass(t.pieces), Stem: kinds, Marked: sc.marked(b, j, t)})
}
