//ff:func feature=token type=helper control=selection
//ff:what nominalStem — is a stem a nominal: the word table's 체언, a dependent noun, a pronoun, a number, or a word no table names (a new word); a part that is no stem is none
package token

func nominalStem(p Part) bool {
	switch {
	case !p.Stem:
		return false
	case p.Name.Kind == "" && p.Name.Form == "":
		return true
	case p.Name.Kind == "열쇠":
		return p.Name.Form == "체언"
	}
	return p.Name.Kind == "의존명사" || p.Name.Kind == "대명사" || p.Name.Kind == "수"
}
