//ff:func feature=parse type=helper control=sequence
//ff:what unitSign — whether the letter at place k is a sign of the unit table written right after a digit (the % of 90%, the ℃ of 3℃): it is read with its number, a piece of the word and no sign standing by itself
package parse

import (
	"slices"
	"unicode"
)

func unitSign(all []rune, k int) bool {
	if k < 1 || k >= len(all) || !signRune(all[k]) || !unicode.IsDigit(all[k-1]) {
		return false
	}
	return slices.Contains(unitWords, string(all[k]))
}
