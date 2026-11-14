//ff:func feature=token type=reader control=sequence
//ff:what fitsUnit — article 43: the word before the boundary is a number that does not begin with a digit, and the stem of the word after it is a unit of the unit table
package token

import "slices"

func fitsUnit(g gap) bool {
	l := stemPart(g.left)
	return numberWord(g.left) && !isDigitFirst(l.Text) && slices.Contains(unitWords, stemPart(g.right).Text)
}
