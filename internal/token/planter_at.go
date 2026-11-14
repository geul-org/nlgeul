//ff:func feature=token type=reader control=iteration dimension=1
//ff:what at — the states of the whole unit that end at place J as the letters after the sign take them (carried over, one word each, after the root of those letters): a first word with the sign taken off, and a word right after the sign alone; each ends in the places of the letters after the sign what a modifying closing or a determiner ends in the whole unit's; each keeps the number of birth it has
package token

func (pl *planter) at(J int, ps *scan) []*path {
	var out []*path
	root := ps.registry().root
	for _, q := range pl.best[J] {
		e, ok := pl.firstOf(q)
		if !ok {
			e, ok = pl.afterOf(q)
		}
		if !ok {
			continue
		}
		ps.adnEnds[e.end] = ps.adnEnds[e.end] || e.adn
		ps.detEnds[e.end] = ps.detEnds[e.end] || e.det
		out = append(out, grow(root, e))
	}
	return out
}
