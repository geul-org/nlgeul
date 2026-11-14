//ff:func feature=token type=helper control=iteration dimension=1
//ff:what relink — a state that stays after a cut at q has before it only what stays: the seed of each root before a state that begins at q, the states that stay before a state that begins after q; the first of them is the state before it (prev)
package token

func relink(p *path, q int, roots, alive map[*path]bool, seedOf map[*path]*path) {
	var preds []*path
	for _, b := range p.preds {
		switch {
		case p.start == q && roots[b]:
			preds = append(preds, seedOf[b])
		case p.start > q && alive[b]:
			preds = append(preds, b)
		}
	}
	p.preds, p.prev = preds, preds[0]
}
