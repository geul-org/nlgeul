//ff:func feature=token type=reader control=iteration dimension=2
//ff:what Lines — read a token stream and print what is sent up: each unit's trace when asked for, then its bundles one after another (render); the readings asleep are printed only when the reader is set to
package token

// A unit is read without looking outside it, so the same unit always
// gives the same lines. The bundles are those Bundles gives.
func (r *Reader) Lines(sentence string, trace bool) []string {
	var out []string
	for _, u := range unitsOf(sentence) {
		bundles, notes := readUnit(r.ix, u, trace)
		out = append(out, notes...)
		for _, b := range bundles {
			if r.ix.asleep || !b.Asleep() {
				out = append(out, render(b)...)
			}
		}
	}
	return out
}
