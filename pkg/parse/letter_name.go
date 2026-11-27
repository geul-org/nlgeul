//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what letterName — the name of a number written in one letter that its shape gives: a circled or parenthesized number (enclosedNumber) is named as the tables of special letters name it, by its shape and what it writes (① 원:수, ⑴ 괄호:수), a Latin letter 특수:로마자; none (empty) for any other letter
package parse

func letterName(r rune) string {
	for _, f := range forms[r] {
		if enclosedNumber(r) && f.content == numName {
			return f.name()
		}
	}
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return latinLetterName
	}
	return ""
}
