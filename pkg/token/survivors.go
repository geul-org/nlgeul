//ff:func feature=token type=helper control=iteration dimension=1
//ff:what survivors — the states that stay when the net is kept from place q on standing on the roots (states that end at q): a state that begins at q stays when a root is before it, a state that begins after q when a state that stays is before it (reaches); the states are given in the order of where they end
package token

func survivors(order []*path, q int, roots map[*path]bool) map[*path]bool {
	alive := map[*path]bool{}
	for _, p := range order {
		if p.start >= q && reaches(p, q, roots, alive) {
			alive[p] = true
		}
	}
	return alive
}
