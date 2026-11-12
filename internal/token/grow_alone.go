//ff:func feature=token type=reader control=iteration dimension=2
//ff:what growAlone — the readings kept at b, each grown by one word from b to j of no markers; a word of the word table that only stands alone grows none of the readings whose last word ends in a modifying closing or is a determiner (afterAdn, for the kind the stem stands as), and where a determiner ends is kept
package token

func (sc *scan) growAlone(cands []*path, kept []*path, e step, b, j int) []*path {
	adn, det := sc.afterAdn(b, tail{start: j}, false, e.kind), sc.afterAdn(b, tail{start: j}, true, e.kind)
	reads := sc.numberReads(e, b, j)
	for _, p := range kept {
		if (adn && p.last.adn) || (det && p.last.det) {
			continue
		}
		for _, r := range reads {
			cands = append(cands, grow(p, r))
		}
		sc.detEnds[j] = sc.detEnds[j] || e.det
	}
	return cands
}
