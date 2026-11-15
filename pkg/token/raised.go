//ff:func feature=token type=helper control=iteration dimension=1
//ff:what raised — a candidate with the lifts that name it (by how it is printed) added to its preference, in the branch and in the score kept of its word (a copy of the scores)
package token

func raised(br Branch, own []lift) Branch {
	add, word := 0.0, br.String()
	for _, l := range own {
		if l.word == word {
			add += l.add
		}
	}
	if add == 0 {
		return br
	}
	br.Pref += add
	if len(br.scores) > 0 {
		br.scores = append([]wordScore{}, br.scores...)
		br.scores[0].pref += add
	}
	return br
}
