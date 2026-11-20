//ff:func feature=parse type=helper control=selection
//ff:what isLatinLetter — is a letter an ASCII Latin letter or a fullwidth Latin letter (U+FF21~FF3A, U+FF41~FF5A)
package parse

func isLatinLetter(r rune) bool {
	switch {
	case r >= 'A' && r <= 'Z', r >= 'a' && r <= 'z':
		return true
	case r >= 0xFF21 && r <= 0xFF3A, r >= 0xFF41 && r <= 0xFF5A:
		return true
	}
	return false
}
