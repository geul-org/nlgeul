//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what enclosedAtoms — each circled or parenthesized number of the unit (enclosedNumber) is an atom of its own, a number named 수 (numEnds stands after it), as a run of digits is
package parse

func (sc *scan) enclosedAtoms() {
	for i, r := range sc.s {
		if enclosedNumber(r) {
			sc.closeAtom(i, i+1, nil, true)
		}
	}
}
