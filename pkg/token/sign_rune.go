//ff:func feature=token type=helper control=sequence
//ff:what signRune — whether a letter is a punctuation mark or a symbol
package token

import "unicode"

func signRune(r rune) bool {
	return unicode.IsPunct(r) || unicode.IsSymbol(r)
}
