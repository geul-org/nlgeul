//ff:func feature=token type=helper control=sequence
//ff:what ontoEnd — where what is written onto a closing sign or a span of code ends: the letters from a place up to the next space, sign of a pair or span of code; the place itself when a sign, a space or nothing stands there (a sign after it is a unit of its own)
package token

import "unicode"

func (c *cutter) ontoEnd(from int) int {
	if from >= len(c.r) || unicode.IsSpace(c.r[from]) || unicode.IsPunct(c.r[from]) || unicode.IsSymbol(c.r[from]) {
		return from
	}
	return c.wordEnd(from)
}
