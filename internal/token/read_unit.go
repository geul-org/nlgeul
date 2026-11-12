//ff:func feature=token type=reader control=sequence
//ff:what readUnit — one unit (the text between fixed boundaries: spaces and the signs of a pair) read into bundles: the signs before its letters, the stretches its candidates agree on (a sign of a pair at its head once for each way it is read: spread, with the candidates as a unit of measure after a digit: measured; the signs between letters read as words of their own: expandMid), then the signs after its letters; and its trace when asked for; a sign of a pair with letters written onto it is read in two (readOnto)
package token

// A unit with nothing to read (signs only) is the bundles of its
// signs.
func readUnit(ix *index, u unit, trace bool) ([]Bundle, []string) {
	if u.name == pairName && u.head < len([]rune(u.text)) {
		return readOnto(ix, u, trace)
	}
	sc := unitScan(ix, u)
	var notes []string
	if trace {
		notes = append(notes, u.notes...)
		sc.trace = &notes
	}
	lead, trail, at := edgesOf(u, sc.s)
	out := sc.signBundles(u, lead, 0, false)
	out = append(out, ix.measured(sc.expandMid(spread(sc.bundles(sc.search()), u)), u)...)
	return append(out, sc.signBundles(u, trail, at, true)...), notes
}
