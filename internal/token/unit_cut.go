//ff:func feature=token type=reader control=sequence
//ff:what unitCut — where a run of digits and a Latin unit of the table is cut in two (the place the unit begins), or 0 and false for a run that is no such pair
package token

import "strings"

func unitCut(s []rune, r [2]int) (int, bool) {
	unit := strings.TrimLeft(string(s[r[0]:r[1]]), "0123456789")
	cut := r[1] - len([]rune(unit))
	return cut, cut > r[0] && isLatinUnit(unit)
}
