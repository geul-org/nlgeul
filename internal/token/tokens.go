//ff:func feature=token type=reader control=iteration dimension=1
//ff:what tokens — how many tokens a stem from b to end holds for preference: an atom counts as one
package token

func (sc *scan) tokens(b, end int) int {
	n := 0
	for p := b + 1; p <= end; p++ {
		if !sc.inside[p] {
			n++
		}
	}
	return n
}
