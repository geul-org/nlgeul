//ff:func feature=token type=reader control=sequence
//ff:what opened — the ways the signs may stand open after the sign of a pair at place i, and whether it may be read as a unit of measure (a sign of the unit table right after a digit: 3", 5'): then the ways before it stay too, as a unit opens and closes nothing
package token

import (
	"slices"
	"unicode"
)

func (c *cutter) opened(i int, after []string) ([]string, bool) {
	if i == 0 || !unicode.IsDigit(c.r[i-1]) || !slices.Contains(unitWords, string(c.r[i])) {
		return after, false
	}
	return distinct(append(append([]string{}, after...), c.open...)), true
}
