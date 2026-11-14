//ff:func feature=token type=helper control=iteration dimension=2
//ff:what hasUnitWord — whether some candidate of some contest of a sentence prints exactly as the given text
package token

func hasUnitWord(r *Reader, s, text string) bool {
	for _, c := range r.Contests(s) {
		if hasBranch(c.Branches, text) {
			return true
		}
	}
	return false
}
