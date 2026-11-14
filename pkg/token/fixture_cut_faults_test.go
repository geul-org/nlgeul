//ff:func feature=token type=helper control=iteration dimension=2
//ff:what cutFaults — the contests whose branches of more than one word do not end alike (the last word's End), or whose cut words do not follow one another within the stretch (from where it begins up to its letters); the letters of such a contest
package token

func cutFaults(contests []Contest) []string {
	var bad []string
	at := 0
	for _, c := range contests {
		n := len([]rune(c.Text))
		last := map[int]bool{}
		for _, b := range c.Branches {
			if len(b.Words) > 1 {
				last[b.Words[len(b.Words)-1].End] = true
			}
			if !cutFits(cutOf(b, at, n), at, n) {
				bad = append(bad, c.Text)
			}
		}
		if len(last) > 1 {
			bad = append(bad, c.Text)
		}
		at += n
	}
	return bad
}
