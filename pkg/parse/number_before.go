//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what numberBefore — do the letters right before place j end a number: a digit, or a word of the unit table right after a digit (36조│제1항, 조의2│제1항)
package parse

import "unicode"

func (sc *scan) numberBefore(j int) bool {
	if j >= 1 && unicode.IsDigit(sc.s[j-1]) {
		return true
	}
	for _, u := range unitWords {
		n := len([]rune(u))
		if j-n >= 1 && string(sc.s[j-n:j]) == u && unicode.IsDigit(sc.s[j-n-1]) {
			return true
		}
	}
	return false
}
