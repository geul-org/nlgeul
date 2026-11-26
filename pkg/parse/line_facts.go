//ff:func feature=parse type=reader control=sequence
//ff:what lineFacts — the facts of a unit's place that its own letters do not show, for the unit from place i to end of a sentence: its input line holds no letter and no digit (MarkBare), a space or the sentence's edge stands on both sides of it (MarkApart), it is one hyphen alone where a sum is written (MarkSum)
package parse

import (
	"unicode"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (c *cutter) lineFacts(i, end int) int {
	facts := 0
	if c.lineBare(i) {
		facts |= rule.MarkBare
	}
	if (i == 0 || unicode.IsSpace(c.r[i-1])) && (end >= len(c.r) || unicode.IsSpace(c.r[end])) {
		facts |= rule.MarkApart
	}
	if c.sumPlace(i, end) {
		facts |= rule.MarkSum
	}
	return facts
}
