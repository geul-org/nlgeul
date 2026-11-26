//ff:func feature=parse type=reader control=sequence
//ff:what readUnit — one unit (the text between fixed boundaries: spaces and the signs of a pair) read into contests: the signs before its letters, the stretches its candidates agree on (a sign of a pair at its head once for each way it is read: spread, with the candidates as a unit of measure after a digit: measured; the signs between letters read as words of their own: expandMid; a court case number as its three pieces: caseNumberParts), then the signs after its letters; and its trace when asked for (a sign of a pair with letters written onto it is read by ontoReader, which gives it only the sign)
package parse

// A unit with nothing to read (signs only) is the contests of its
// signs.
func readUnit(ix *index, u unit, trace bool) ([]Contest, []string) {
	sc := unitScan(ix, u)
	var notes []string
	if trace {
		notes = append(notes, u.notes...)
		sc.trace = &notes
	}
	lead, trail, at := edgesOf(u, sc.s)
	out := sc.signContests(u, lead, 0, false)
	out = append(out, ix.measured(caseNumberParts(sc.expandMid(spread(sc.contests(sc.search()), u))), u)...)
	return append(out, sc.signContests(u, trail, at, true)...), notes
}
