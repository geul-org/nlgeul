//ff:func feature=mini type=reader control=iteration dimension=1
//ff:what readChain — the longest particle chain at the end of s
package mini

func readChain(s []rune) chainFact {
	var best chainFact
	for _, p := range particles {
		start, ok := matchParticle(s, len(s), p)
		if !ok {
			continue
		}
		c, inner, left := chainStart(s, start, 1, p.slot)
		if n := 2 * (len(s) - c); n > best.size {
			best = chainFact{n, left + p.form, p.form, p.count, outer(p.kase, inner)}
		}
	}
	return best
}
