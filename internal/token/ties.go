//ff:func feature=token type=reader control=sequence
//ff:what ties — may the sign between two runs tie them: a full stop always; another sign of numbers only between two runs of digits
package token

import "github.com/geul-org/koparse/pkg/rule"

func ties(s []rune, a, b [2]int) bool {
	sign := s[a[1]]
	if sign == '.' {
		return true
	}
	_, ok := signShapes[sign]
	return ok && sideOf(string(s[a[0]:a[1]]))&rule.SideDigit != 0 && sideOf(string(s[b[0]:b[1]]))&rule.SideDigit != 0
}
