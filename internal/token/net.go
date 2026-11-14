//ff:func feature=token type=reader control=iteration dimension=2
//ff:what net — read a unit token by token as search does (the same places, tails, words and graphs), but keep the candidates that end at a place as states, one for each last word (the only thing of the words before it that a word after it asks of): a state keeps every state before it (preds) instead of each way to it, so a unit that has very many ways is read once for each stretch; what comes back is the net, the states that end the unit and every state made, nil for a unit of no letters
package token

func (sc *scan) net() *lattice {
	n := len(sc.s)
	if n == 0 {
		return nil
	}
	sc.netted = true
	sc.prepare()
	best := make([][]*path, n+1)
	best[0] = []*path{sc.registry().root}
	for j := 1; j <= n; j++ {
		var cands []*path
		tails := sc.tailsAt(j)
		sc.marks[j] = markedFrom(tails)
		for _, t := range tails {
			cands = sc.extend(cands, best, t, j)
		}
		cands = sc.signWords(cands, best, j)
		best[j] = sc.settle(sc.whole(cands, best, j, false))
	}
	reg := sc.registry()
	return &lattice{n: n, ends: sc.settle(append(best[n], sc.attached()...)), all: reg.all, twins: reg.twins}
}
