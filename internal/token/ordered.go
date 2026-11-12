//ff:func feature=token type=helper control=iteration dimension=1
//ff:what ordered — the candidates of a sign of a pair, each once: those that open before those that close, each kind from the least depth up
package token

import "sort"

func ordered(reads []pairing) []pairing {
	var out []pairing
	seen := map[pairing]bool{}
	for _, rd := range reads {
		if !seen[rd] {
			seen[rd] = true
			out = append(out, rd)
		}
	}
	sort.SliceStable(out, func(a, b int) bool {
		if out[a].name != out[b].name {
			return out[a].name == openName
		}
		return out[a].depth < out[b].depth
	})
	return out
}
