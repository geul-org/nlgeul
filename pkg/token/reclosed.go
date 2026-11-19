//ff:func feature=token type=reader control=iteration dimension=1
//ff:what reclosed — the tails of the closings read at j once more as endings right after a final closing of a quoted sentence (closedBefore), in the order of groupKeys
package token

func (sc *scan) reclosed(pieces map[string]piece, from map[string]*entry, starts map[string]int, j int) []tail {
	var out []tail
	for _, k := range sc.groupKeys(pieces, from, j) {
		out = append(out, sc.closedBefore(pieces[k], from[k], starts[k], j)...)
	}
	return out
}
