//ff:func feature=token type=reader control=iteration dimension=1
//ff:what growTypo — the candidates kept at b, each grown by one word from b to j of no markers written wrong, as a known word of no markers is grown (the candidates that end in a modifying closing or a determiner are not grown by a word that only stands alone)
package token

func (sc *scan) growTypo(cands []*path, kept []*path, t tail, e step, b, j int) []*path {
	adn, det := sc.afterAdn(b, t, false, e.kind), sc.afterAdn(b, t, true, e.kind)
	for _, p := range kept {
		if e.verdict <= 0 || (adn && p.last.adn) || (det && p.last.det) {
			continue
		}
		cands = append(cands, grow(p, e))
		sc.detEnds[j] = sc.detEnds[j] || e.det
	}
	return cands
}
