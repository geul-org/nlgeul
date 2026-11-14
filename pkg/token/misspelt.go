//ff:func feature=token type=rule control=iteration dimension=1
//ff:what misspelt — is the right form of a misspelt stem a known word with what follows it: the best verdict among its candidates (misspeltCandidates), 0 when it has none
package token

func (ix *index) misspelt(right string, kinds, tail int, sole bool) float64 {
	best := 0.0
	for _, r := range ix.misspeltCandidates(right, kinds, tail, sole) {
		best = max(best, r.v)
	}
	return best
}
