//ff:func feature=token type=helper control=iteration dimension=1
//ff:what raised — a candidate with the lifts that name it (by how it is printed) applied, in the branch and in the score kept of its word (a copy of the scores): what they add is added to its preference, and its verdict is multiplied by 2/(1+force)-1 for each force of rebuttal (several multiply; a force of 0 multiplies by 1)
package token

func raised(br Branch, own []lift) Branch {
	add, times, word := 0.0, 1.0, br.String()
	for _, l := range own {
		if l.word != word {
			continue
		}
		add += l.add
		times *= 2/(1+l.force) - 1
	}
	if add == 0 && times == 1 {
		return br
	}
	br.Pref += add
	br.Verdict *= times
	if len(br.scores) > 0 {
		br.scores = append([]wordScore{}, br.scores...)
		br.scores[0].pref += add
		br.scores[0].verdict *= times
	}
	return br
}
