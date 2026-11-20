//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what net — read a unit token by token as search does (the same places, tails, words and graphs), but keep the candidates that end at a place as states, one for each last word: the whole unit's net at once, with no window (the stream reads it by netRun); what comes back is the states that end the unit and every state made, nil for a unit of no letters
package parse

func (sc *scan) net() *lattice {
	if len(sc.s) == 0 {
		return nil
	}
	best := sc.netStart()
	for j := 1; j <= len(sc.s); j++ {
		sc.netStep(best, j, nil)
	}
	return sc.netClose(best, 0)
}
