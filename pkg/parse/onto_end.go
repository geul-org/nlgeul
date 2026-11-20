//ff:func feature=parse type=helper control=sequence
//ff:what ontoEnd — where what is written onto a closing sign or a span of code ends: the letters from a place up to the next space, sign of a pair or span of code; the place itself when a sign (a punctuation mark, a symbol, or a named sign that is a letter for Unicode: ㆍ), a space or nothing stands there (a sign after it is a unit of its own, its edge read as trimSigns reads one)
package parse

import "unicode"

func (c *cutter) ontoEnd(from int) int {
	if from >= len(c.r) || unicode.IsSpace(c.r[from]) || unicode.IsPunct(c.r[from]) || unicode.IsSymbol(c.r[from]) || len(signNames[c.r[from]]) > 0 {
		return from
	}
	return c.wordEnd(from)
}
