//ff:func feature=token type=reader control=sequence
//ff:what change — does the script change at place j: a run of Latin letters and digits on one side and a Hangul syllable on the other (a place where a word may end with no marker)
package token

import "github.com/geul-org/koparse/pkg/rule"

func (sc *scan) change(j int) bool {
	if j < 1 || j >= len(sc.s) {
		return false
	}
	a, b := sideOf(string(sc.s[j-1])), sideOf(string(sc.s[j]))
	return a != 0 && b != 0 && (a&rule.SideHangul != 0) != (b&rule.SideHangul != 0)
}
