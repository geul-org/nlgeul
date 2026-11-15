//ff:func feature=token type=rule control=sequence
//ff:what unlisted — the verdict of the new word's graph for a stem no table lists, from b before tail t (which ends at j): the graph is told what unlistedPiece tells (the stem as it was before a piece was fused into it, what follows it, the kinds the tables know its letters as, whether its letters read as markers there, whether they hold, right after a run of Latin letters at the stem's head, only a word the spacing articles write apart) and whether it stands over a place where the script changes whose two sides each read as a known word or a run of Latin letters and digits (crossesChange)
package token

func (sc *scan) unlisted(b, j int, stem string, t tail) float64 {
	p := sc.unlistedPiece(b, j, stem, t)
	p.Change = sc.crossesChange(b, j, t)
	return sc.ix.judged(sc.ix.stem, p, "")
}
