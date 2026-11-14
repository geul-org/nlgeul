//ff:func feature=token type=reader control=iteration dimension=2
//ff:what Contests — read a sentence and send up its words as a structure: the sentence is cut into units at spaces and at the signs of a pair, and each unit comes back as contests in order, each marked with which unit it is in and whether no space stood before it (Glued)
package token

// A layer above reads this instead of the printed lines: each contest
// holds its branches, each branch its words, each word its parts with
// their candidates. Lines prints the same contests.
func (r *Reader) Contests(sentence string) []Contest {
	var out []Contest
	spaced := true
	for i, u := range unitsOf(r.ix, sentence, false) {
		contests, _ := readUnit(r.ix, u, false)
		for k, b := range contests {
			b.Unit, b.Glued = i, k > 0 || (i > 0 && !spaced)
			out = append(out, b)
		}
		spaced = u.spaced
	}
	return out
}
