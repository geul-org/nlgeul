//ff:func feature=token type=helper control=iteration dimension=1
//ff:what stateFor — the state of the net for a word that begins at from (by where it ends, how it is printed, its verdict and preference), nil when there is none
package token

func stateFor(lat *lattice, from int, e step) *path {
	for _, c := range lat.all {
		if c.start == from && c.last.end == e.end && c.last.sig == e.sig && c.last.verdict == e.verdict && c.last.pref == e.pref {
			return c
		}
	}
	return nil
}
