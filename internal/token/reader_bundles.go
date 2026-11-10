//ff:func feature=token type=reader control=iteration dimension=2
//ff:what Bundles — read a sentence and send up its words as a structure: the sentence is cut into units at spaces and at the signs of a pair, and each unit comes back as bundles in order, each marked with which unit it is in
package token

// A layer above reads this instead of the printed lines: each bundle
// holds its branches, each branch its words, each word its parts with
// their readings. Lines prints the same bundles.
func (r *Reader) Bundles(sentence string) []Bundle {
	var out []Bundle
	for i, u := range unitsOf(sentence) {
		bundles, _ := readUnit(r.ix, u, false)
		for _, b := range bundles {
			b.Unit = i
			out = append(out, b)
		}
	}
	return out
}
