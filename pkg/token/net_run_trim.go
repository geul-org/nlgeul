//ff:func feature=token type=helper control=iteration dimension=1
//ff:what trim — the places after a cut: those before the cut are emptied (the states kept there, the candidates kept at a sign, which sign-between-letters places keep, the markers worked out there), the cut keeps the seeds of the roots that were among the states kept at it (best) and, apart, those of the roots that were among the candidates kept there at a sign (kept), so that a word after the cut grows from the same states as it would have without the cut, and the places after it keep only the states that stay, the best lambdas of the spans before it are forgotten; the scan is told where the first place is (floor)
package token

func (r *netRun) trim(q int, seedOf map[*path]*path, order []*path, alive map[*path]bool) {
	sc := r.sc
	seeds := seedsAmong(order, seedOf, r.best[q])
	keptSeeds := seedsAmong(order, seedOf, sc.kept[q])
	for k := r.c; k <= r.j; k++ {
		switch {
		case k < q:
			r.best[k], sc.marks[k] = nil, nil
		case k == q:
			r.best[k] = seeds
		default:
			r.best[k] = keepAlive(r.best[k], alive)
		}
	}
	for k, v := range sc.kept {
		switch {
		case k < q:
			delete(sc.kept, k)
		case k > q:
			sc.kept[k] = keepAlive(v, alive)
		}
	}
	for sp := range r.lam {
		if sp.Start < q {
			delete(r.lam, sp)
		}
	}
	sc.kept[q], sc.floor, r.seeds, r.keptSeeds = keptSeeds, q, seeds, keptSeeds
}
