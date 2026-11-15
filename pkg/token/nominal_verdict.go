//ff:func feature=token type=rule control=iteration dimension=1
//ff:what nominalVerdict — how well a stem stands as a nominal of the tables with what follows it (the best verdict among its candidates of that kind, knownCandidates), its word taken as the whole of its unit when sole is true; 0 when it does not stand as one
package token

func (ix *index) nominalVerdict(stem string, tail int, sole bool) float64 {
	best := 0.0
	for _, r := range ix.knownCandidates(stem, tail, sole, false, false, false) {
		if r.kind == kindNominal {
			best = max(best, r.v)
		}
	}
	return best
}
