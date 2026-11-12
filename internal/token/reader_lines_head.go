//ff:func feature=token type=reader control=iteration dimension=2
//ff:what LinesHead — read a token stream and print what is sent up: each unit's trace when asked for, then its bundles one after another (render), a contest printing only its first head candidates (0 prints all); the candidates are all made and Bundles gives all of them
package token

// head only limits the lines printed for a contest; the trace lines
// are never cut.
func (r *Reader) LinesHead(sentence string, trace bool, head int) []string {
	var out []string
	for _, u := range unitsOf(r.ix, sentence, trace) {
		bundles, notes := readUnit(r.ix, u, trace)
		out = append(out, notes...)
		for _, b := range bundles {
			out = append(out, render(b, head)...)
		}
	}
	return out
}
