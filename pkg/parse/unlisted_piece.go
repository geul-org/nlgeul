//ff:func feature=parse type=rule control=sequence
//ff:what unlistedPiece — what the new word's graph is told of a stem no table lists, from b before tail t (which ends at j), but whether it stands over a place where the script changes (unlisted tells that, crossesChange): the stem as it was before a piece was fused into it, what follows it, the kinds the tables know its letters as, whether its letters read as markers there (marked), and whether they hold, right after a run of Latin letters at the stem's head, only a word the spacing articles write apart (writtenApart)
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

func (sc *scan) unlistedPiece(b, j int, stem string, t tail) rule.Piece {
	kinds := sc.ix.lex[stem].kinds | sc.ix.closed[stem]
	return rule.Piece{Form: stem, Tail: tailClass(t.pieces), Stem: kinds, Marked: sc.marked(b, j, t), Apart: sc.writtenApart(b, j, t)}
}
