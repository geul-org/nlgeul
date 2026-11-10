//ff:func feature=token type=reader control=iteration dimension=1
//ff:what readUnit — the lines of one unit (the text between fixed boundaries): its trace when asked for, then its bundles, then the dots that close it
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
	for _, b := range sc.bundles(sc.search()) {
		out = append(out, render(b)...)
	}
	out = append(out, sc.closingDots(field)...)
	return append(notes, out...)
}
