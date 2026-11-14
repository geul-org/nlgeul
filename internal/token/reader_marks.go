//ff:func feature=token type=reader control=iteration dimension=1
//ff:what Marks — for measuring: one line for each unit, its text, then for each candidate, best first, the word ends (token offsets) each with the spacing rules (Options.Spacing) that the boundary after it is a place of, candidates separated by ";"
package token

func (r *Reader) Marks(sentence string) []string {
	var out []string
	for _, u := range unitsOf(r.ix, sentence, false) {
		sc := unitScan(r.ix, u)
		out = append(out, string(sc.s)+"\t"+sc.marksOf(sc.search()))
	}
	return out
}
