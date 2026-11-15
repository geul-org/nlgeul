//ff:func feature=token type=helper control=sequence
//ff:what stemlessAfter — whether a whole form right after the letter given has no stem before it: after a middle dot (a sign between two words) no form has; after a sign of a pair or another punctuation mark a slot has none, for a slot stands only on a predicate's stem (70개(26.2%)였다 is the copula 이 with 었, not the slot 였)
package token

import "slices"

// The dictionary gives the slot -였- after the stem of 하다 only, and
// every slot after a predicate's stem; what a closing sign ends is no
// such stem.
func stemlessAfter(e *entry, prev rune) bool {
	if slices.Contains(signNames[prev], "가운뎃점") {
		return true
	}
	_, pair := pairRows[prev]
	return e.class == classSlot && (pair || len(signNames[prev]) > 0)
}
