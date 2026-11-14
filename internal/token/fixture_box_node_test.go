//ff:func feature=token type=helper control=iteration dimension=1
//ff:what boxNode — a node made up for the box tests: its place and letters, and a candidate for each text given (one stem, verdict 1.00, 0.75, 0.50 … and preference -3.25, -6.50, -9.75 …, one step further for each)
package token

func boxNode(start, end int, text string, cands ...string) Node {
	n := Node{Span: Span{start, end}, Contest: Contest{Text: text}}
	for i, c := range cands {
		n.Contest.Branches = append(n.Contest.Branches, Branch{
			Words:   []Eojeol{{Parts: []Part{{Text: c, Stem: true}}}},
			Verdict: 1 - 0.25*float64(i),
			Pref:    -3.25 * float64(i+1),
		})
	}
	return n
}
