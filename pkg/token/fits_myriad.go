//ff:func feature=token type=reader control=sequence
//ff:what fitsMyriad — article 44: the word before the boundary is a number that ends in 만, 억 or 조 (sinoMyriads), and the stem of the word after it is a number
package token

import "strings"

func fitsMyriad(g gap) bool {
	return numberWord(g.left) && strings.ContainsRune(sinoMyriads, lastLetter(g.left)) && numberStem(stemPart(g.right))
}
