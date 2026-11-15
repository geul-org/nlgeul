//ff:func feature=token type=helper control=iteration dimension=1
//ff:what digitsFrom — how many digits stand in a row from a place of a text
package token

import "unicode"

func digitsFrom(s []rune, from int) int {
	n := 0
	for from+n < len(s) && unicode.IsDigit(s[from+n]) {
		n++
	}
	return n
}
