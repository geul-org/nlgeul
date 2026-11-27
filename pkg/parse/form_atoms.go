//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what formAtoms — each letter of the unit that the tables of special letters hold and that is read as a letter of its own (formLetter) is an atom of one letter: a letter that writes a number (numberLetter: ①, ⒈, ½) a number, as a run of digits is (numEnds stands after it; its name is the table's, letterName), any other (㊀, ㈠) an atom with the names the tables give it
package parse

func (sc *scan) formAtoms() {
	for i, r := range sc.s {
		switch {
		case numberLetter(r):
			sc.closeAtom(i, i+1, nil, true)
		case formLetter(r):
			sc.closeAtom(i, i+1, formNames(r), false)
		}
	}
}
