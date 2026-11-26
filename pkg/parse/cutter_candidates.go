//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what candidates — the candidates of the sign at place i as a sign of a pair, for every way the signs before it may stand open: each candidate that stands (열기 or 닫기, with its depth) once, whether the sign is also read as its letter alone (it is unsure, in some way no candidate stands, or it stands as a round bracket that numbers an item of a list: listNumber, a closing at depth 0), the ways the signs may stand open after it, and the trace lines; nothing when the letter is no sign of a pair, stands in a span of code, or stands as such a sign in no way; every way is judged knowing whether a sign of its kind stands open in any of them (MarkOpenedAny) and whether a sign written apart has its partner written apart too (MarkMated); the ways are taken in the order leaned gives (a quotation mark written the same way at both ends is read first as the candidate its sides lean to)
package parse

import (
	"slices"
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The cutter is not changed: next moves it when the sign is cut off.
func (c *cutter) candidates(i int) ([]pairing, bool, []string, []string) {
	row, ok := pairRows[c.r[i]]
	if !ok || c.inside[i] {
		return nil, false, nil, nil
	}
	var candidates []pairing
	var open, notes []string
	bare := false
	m := c.facts(i)
	if c.mated(i, row) {
		m.Facts |= rule.MarkMated
	}
	if strings.ContainsRune(strings.Join(c.open, ""), row.kind) {
		m.Facts |= rule.MarkOpenedAny
	}
	for _, way := range c.leaned(i, row) {
		stood, after, lines := c.judge(i, row, m, way)
		candidates, open, notes = append(candidates, stood...), append(open, after...), append(notes, lines...)
		if len(stood) == 0 || row.unsure || (listNumber(c.r[i], m) && slices.Contains(stood, pairing{shutName, 0})) {
			open, bare = append(open, way), true
		}
	}
	if len(candidates) == 0 {
		return nil, false, nil, nil
	}
	return distinct(candidates), bare, distinct(open), distinct(notes)
}
