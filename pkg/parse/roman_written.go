//ff:func feature=parse type=helper control=sequence
//ff:what romanWritten — is a text a Roman numeral written in the Latin capitals I, V and X in the numeral's own order (II, IV, XIV, XXIII, and I, V, X alone); false for an empty text
package parse

// Only the numerals of I, V and X are read: the numerals of L, C, D and
// M are written as abbreviations are (CD, MD, DC, MC, CM, MM, CI, MV,
// XL), and nothing in the letters tells the two apart.
func romanWritten(text string) bool {
	return text != "" && romanShape.MatchString(text)
}
