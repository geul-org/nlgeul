//ff:func feature=token type=reader control=sequence
//ff:what wholeAtom — note the atom a unit begins with when the unit is no sentence's word but an address, a mail address, a path or an option: one piece, not read inside, named by what it is; where it ends (0 when the unit is none of these)
package token

func (sc *scan) wholeAtom() int {
	name, end := special(sc.s)
	if end == 0 {
		return 0
	}
	sc.atoms[0] = atom{end, name, false}
	sc.fill(0, end)
	return end
}
