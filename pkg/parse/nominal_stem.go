//ff:func feature=parse type=helper control=selection
//ff:what nominalStem — is a stem a nominal: the word table's 체언, a dependent noun, a pronoun, a number (a decimal number too: 3.5), or a word no table names (a new word); a part that is no stem is none
package parse

func nominalStem(p Part) bool {
	switch {
	case !p.Stem:
		return false
	case p.Name.Kind == "" && p.Name.Form == "":
		return true
	case p.Name.Kind == "열쇠":
		return p.Name.Form == "체언"
	}
	// A decimal number (3.5) is named by its dot, 소수점; it is a number
	// all the same.
	return p.Name.Kind == "의존명사" || p.Name.Kind == "대명사" || numberName(p.Name) || p.Name.Kind == "소수점"
}
