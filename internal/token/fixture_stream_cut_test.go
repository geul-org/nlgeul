//ff:func feature=token type=helper control=iteration dimension=1
//ff:what streamCut — a way of cutting made by hand: for each span the word as a branch of its own (its letters, verdict 1, the preference given for it)
package token

func streamCut(spans []Span, texts []string, prefs []float64) cut {
	c := cut{spans: spans}
	for i, text := range texts {
		w := Eojeol{Parts: []Part{{Text: text, Stem: true}}, End: spans[i].End}
		c.words = append(c.words, Branch{Words: []Eojeol{w}, Verdict: 1, Pref: prefs[i]})
	}
	return c
}
