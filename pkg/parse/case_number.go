//ff:func feature=parse type=reader control=sequence
//ff:what caseNumber — a court case number that the unit's letters begin with (caseNumberEnd) is one atom named caseKind: no word begins or ends within it, it stands as the stem of one word with the verdict of an atom (2004다25123), and what follows it is read after it (a unit, particles); it is sent up as its three pieces (caseNumberParts)
package parse

func (sc *scan) caseNumber() {
	if end := sc.caseNumberEnd(sc.head); end > 0 {
		sc.atoms[sc.head] = atom{end, []string{caseKind}, false}
		sc.fill(sc.head, end)
	}
}
