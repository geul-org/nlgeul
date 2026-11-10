//ff:func feature=token type=reader control=sequence
//ff:what closingDots — the dots that close a unit, right after its core, read as a line of their own (올렸다. → .[마침표:.], 글쎄... → ...[줄임표:.]); nothing when the unit does not end so
package token

import "strings"

// Nothing stands after such dots, so one of the two closing readings
// always holds: a single dot is 마침표, dots in a row 줄임표.
func (sc *scan) closingDots(field string) []string {
	at := strings.Index(field, string(sc.s)) + len(string(sc.s))
	rest := field[at:]
	dots := rest[:len(rest)-len(strings.TrimLeft(rest, "."))]
	if dots == "" {
		return nil
	}
	m := sc.stopFacts(len(sc.s), len(sc.s))
	m.Run = len(dots)
	names, _ := sc.judgeMark(sc.ix.stops, m, dots)
	return []string{dots + "[" + names + "]"}
}
