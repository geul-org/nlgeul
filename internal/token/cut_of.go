//ff:func feature=token type=helper control=iteration dimension=1
//ff:what cutOf — a branch as one way of cutting its stretch: the span of each word and the word as a branch of its own with its verdict, preference and number of birth; at is the place of the stretch in the sentence and n its letters; the words' ends are token offsets of the unit, so the last end stands for the stretch's end and the others lie before it by the difference
package token

func cutOf(br Branch, at, n int) cut {
	scores := scoresOf(br)
	base := br.Words[len(br.Words)-1].End - n
	var c cut
	start := at
	for i, w := range br.Words {
		end := at + w.End - base
		c.spans = append(c.spans, Span{Start: start, End: end})
		c.words = append(c.words, Branch{Words: []Eojeol{w}, Verdict: scores[i].verdict, Pref: scores[i].pref, scores: []wordScore{scores[i]}})
		start = end
	}
	return c
}
