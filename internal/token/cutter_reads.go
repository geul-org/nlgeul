//ff:func feature=token type=reader control=iteration dimension=1
//ff:what reads — the readings of the sign at place i as a sign of a pair, for every way the signs before it may stand open: each reading that stands (열기 or 닫기, with its depth) once, whether the sign is also read as its letter alone (it is unsure, or in some way no reading stands), the ways the signs may stand open after it, and the trace lines; nothing when the letter is no sign of a pair, stands in a span of code, or stands as such a sign in no way
package token

// The cutter is not changed: next moves it when the sign is cut off.
func (c *cutter) reads(i int) ([]pairing, bool, []string, []string) {
	row, ok := pairRows[c.r[i]]
	if !ok || c.inside[i] {
		return nil, false, nil, nil
	}
	var reads []pairing
	var open, notes []string
	bare := false
	m := c.facts(i)
	for _, way := range c.open {
		stood, after, lines := c.judge(i, row, m, way)
		reads, open, notes = append(reads, stood...), append(open, after...), append(notes, lines...)
		if len(stood) == 0 || row.unsure {
			open, bare = append(open, way), true
		}
	}
	if len(reads) == 0 {
		return nil, false, nil, nil
	}
	return ordered(reads), bare, distinct(open), distinct(notes)
}
