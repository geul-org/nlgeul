//ff:func feature=token type=helper control=iteration dimension=1
//ff:what spansAt — the spans of the nodes given by their places
package token

func (rg *region) spansAt(ids []int) []Span {
	out := make([]Span, 0, len(ids))
	for _, id := range ids {
		out = append(out, rg.nodes[id].Span)
	}
	return out
}
