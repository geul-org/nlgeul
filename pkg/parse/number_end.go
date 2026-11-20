//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what numberEnd — do the letters end in a number: a digit, or greater numbers written right after a digit (17억, 3천만), which are pieces of the number before them
package parse

import (
	"strings"
	"unicode"
)

func numberEnd(s []rune) bool {
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(s[i]) {
			return true
		}
		if !strings.ContainsRune(sinoBigs, s[i]) {
			return false
		}
	}
	return false
}
