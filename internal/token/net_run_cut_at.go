//ff:func feature=token type=reader control=iteration dimension=1
//ff:what cutAt — a stretch cut off at the end of the first span of the agreed part that a stretch reads the same way (tryAt), the ends in order; a place where a sign shows a boundary has the candidates kept there worked out first (keptAt, so the states are the ones the unit's end would see) and then the cut is tried again; false when no end of the part serves
package token

func (r *netRun) cutAt(prefix []Span) (unitStretch, bool) {
	for k, sp := range prefix {
		if r.sc.late(sp.End) {
			r.sc.keptAt(r.best, sp.End)
			return r.tryCut()
		}
		if st, ok := r.tryAt(sp.End, prefix[:k+1]); ok {
			return st, true
		}
	}
	return unitStretch{}, false
}
