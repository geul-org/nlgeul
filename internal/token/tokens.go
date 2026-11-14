//ff:func feature=token type=reader control=iteration dimension=1
//ff:what tokens — how many tokens a stem from b to end holds for preference: an atom counts as one (from the counts of a scan that was prepared, else counted)
package token

func (sc *scan) tokens(b, end int) int {
	if end <= b {
		return 0
	}
	if len(sc.count) == len(sc.s)+1 && end < len(sc.count) {
		return sc.count[end] - sc.count[b]
	}
	n := 0
	for p := b + 1; p <= end; p++ {
		if !sc.inside[p] {
			n++
		}
	}
	return n
}
