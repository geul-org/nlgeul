//ff:func feature=token type=reader control=iteration dimension=1
//ff:what unitsOf — cut a sentence into the units that are read one by one: at spaces, and at the signs that open and close a pair (brackets, quotation marks), wherever in the sentence the two stand; a span of code between backticks is one unit with the spaces inside it
package token

import "unicode"

// A sign with no partner is left in its unit, where the edges of a
// unit are trimmed as before.
func unitsOf(sentence string) []unit {
	r := []rune(sentence)
	pairs := pairsOf(r)
	var out []unit
	for i := 0; i < len(r); {
		if unicode.IsSpace(r[i]) {
			i++
			continue
		}
		u, next := nextUnit(r, i, pairs)
		out, i = append(out, u), next
	}
	return out
}
