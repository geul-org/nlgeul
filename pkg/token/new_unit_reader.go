//ff:func feature=token type=reader control=sequence
//ff:what newUnitReader — begin to read one unit as readUnitNet reads it (the signs before its letters, the letters by the net, the signs after, and its trace when asked for), but the letters are worked out only as stretches are asked for (next); a sign of a pair with letters written onto it is read in two (ontoReader)
package token

func newUnitReader(ix *index, u unit, trace bool, stats *netStats) *unitReader {
	if u.name == pairName && u.head < len([]rune(u.text)) {
		return ontoReader(ix, u, trace, stats)
	}
	r := &unitReader{}
	sc := unitScan(ix, u)
	if trace {
		r.notes = append(r.notes, u.notes...)
		sc.trace = &r.notes
	}
	lead, trail, at := edgesOf(u, sc.s)
	r.queue = stretchesOf(sc.signContests(u, lead, 0, false))
	r.run = newNetRun(sc, u, true, nil, stats)
	r.after = func() []unitStretch { return stretchesOf(sc.signContests(u, trail, at, true)) }
	return r
}
