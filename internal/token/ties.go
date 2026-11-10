//ff:func feature=token type=reader control=sequence
//ff:what ties — may the sign between two runs tie them: a full stop, and a sign that joins the runs of a name (a slash, a hyphen), always; another sign of numbers only between two runs of digits
package token

import "github.com/geul-org/koparse/pkg/rule"

func ties(s []rune, a, b [2]int) bool {
	sign := s[a[1]]
	if _, name := nameSigns[sign]; sign == '.' || name {
		return true
	}
	_, ok := signShapes[sign]
	return ok && sideOf(string(s[a[0]:a[1]]))&rule.SideDigit != 0 && sideOf(string(s[b[0]:b[1]]))&rule.SideDigit != 0
}
