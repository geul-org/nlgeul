//ff:func feature=parse type=reader control=sequence
//ff:what fitsUnitDigit — the proviso of article 43: as fitsUnit, but the number before the boundary begins with a digit
package parse

import "slices"

func fitsUnitDigit(g gap) bool {
	l := stemPart(g.left)
	return numberWord(g.left) && isDigitFirst(l.Text) && slices.Contains(unitWords, stemPart(g.right).Text)
}
