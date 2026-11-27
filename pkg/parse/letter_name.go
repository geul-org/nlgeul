//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what letterName — the name of a number written in one letter that its shape gives: a circled number (①, ㉑, ㊿) 원문자:수, a parenthesized number (⑴-⒇) 괄호문자:수 (the tables of special letters tell the shape: legacyShapes), a Latin letter 특수:로마자; none (empty) for any other letter
package parse

func letterName(r rune) string {
	for _, f := range forms[r] {
		if name, ok := legacyShapes[f.shape]; ok && f.content == numName {
			return name
		}
	}
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return latinLetterName
	}
	return ""
}
