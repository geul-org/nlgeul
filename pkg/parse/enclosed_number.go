//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what enclosedNumber — is r a circled or parenthesized number (①, ⓪, ⑴, ㉑, ㊿): a letter the tables of special letters give one of the shapes of legacyShapes and the content 수
package parse

func enclosedNumber(r rune) bool {
	for _, f := range forms[r] {
		if _, ok := legacyShapes[f.shape]; ok && f.content == numName {
			return true
		}
	}
	return false
}
