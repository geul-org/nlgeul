//ff:func feature=token type=helper control=iteration dimension=1
//ff:what pathInNet — is each word of a candidate a state of the net (by where it stands, how it is printed, its verdict and preference), and each one next to the one before it a pred of it
package token

import "slices"

func pathInNet(lat *lattice, p *path) bool {
	var before *path
	from := 0
	for _, e := range wordsOf(p) {
		s := stateFor(lat, from, e)
		if s == nil || (before != nil && !slices.Contains(s.preds, before)) {
			return false
		}
		before, from = s, e.end
	}
	return before != nil
}
