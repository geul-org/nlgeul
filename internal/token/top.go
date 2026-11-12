//ff:func feature=token type=reader control=iteration dimension=1
//ff:what top — the readings in order, best first, the same reading only once (none is left out for its place)
package token

import "sort"

func (ix *index) top(cands []*path) []*path {
	sort.SliceStable(cands, func(a, b int) bool { return ix.before(cands[a], cands[b]) })
	seen := map[string]bool{}
	var out []*path
	for _, p := range cands {
		if !seen[p.sig] {
			seen[p.sig] = true
			out = append(out, p)
		}
	}
	return out
}
