//ff:func feature=token type=helper control=iteration dimension=1
//ff:what particleAfterClosing — does a word have a particle after a closing among the parts after its stem (a particle written onto the word before)
package token

func particleAfterClosing(e Eojeol) bool {
	closed := false
	for _, p := range tailParts(e) {
		if closed && particlePart(p) {
			return true
		}
		closed = closed || p.Name.Kind == "맺음"
	}
	return false
}
