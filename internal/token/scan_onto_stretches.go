//ff:func feature=token type=reader control=iteration dimension=1
//ff:what ontoStretches — the stretches of the letters written onto the sign of a pair at a unit's head, read by the net: the letters after the sign are read by a scan of their own whose words are numbered on from ontoBorn, and the whole unit is read beside them by a scan that makes only the words that begin at its head or after the sign (planter), place by place, the words it makes that the letters may begin with (the first word with the sign taken off, the words after the sign alone) planted in the net of the letters at the place they end; the stretches are those of the live states, as netStretches cuts them, but the letters are read as the letters of a unit with no head (no spread, no unit of measure)
package token

func (sc *scan) ontoStretches(u unit) []unitStretch {
	ps := &scan{ix: sc.ix, s: sc.s[u.head:], closed: sc.closed, trace: sc.trace, reg: newRegistry(ontoBorn)}
	pl := newPlanter(sc, u.head)
	best := ps.netStart()
	for j := 1; j <= len(ps.s); j++ {
		pl.step(j + u.head)
		ps.netStep(best, j, pl.at(j+u.head, ps))
	}
	var out []unitStretch
	for _, pt := range ps.netEnd(best).portions() {
		out = append(out, ps.stretchOf(u, pt, false))
	}
	return out
}
