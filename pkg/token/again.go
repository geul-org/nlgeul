//ff:func feature=token type=rule control=sequence
//ff:what again — does a tail that was judged by the place it stands at stand on one candidate: its graph is asked once more with the fact of that candidate, whether the candidate's last word ends in a modifying closing (게, 건, 걸 as 거 with a particle are read only right after one); a tail with no such graph stands on any candidate
package token

func (sc *scan) again(t tail, adn bool) bool {
	if t.facts == nil {
		return true
	}
	p := *t.facts
	p.Adn = adn
	return sc.ix.judged(t.again.g, p, "") > 0
}
