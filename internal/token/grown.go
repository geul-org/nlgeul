//ff:func feature=token type=reader control=iteration dimension=1
//ff:what grown — the readings kept at b, each grown by the word from b to j that tail t makes, when the word before the tail passes the tail's graph, the stem's last syllable is one it may have (restored), and the stem stands (a new word as its graph judges it); a tail read only after a modifying closing grows only the readings on which its graph stands (again), a predicate of the word table grows none of the readings whose last word ends in one or is a determiner (afterAdn), and where a word that ends in a modifying closing ends is kept
package token

func (sc *scan) grown(cands []*path, kept []*path, t tail, b, j int) []*path {
	if !sc.gated(b, t) || !sc.restored(b, t) {
		return cands
	}
	e, adn, det := sc.wordOf(b, j, t), sc.afterAdn(b, t, false), sc.afterAdn(b, t, true)
	for _, p := range kept {
		if e.verdict <= 0 || !sc.again(t, p.last.adn) || (adn && p.last.adn) || (det && p.last.det) {
			continue
		}
		cands = append(cands, grow(p, e))
		sc.adnEnds[j] = sc.adnEnds[j] || e.adn
	}
	return cands
}
