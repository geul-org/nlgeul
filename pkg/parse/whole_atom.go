//ff:func feature=parse type=reader control=sequence
//ff:what wholeAtom — note the atom a unit begins with: the sign of a pair or the span of code at its head, or, when the unit is no sentence's word, an address, a mail address, a path or an option: one piece, not read inside, named by what it is; where it ends (0 when the unit is none of these); and where a path that ends in a slash is cut from the Hangul after it (pathCut), or, when the unit begins with none of these, where it ends as a path with no head and one slash (weakPath)
package parse

func (sc *scan) wholeAtom() int {
	name, end := special(sc.s)
	if sc.head > 0 {
		name, end = sc.headName, sc.head
	}
	if end == 0 {
		sc.weak = weakPath(sc.s)
		return 0
	}
	sc.atoms[0] = atom{end, []string{name}, false}
	sc.fill(0, end)
	sc.cut = pathCut(sc.s, name, end)
	return end
}
