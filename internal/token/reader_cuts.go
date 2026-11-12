//ff:func feature=token type=reader control=iteration dimension=1
//ff:what Cuts — for measuring: one line for each unit, its text, the word ends (token offsets) of each reading, best first, readings separated by ";"
package token

func (r *Reader) Cuts(sentence string) []string {
	var out []string
	for _, u := range unitsOf(sentence) {
		sc := &scan{ix: r.ix, s: coreOf(u), head: u.head, headName: u.name}
		out = append(out, string(sc.s)+"\t"+endsOf(sc.search()))
	}
	return out
}
