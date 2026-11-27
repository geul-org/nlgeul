//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what enclosedNumber — is r a circled or parenthesized number (①, ⓪, ⑴, ㉑, ㊿): a letter the tables of special letters give one of the shapes of enclosedShapes and the content 수
package parse

import "slices"

func enclosedNumber(r rune) bool {
	for _, f := range forms[r] {
		if slices.Contains(enclosedShapes, f.shape) && f.content == numName {
			return true
		}
	}
	return false
}
