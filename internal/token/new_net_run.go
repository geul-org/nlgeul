//ff:func feature=token type=reader control=sequence
//ff:what newNetRun — begin the net of a unit that is worked out as the stream asks for stretches: the scan is made ready and the root is kept at place 0 (netStart), the window cuts unless the index was made to work without it; a unit with no letters has nothing to work out; first tells that the first stretch may touch the unit's head, plant the words planted at each place (the letters after a sign of a pair), and stats is what the net counts of itself
package token

func newNetRun(sc *scan, u unit, first bool, plant func(j int) []*path, stats *netStats) *netRun {
	r := &netRun{sc: sc, u: u, n: len(sc.s), first: first, plant: plant, stats: stats, window: !sc.ix.noWindow, lam: map[Span]spanLambda{}}
	if r.n == 0 {
		r.done = true
		return r
	}
	r.best = sc.netStart()
	r.base = sc.registry().next
	return r
}
