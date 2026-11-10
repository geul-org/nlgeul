//ff:func feature=token type=reader control=iteration dimension=1
//ff:what Lines — read a token stream and send up its words: the sentence is cut into units at spaces and at the signs of a pair (fixed boundaries), and each unit comes back as bundles of words, one line each
package token

// A unit is read without looking outside it, so the same unit always
// gives the same lines.
func (r *Reader) Lines(sentence string, trace bool) []string {
	var out []string
	for _, u := range unitsOf(sentence) {
		out = append(out, readUnit(r.ix, u, trace)...)
	}
	return out
}
