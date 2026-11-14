//ff:func feature=token type=reader control=iteration dimension=1
//ff:what unitsOf — cut a sentence into the units that are read one by one, from its first letter on: at spaces, and at the signs that open and close a pair (brackets, quotation marks), each read for every way the signs before it may stand open; a span of code between backticks is one unit with the spaces inside it
package token

import "unicode"

// Nothing is matched ahead: a sign of a pair is cut off where it
// stands, and a closing sign that none may stand open for is left in
// its unit as a letter.
func unitsOf(ix *index, sentence string, trace bool) []unit {
	c := newCutter(ix, sentence, trace)
	var out []unit
	for i := 0; i < len(c.r); {
		if unicode.IsSpace(c.r[i]) {
			i++
			continue
		}
		u, next := c.next(i)
		out, i = append(out, u), next
	}
	return out
}
