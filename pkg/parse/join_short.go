//ff:func feature=parse type=reader control=sequence
//ff:what joinShort — read a shortened 것 (게, 건, 걸) at the word's end after a demonstrative (이게) or a modifying closing (된다는게), as a chain with the case of the particle it holds
package parse

import "slices"

// The chain's length counts what stands before the shortened form too
// (2 for a demonstrative).
func joinShort(s []rune) (chainFact, bool) {
	n := len(s)
	if n < 2 {
		return chainFact{}, false
	}
	form := string(s[n-1])
	kase, ok := shortened[form]
	left, size, adn := leftAdn(s[:n-1])
	if slices.Contains(pointers, string(s[:n-1])) {
		left, size, adn = string(s[:n-1]), 2, true
	}
	if !ok || !adn {
		return chainFact{}, false
	}
	return chainFact{size + 2, form, form, 0, kase, left + "+" + form}, true
}
