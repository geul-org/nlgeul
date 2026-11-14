//ff:func feature=token type=reader control=sequence
//ff:what tryAt — the stretch from the last cut to place q (the states on the ways that end at q: between) made as the stream takes it, when its best way is the first spans of the agreed part (want); then the cut is made (commit); false when it reads another way
package token

func (r *netRun) tryAt(q int, want []Span) (unitStretch, bool) {
	pt := r.sc.between(r.c, q)
	st := r.sc.stretchOf(r.u, pt, r.first && r.c == 0)
	rg := newRegionNet(Contest{Text: string(r.sc.s[r.c:q])}, r.c, *st.cells)
	seg, ok := rg.peek()
	if !ok || !sameSpans(seg.Nodes, want) {
		return unitStretch{}, false
	}
	return r.commit(q, pt, seg, st), true
}
