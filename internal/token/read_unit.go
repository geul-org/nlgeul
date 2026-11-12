//ff:func feature=token type=reader control=sequence
//ff:what readUnit — one unit (the text between fixed boundaries: spaces and the signs of a pair) read into bundles: the stretches its readings agree on, then the dots that close it; and its trace when asked for
package token

// A unit with nothing to read (signs only) is one bundle with its
// letters as they are.
func readUnit(ix *index, u unit, trace bool) ([]Bundle, []string) {
	sc := unitScan(ix, u)
	if len(sc.s) == 0 {
		return []Bundle{{Text: u.text, Branches: []Branch{{Words: []Eojeol{{Parts: []Part{{Text: u.text, Stem: true}}}}, Verdict: 1}}}}, nil
	}
	var notes []string
	if trace {
		sc.trace = &notes
	}
	return append(sc.bundles(sc.search()), sc.closingDots(u.text)...), notes
}
