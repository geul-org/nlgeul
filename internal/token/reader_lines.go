//ff:func feature=token type=reader control=iteration dimension=2
//ff:what Lines — read a token stream and print what is sent up: each unit's trace when asked for, then its bundles one after another (render)
package token

// A unit is read without looking outside it, so the same unit always
// gives the same lines. The bundles are those Bundles gives.
func (r *Reader) Lines(sentence string, trace bool) []string {
	var out []string
	for _, u := range unitsOf(r.ix, sentence, trace) {
		bundles, notes := readUnit(r.ix, u, trace)
		out = append(out, notes...)
		for _, b := range bundles {
			out = append(out, render(b)...)
		}
	}
	return out
}
