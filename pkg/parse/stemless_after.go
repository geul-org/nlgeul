//ff:func feature=parse type=helper control=sequence
//ff:what stemlessAfter — whether a whole form right after the letter given has no stem before it: after a middle dot (a sign between two words) no form has; after a sign of a pair or another punctuation mark a slot has none, for a slot stands only on a predicate's stem (70개(26.2%)였다 is the copula 이 with 었, not the slot 였); after a digit a slot or a closing has none either, but for a closing that stands on the copula only, whose stem is read between (3+(이)+야)
package parse

import (
	"slices"
	"unicode"
)

// The dictionary gives the slot -였- after the stem of 하다 only, and
// every slot after a predicate's stem; what a closing sign ends is no
// such stem. A number is a nominal: a slot or a closing follows it on
// the copula's stem only (3이면, 3+(이)+야), so none stands right on its
// last digit (3면 is no 3+면, 제4면 no 제4+면).
func stemlessAfter(e *entry, prev rune) bool {
	if slices.Contains(signNames[prev], "가운뎃점") {
		return true
	}
	if unicode.IsDigit(prev) && (e.class == classSlot || (e.class == classClosing && !e.needs)) {
		return true
	}
	_, pair := pairRows[prev]
	return e.class == classSlot && (pair || len(signNames[prev]) > 0)
}
