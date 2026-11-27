//ff:func feature=parse type=helper control=selection
//ff:what letterName — the name of a number written in one letter that its shape gives: a circled number (①-⑳, ㉑-㉟, ㊱-㊿) 원문자:수, a parenthesized number (⑴-⒇) 괄호문자:수, a Latin letter 특수:로마자; none (empty) for any other letter
package parse

func letterName(r rune) string {
	switch {
	case r >= 0x2474 && r <= 0x2487:
		return parenNumName
	case enclosedNumber(r):
		return circledNumName
	case r >= 'A' && r <= 'Z', r >= 'a' && r <= 'z':
		return latinLetterName
	}
	return ""
}
