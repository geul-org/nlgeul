//ff:func feature=token type=reader control=sequence
//ff:what fitsMyriad — article 44: the word before the boundary is a number that ends in 만, 억 or 조, and the stem of the word after it is a number
package token

import "strings"

func fitsMyriad(g gap) bool {
	return numberWord(g.left) && strings.ContainsRune("만억조", lastLetter(g.left)) && numberStem(stemPart(g.right))
}
