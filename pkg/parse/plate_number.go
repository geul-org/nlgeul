//ff:func feature=parse type=reader control=sequence
//ff:what plateNumber — a motor vehicle's registration number that the unit's letters begin with (plateEnd) is one atom named plateKind: no word begins or ends within it, it stands as the stem of one word with the verdict of an atom, and what follows it is read after it (the unit 호, particles); where the same letters are a court case number too (a year of two digits, a mark of one syllable and a serial of four digits: 97도2836) the atom keeps that name first and takes this one after it, so that both candidates stand; it is sent up as its pieces (plateWord)
package parse

import "slices"

func (sc *scan) plateNumber() {
	end := sc.plateEnd(sc.head)
	if end == 0 {
		return
	}
	names := []string{plateKind}
	if a, ok := sc.atoms[sc.head]; ok && a.end == end && slices.Contains(a.names, caseKind) {
		names = []string{caseKind, plateKind}
	}
	sc.atoms[sc.head] = atom{end, names, false}
	sc.fill(sc.head, end)
}
