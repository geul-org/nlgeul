//ff:func feature=parse type=reader control=sequence
//ff:what caseNumber — a court case number that the unit's letters begin with (caseNumberEnd) is not read inside: no word begins or ends within it, so it stands as one word that no table lists (2004다25123), and what follows it is read after it (a unit, particles)
package parse

func (sc *scan) caseNumber() {
	if end := sc.caseNumberEnd(sc.head); end > 0 {
		sc.fill(sc.head, end)
	}
}
