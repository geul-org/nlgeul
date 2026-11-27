//ff:func feature=parse type=helper control=sequence
//ff:what namedSign — is r a sign with a name of its own: a sign of the sign table (signNames: a comma, ㆍ), or a punctuation mark or symbol the table of shapes (special-forms.tsv) names by its shape (㈎, ⓐ, ？); a sign the other tables of special letters name (○, *, →) is none, as it was none when it had no name
package parse

func namedSign(r rune) bool {
	return len(signNames[r]) > 0 || (signRune(r) && len(forms[r]) > 0 && forms[r][0].origin == madeTable+".tsv")
}
