//ff:func feature=parse type=helper control=sequence
//ff:what firstRune — the first letter of a text (the replacement character when the text is empty or not UTF-8)
package parse

import "unicode/utf8"

func firstRune(text string) rune {
	r, _ := utf8.DecodeRuneInString(text)
	return r
}
