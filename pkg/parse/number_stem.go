//ff:func feature=parse type=helper control=sequence
//ff:what numberStem — is a stem a number: a stem named 수, or whose letters are Sino-Korean or native numerals (before a unit)
package parse

func numberStem(p Part) bool {
	return p.Stem && (p.Name.Kind == "수" || sinoTokens(p.Text) > 0 || nativeTokens(p.Text) > 0)
}
