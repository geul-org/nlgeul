//ff:func feature=token type=reader control=iteration dimension=1
//ff:what top — the k best of the readings, the same reading only once
package token

import "sort"

func (ix *index) top(cands []*path, k int) []*path {
	sort.SliceStable(cands, func(a, b int) bool { return ix.before(cands[a], cands[b]) })
	seen := map[string]bool{}
	var out []*path
	for _, p := range cands {
		if len(out) < k && !seen[p.sig] {
			seen[p.sig] = true
			out = append(out, p)
		}
	}
	return out
}
