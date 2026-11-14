//ff:func feature=token type=reader control=sequence
//ff:what readOnto — a unit whose head is a sign of a pair with letters written onto it, read as the contests of the sign alone, then the contests of the letters: those read on their own (their words are numbered on from where the whole unit's were), with the candidates that took the sign as their first word moved over with the sign taken off (the first word then has no stem: +를[목적격:를]); then the signs after the letters
package token

// The sign's contests are read as a sign alone is. The candidates of the
// letters are those of a unit with no head, beside the candidates of the
// whole unit that have the sign as a word of its own; those that took
// the sign into a word's letters are not moved (the letters read alone
// stand for them).
func readOnto(ix *index, u unit, trace bool) ([]Contest, []string) {
	su := u
	su.text = string([]rune(u.text)[:u.head])
	out, notes := readUnit(ix, su, trace)
	sc := unitScan(ix, u)
	if trace {
		sc.trace = &notes
	}
	_, trail, at := edgesOf(u, sc.s)
	paths := sc.search()
	if len(sc.s) > u.head {
		ps := &scan{ix: ix, s: sc.s[u.head:], closed: sc.closed, trace: sc.trace, reg: newRegistry(sc.registry().next)}
		onto := ps.search()
		all := ix.top(append(sc.moved(paths, u.head), onto...))
		out = append(out, ps.expandMid(ps.contests(all))...)
	}
	return append(out, sc.signContests(u, trail, at, true)...), notes
}
