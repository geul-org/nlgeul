//ff:func feature=token type=reader control=sequence
//ff:what stopFacts — the facts of the dots from start to end: what stands on each side (nothing at the unit's edges) and how many dots there are
package token

import "github.com/geul-org/koparse/pkg/rule"

func (sc *scan) stopFacts(start, end int) rule.Mark {
	m := rule.Mark{Run: end - start}
	if start > 0 {
		m.Left = sideOf(string(sc.s[start-1]))
	}
	if end < len(sc.s) {
		m.Right = sideOf(string(sc.s[end]))
	}
	return m
}
