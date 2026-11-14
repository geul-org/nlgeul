//ff:func feature=token type=reader control=sequence
//ff:what commit — make the cut at place q: how wide the window was is noted, the stretch is given what it takes to be read again (redo; not the first stretch of the letters after a sign of a pair, which grows from planted words), the last states of its best way are found (lastStates) and the net is kept only from it on (prune), and the cut is where the next stretch begins
package token

func (r *netRun) commit(q int, pt portion, seg Segment, st unitStretch) unitStretch {
	r.stats.cuts++
	r.stats.widths = append(r.stats.widths, r.sc.tokens(r.c, r.j))
	if r.plant == nil || r.c > 0 {
		st.redo = &redo{sc: r.sc, u: r.u, from: r.c, to: q, seeds: r.seeds, keptSeeds: r.keptSeeds, first: r.first && r.c == 0, base: r.base}
	}
	r.prune(q, lastStates(pt, seg))
	r.c = q
	return st
}
