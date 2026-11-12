//ff:func feature=token type=reader control=sequence
//ff:what closingDots — the dots that close a unit, right after its core, read as a bundle of their own (올렸다. → .[마침표:.], 글쎄... → ...[줄임표:.]); nothing when the unit does not end so; the core is looked for after the unit's head (a sign of a pair or a span of code), as a sign dropped between the head and the core leaves the two apart in the unit's letters
package token

import "strings"

// Nothing stands after such dots, so one of the two closing readings
// always holds: a single dot is 마침표, dots in a row 줄임표.
func (sc *scan) closingDots(field string) []Bundle {
	head, body := string(sc.s[:sc.head]), string(sc.s[sc.head:])
	at := len(head) + strings.Index(field[len(head):], body) + len(body)
	rest := field[at:]
	dots := rest[:len(rest)-len(strings.TrimLeft(rest, "."))]
	if dots == "" {
		return nil
	}
	m := sc.stopFacts(len(sc.s), len(sc.s))
	m.Run = len(dots)
	names, _ := sc.judgeMark(sc.ix.stops, m, dots)
	return []Bundle{{Text: dots, Branches: []Branch{{Words: []Eojeol{{Parts: []Part{{Text: dots, Names: namesOf(names)}}, End: len(sc.s)}}, Verdict: 1}}}}
}
