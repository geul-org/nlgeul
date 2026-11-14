//ff:func feature=token type=helper control=sequence
//ff:what take — a walk that has reached the end of the region as a segment, when it is one of the ways of cutting there are; a way made by joining words of different ways is not one of them: it is counted (skipped) and nothing is given (the cuts there are are only put in order)
package token

func (rg *region) take(w *walk) (Segment, bool) {
	ids := w.ids()
	if !rg.member(rg.spansAt(ids)) {
		rg.skipped++
		return Segment{}, false
	}
	return rg.segment(ids), true
}
