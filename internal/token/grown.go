//ff:func feature=token type=reader control=iteration dimension=1
//ff:what grown — the readings kept at b, each grown by the word from b to j that tail t makes, when the word before the tail passes the tail's graph and is no new word whose letters read as markers there, no dropped ㅡ is put back where the stem does not drop it, and no word the tables know only as a nominal stands as a new word before a predicate tail; a tail read only after a modifying closing grows only the readings whose last word ends in one, and where a word that ends in a modifying closing ends is kept
package token

func (sc *scan) grown(cands []*path, kept []*path, t tail, b, j int) []*path {
	if !sc.gated(b, t) || sc.markerStem(b, j, t) || sc.euDropped(b, t) || sc.knownAsNew(b, t) {
		return cands
	}
	e := sc.wordOf(b, j, t)
	for _, p := range kept {
		if t.adn && !p.last.adn {
			continue
		}
		cands = append(cands, grow(p, e))
		sc.adnEnds[j] = sc.adnEnds[j] || e.adn
	}
	return cands
}
