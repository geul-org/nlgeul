//ff:func feature=token type=helper control=iteration dimension=2
//ff:what edges — the pairs of spans where a word at the second can follow a word at the first, in the order of the spans
package token

func (w *weave) edges() []edge {
	var out []edge
	for i, ids := range w.succ {
		for _, j := range ids {
			out = append(out, edge{from: w.spans[i], to: w.spans[j]})
		}
	}
	return out
}
