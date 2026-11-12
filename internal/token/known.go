//ff:func feature=token type=rule control=iteration dimension=1
//ff:what known — is a stem a known word with what follows it: the best verdict among its readings (knownReads), 0 when it has none
package token

func (ix *index) known(stem string, tail int, sole bool) float64 {
	best := 0.0
	for _, r := range ix.knownReads(stem, tail, sole) {
		best = max(best, r.v)
	}
	return best
}
