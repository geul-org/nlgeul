//ff:func feature=parse type=helper control=sequence
//ff:what stemlessAfter — whether a whole form right after the letter given has no stem before it: after a middle dot (a sign between two words) no form has; after a sign of a pair or another punctuation mark a slot has none, for a slot stands only on a predicate's stem (70개(26.2%)였다 is the copula 이 with 었, not the slot 였); after a digit a slot has none either; a closing is read there (closingTails marks it as a tail no stem takes, so only the copula read on it stands: 3+(이)+야, 3+(이)+라면)
package parse

import (
	"slices"
	"unicode"
)

// The dictionary gives the slot -였- after the stem of 하다 only, and
// every slot after a predicate's stem; what a closing sign ends is no
// such stem. A number is a nominal: a slot or a closing follows it on
// the copula's stem only (3이면, 3+(이)+야), so none stands right on its
// last digit (3면 is no 3+면, 제4면 no 제4+면); a closing after a digit is
// kept as a tail that no stem takes (closingTails), so that only the
// copula's read of it, with the stem dropped, stands.
func stemlessAfter(e *entry, prev rune) bool {
	if slices.Contains(signNames[prev], "가운뎃점") {
		return true
	}
	if unicode.IsDigit(prev) && e.class == classSlot {
		return true
	}
	_, pair := pairRows[prev]
	return e.class == classSlot && (pair || namedSign(prev))
}
