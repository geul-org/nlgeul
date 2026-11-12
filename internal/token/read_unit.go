//ff:func feature=token type=reader control=sequence
//ff:what readUnit — one unit (the text between fixed boundaries: spaces and the signs of a pair) read into bundles: the signs before its letters, the stretches its readings agree on (a sign of a pair at its head once for each way it is read: spread), then the signs after its letters; and its trace when asked for
package token

// A unit with nothing to read (signs only) is the bundles of its
// signs.
func readUnit(ix *index, u unit, trace bool) ([]Bundle, []string) {
	sc := unitScan(ix, u)
	var notes []string
	if trace {
		notes = append(notes, u.notes...)
		sc.trace = &notes
	}
	lead, trail, at := edgesOf(u, sc.s)
	out := sc.signBundles(u, lead, 0, false)
	out = append(out, spread(sc.bundles(sc.search()), u)...)
	return append(out, sc.signBundles(u, trail, at, true)...), notes
}
