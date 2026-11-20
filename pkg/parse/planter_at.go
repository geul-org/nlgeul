//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what at — the states of the whole unit that end at place J as the letters after the sign take them (carried over, one word each, after the root of those letters): a first word with the sign taken off, and a word right after the sign alone; each ends in the places of the letters after the sign what a modifying closing, a determiner or the linking 어/아 ends in the whole unit's; each keeps the number of birth it has
package parse

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
		ps.connEnds[e.end] = ps.connEnds[e.end] || e.conn
		out = append(out, grow(root, e))
	}
	return out
}
