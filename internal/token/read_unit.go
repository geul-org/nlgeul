//ff:func feature=token type=reader control=iteration dimension=1
//ff:what readUnit — the lines of one unit (the text between fixed boundaries): its trace when asked for, then the bundles of its readings that are awake, then (when asked for) the readings that are asleep, then the dots that close it
package token

func readUnit(ix *index, field string, trace bool) []string {
	sc := &scan{ix: ix, s: trimUnit(field)}
	if len(sc.s) == 0 {
		return []string{field}
	}
	var notes []string
	if trace {
		sc.trace = &notes
	}
	var out []string
	awake, asleep := wake(sc.search())
	for _, b := range sc.bundles(awake) {
		out = append(out, render(b)...)
	}
	if ix.asleep {
		out = append(out, renderAsleep(string(sc.s), asleep)...)
	}
	out = append(out, sc.closingDots(field)...)
	return append(notes, out...)
}
