//ff:func feature=token type=reader control=iteration dimension=1
//ff:what net — read a unit token by token as search does (the same places, tails, words and graphs), but keep the candidates that end at a place as states, one for each last word (the only thing of the words before it that a word after it asks of): a state keeps every state before it (preds) instead of each way to it, so a unit that has very many ways is read once for each stretch; what comes back is the net, the states that end the unit and every state made, nil for a unit of no letters (the places are worked out one after another by netStep)
package token

func (sc *scan) net() *lattice {
	if len(sc.s) == 0 {
		return nil
	}
	best := sc.netStart()
	for j := 1; j <= len(sc.s); j++ {
		sc.netStep(best, j, nil)
	}
	return sc.netEnd(best)
}
