//ff:func feature=token type=reader control=sequence
//ff:what ontoReader — begin to read a unit whose head is a sign of a pair with letters written onto it, as readOntoNet reads it (the contests of the sign alone, the letters after it by the net beside the unit as a whole, the signs after the letters), the letters worked out as stretches are asked for; the words of the whole unit that the letters may begin with are planted while no cut has been made (planter), and the letters are read as the letters of a unit with no head
package token

func ontoReader(ix *index, u unit, trace bool, stats *netStats) *unitReader {
	su := u
	su.text = string([]rune(u.text)[:u.head])
	contests, notes := readUnit(ix, su, trace)
	r := &unitReader{queue: stretchesOf(contests), notes: notes}
	sc := unitScan(ix, u)
	if trace {
		sc.trace = &r.notes
	}
	_, trail, at := edgesOf(u, sc.s)
	r.after = func() []unitStretch { return stretchesOf(sc.signContests(u, trail, at, true)) }
	if len(sc.s) > u.head {
		ps := &scan{ix: sc.ix, s: sc.s[u.head:], closed: sc.closed, trace: sc.trace, reg: newRegistry(ontoBorn)}
		pl := newPlanter(sc, u.head)
		plant := func(j int) []*path {
			pl.step(j + u.head)
			return pl.at(j+u.head, ps)
		}
		r.run = newNetRun(ps, u, false, plant, stats)
	}
	return r
}
