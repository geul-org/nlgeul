//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what amountEnd — do the letters end in a number or in a number and its unit: a number's end (numberEnd), or a word of the unit table right after one (3,000원, 50%)
package parse

func amountEnd(s []rune) bool {
	if numberEnd(s) {
		return true
	}
	for _, u := range unitWords {
		k := len(s) - len([]rune(u))
		if k >= 1 && string(s[k:]) == u && numberEnd(s[:k]) {
			return true
		}
	}
	return false
}
