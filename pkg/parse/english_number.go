//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what englishNumber — is a text a number written in English: a numeral of the table (englishNumerals: two, hundred), whatever the case of its letters (Two, TWO), or numerals joined by hyphens into one number (twenty-one); false for an empty text
package parse

import (
	"slices"
	"strings"
)

// Numerals joined by a hyphen are one number, as a run of numerals in
// Hangul is (이십일); numerals written apart are words of their own (two
// hundred).
func englishNumber(text string) bool {
	for _, part := range strings.Split(text, "-") {
		if !slices.Contains(englishNumerals, strings.ToLower(part)) {
			return false
		}
	}
	return true
}
