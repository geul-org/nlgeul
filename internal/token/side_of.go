//ff:func feature=token type=reader control=selection
//ff:what sideOf — what a text is, as a sign's side: a run of Latin letters and digits (of digits only, when so), a Hangul syllable, or neither
package token

import (
	"strings"

	"github.com/geul-org/koparse/pkg/rule"
)

func sideOf(text string) int {
	r := []rune(text)
	switch {
	case len(r) == 0:
		return 0
	case rule.JongOf(r[len(r)-1]) >= 0:
		return rule.SideHangul
	case !rule.RunRune(r[0]):
		return 0
	case strings.Trim(text, "0123456789") == "":
		return rule.SideAlnum | rule.SideDigit
	}
	return rule.SideAlnum
}
