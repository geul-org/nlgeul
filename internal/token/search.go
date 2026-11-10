//ff:func feature=token type=reader control=iteration dimension=2
//ff:what search — read a unit token by token: at each place look up the markers that end there, grow the readings kept so far, and keep the best; the readings that reach the end come back, best first
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
		for _, t := range tails {
			cands = sc.extend(cands, best, t, j)
		}
		best[j] = sc.ix.top(sc.whole(cands, best, j, markedFrom(tails)), beam)
	}
	return sc.ix.top(best[n], branches)
}
