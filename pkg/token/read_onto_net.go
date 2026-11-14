//ff:func feature=token type=reader control=sequence
//ff:what readOntoNet — a unit whose head is a sign of a pair with letters written onto it, read as readOnto reads it (the contests of the sign alone, the letters after it, the signs after the letters), but the letters by the net (ontoStretches): the sign and the signs after the letters stay contests, and the letters after the sign are the stretches of the net, with the first word that took the sign carried over to them (+를[목적격:를])
package token

func readOntoNet(ix *index, u unit, trace bool) ([]unitStretch, []string) {
	su := u
	su.text = string([]rune(u.text)[:u.head])
	contests, notes := readUnit(ix, su, trace)
	out := stretchesOf(contests)
	sc := unitScan(ix, u)
	if trace {
		sc.trace = &notes
	}
	_, trail, at := edgesOf(u, sc.s)
	if len(sc.s) > u.head {
		out = append(out, sc.ontoStretches(u)...)
	}
	return append(out, stretchesOf(sc.signContests(u, trail, at, true))...), notes
}
