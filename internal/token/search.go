//ff:func feature=token type=reader control=iteration dimension=2
//ff:what search — read a unit token by token: at each place look up the markers that end there, and grow every reading that reaches the place before them (only a reading a graph rejects is dropped; none is left out for its place among the others); the readings that reach the end come back, best first, with the readings of the whole unit as markers that attach to the unit before
package token

func (sc *scan) search() []*path {
	n := len(sc.s)
	if n == 0 {
		return nil
	}
	sc.prepare()
	best := make([][]*path, n+1)
	best[0] = []*path{{verdict: 1}}
	for j := 1; j <= n; j++ {
		var cands []*path
		tails := sc.tailsAt(j)
		sc.marks[j] = markedFrom(tails)
		for _, t := range tails {
			cands = sc.extend(cands, best, t, j)
		}
		best[j] = sc.ix.top(sc.whole(cands, best, j))
	}
	return sc.ix.top(append(best[n], sc.attached()...))
}
