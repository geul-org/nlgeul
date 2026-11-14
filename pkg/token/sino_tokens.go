//ff:func feature=token type=helper control=iteration dimension=1
//ff:what sinoTokens — how many Sino-Korean numerals a text is made of: digits and greater numbers with no two digits side by side, 수 only at the front before a greater number; 0 when it is no such run
package token

import "strings"

func sinoTokens(text string) int {
	letters := []rune(text)
	digit := false
	for i, r := range letters {
		some := r == sinoSome && i == 0 && len(letters) > 1
		isDigit := strings.ContainsRune(sinoDigits, r) || some
		if (!isDigit && !strings.ContainsRune(sinoBigs, r)) || (isDigit && digit) {
			return 0
		}
		digit = isDigit
	}
	return len(letters)
}
