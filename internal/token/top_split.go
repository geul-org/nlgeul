//ff:func feature=token type=reader control=iteration dimension=1
//ff:what topSplit — the k best readings with no letters written wrong and the k best with some, counted apart, put together in the order of before (the readings with none come out as top would give them, the same reading only once)
package token

import "sort"

// A reading with letters written wrong is one more reading of a place,
// never one that takes another reading's place among the k kept.
func (ix *index) topSplit(cands []*path, k int) []*path {
	var plain, typo []*path
	for _, p := range cands {
		if p.typo {
			typo = append(typo, p)
			continue
		}
		plain = append(plain, p)
	}
	out := append(ix.top(plain, k), ix.top(typo, k)...)
	sort.SliceStable(out, func(a, b int) bool { return ix.before(out[a], out[b]) })
	return out
}
