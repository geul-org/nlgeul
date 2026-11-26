//ff:func feature=parse type=reader control=sequence
//ff:what sumPlace — whether the unit from place i to end of a sentence is one hyphen alone, a space on each side, where a sum is written: a bracket stands open before it in the first way the signs may stand open (sumBrackets: 가 × [나 - 다]), or a number ends right before the space before it (a digit, or a word of the unit table right after a digit) and a digit stands right after the space after it (100 - 20, 3,000원 - 1,000원)
package parse

import (
	"strings"
	"unicode"
)

func (c *cutter) sumPlace(i, end int) bool {
	if end != i+1 || c.r[i] != '-' || i == 0 || end >= len(c.r) || !unicode.IsSpace(c.r[i-1]) || !unicode.IsSpace(c.r[end]) {
		return false
	}
	if len(c.open) > 0 && strings.ContainsAny(c.open[0], sumBrackets) {
		return true
	}
	return amountEnd(c.r[:i-1]) && end+1 < len(c.r) && unicode.IsDigit(c.r[end+1])
}
