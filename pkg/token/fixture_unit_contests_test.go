//ff:func feature=token type=helper control=iteration dimension=2
//ff:what unitContests — the old reading the net is held against: each unit of a sentence read whole by oldUnit (every path spread), its contests in order, each marked with which unit it is in and whether no space stood before it (Glued: true for every contest of a unit after its first, and for the first after a unit with no space after it)
package token

func unitContests(r *Reader, sentence string) []Contest {
	var out []Contest
	spaced := true
	for i, u := range unitsOf(r.ix, sentence, false) {
		contests, _ := oldUnit(r.ix, u, false)
		for k, c := range contests {
			c.Unit, c.Glued = i, k > 0 || (i > 0 && !spaced)
			out = append(out, c)
		}
		spaced = u.spaced
	}
	return out
}
