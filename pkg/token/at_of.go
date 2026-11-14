//ff:func feature=token type=helper control=sequence
//ff:what atOf — the letter at a place of a text, or 0 outside it
package token

func atOf(s []rune, i int) rune {
	if i < 0 || i >= len(s) {
		return 0
	}
	return s[i]
}
