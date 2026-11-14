//ff:func feature=token type=helper control=iteration dimension=1
//ff:what endingAt — the states that end at place q
package token

func endingAt(order []*path, q int) map[*path]bool {
	out := map[*path]bool{}
	for _, p := range order {
		if p.last.end == q {
			out[p] = true
		}
	}
	return out
}
