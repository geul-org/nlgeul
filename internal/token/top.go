//ff:func feature=token type=reader control=iteration dimension=1
//ff:what top — the k best of the readings: by verdict, then by preference; the same reading only once
package token

import "sort"

func top(cands []*path, k int) []*path {
	sort.SliceStable(cands, func(a, b int) bool {
		if cands[a].verdict != cands[b].verdict {
			return cands[a].verdict > cands[b].verdict
		}
		return cands[a].pref > cands[b].pref
	})
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
