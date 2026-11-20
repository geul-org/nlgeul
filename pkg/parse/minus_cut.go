//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what minusCut — whether a hyphen at place i of a sentence is the minus sign of a sum written between two amounts with no space (3,000원-1,000원): a word of the unit table stands right before it after a digit, and a digit stands right after it; not inside a span of code nor inside a stretch that is no sentence's word (an address, a path, an option)
package parse

import "unicode"

func (c *cutter) minusCut(i int) bool {
	if c.r[i] != '-' || c.inside[i] || i+1 >= len(c.r) || !unicode.IsDigit(c.r[i+1]) || c.special(i) {
		return false
	}
	for _, u := range unitWords {
		k := i - len([]rune(u))
		if k >= 1 && string(c.r[k:i]) == u && unicode.IsDigit(c.r[k-1]) {
			return true
		}
	}
	return false
}
