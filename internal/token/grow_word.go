//ff:func feature=token type=reader control=iteration dimension=2
//ff:what growWord — the candidates kept at b, each grown by one word from b to j that tail t makes; a tail read only after a modifying closing grows only the candidates on which its graph stands (again), a predicate of the word table grows none of the candidates whose last word ends in one or is a determiner (afterAdn, for the kind the stem stands as), and where a word that ends in a modifying closing ends is kept
package token

func (sc *scan) growWord(cands []*path, kept []*path, t tail, e step, b, j int) []*path {
	adn, det := sc.afterAdn(b, t, false, e.kind), sc.afterAdn(b, t, true, e.kind)
	candidates := sc.numberCandidates(e, b, t.start)
	for _, p := range kept {
		if e.verdict <= 0 || !sc.again(t, p.last.adn) || (adn && p.last.adn) || (det && p.last.det) {
			continue
		}
		for _, r := range candidates {
			cands = append(cands, grow(p, r))
		}
		sc.adnEnds[j] = sc.adnEnds[j] || e.adn
	}
	return cands
}
