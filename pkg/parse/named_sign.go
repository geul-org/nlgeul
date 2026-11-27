//ff:func feature=parse type=helper control=sequence
//ff:what namedSign — is r a sign with a name of its own: a sign of the sign table (signNames: a comma, ㆍ), or a punctuation mark or symbol the tables of special letters name by its shape (㈎, ⓐ, ？)
package parse

func namedSign(r rune) bool {
	return len(signNames[r]) > 0 || (signRune(r) && len(forms[r]) > 0)
}
