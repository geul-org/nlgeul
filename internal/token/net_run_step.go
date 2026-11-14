//ff:func feature=token type=reader control=sequence
//ff:what step — work out the next place of the unit (netStep), with the words planted there while no cut has been made (the letters after a sign of a pair: only the words that begin at the sign can be planted), and note the most states kept at once
package token

func (r *netRun) step() {
	r.j++
	var planted []*path
	if r.plant != nil && r.c == 0 {
		planted = r.plant(r.j)
	}
	r.sc.netStep(r.best, r.j, planted)
	if k := len(r.sc.registry().all); k > r.stats.states {
		r.stats.states = k
	}
}
