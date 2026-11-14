//ff:func feature=token type=helper control=iteration dimension=1
//ff:what numberWord — is a word a number: its stem is a number (numberStem) and the parts after it are none, or greater numbers and approximations only
package token

func numberWord(e Eojeol) bool {
	if !numberStem(stemPart(e)) {
		return false
	}
	for _, p := range tailParts(e) {
		if p.Name.Kind != "큰 수" && p.Name.Kind != "어림" {
			return false
		}
	}
	return true
}
