//ff:func feature=token type=reader control=sequence
//ff:what readUnitNet — one unit read as readUnit reads it (the signs before its letters, the stretches its ways agree on, the signs after, and its trace), but the letters by the net (netStretches): every stretch is read once, however many ways there are to cut it; the signs at the edges stay contests, and a sign of a pair with letters written onto it is read the old way (readOnto), as is anything readUnit gives
package token

func readUnitNet(ix *index, u unit, trace bool) ([]unitStretch, []string) {
	if u.name == pairName && u.head < len([]rune(u.text)) {
		contests, notes := readUnit(ix, u, trace)
		return stretchesOf(contests), notes
	}
	sc := unitScan(ix, u)
	var notes []string
	if trace {
		notes = append(notes, u.notes...)
		sc.trace = &notes
	}
	lead, trail, at := edgesOf(u, sc.s)
	out := stretchesOf(sc.signContests(u, lead, 0, false))
	out = append(out, sc.netStretches(u)...)
	return append(out, stretchesOf(sc.signContests(u, trail, at, true))...), notes
}
