//ff:func feature=token type=reader control=sequence
//ff:what trimUnit — a unit without the signs at its edges, keeping a sign that belongs to a number there (one that stands at a number's front: -3, $40; a unit sign at its end: 50%, 3℃) and the signs that belong to a name: a path's head and the slash at its end, the dashes of an option, and @ # $ _ . before Latin letters
package token

import (
	"strings"
	"unicode"

	"github.com/geul-org/koparse/pkg/rule"
)

func trimUnit(field string) []rune {
	s := rule.Trim(field)
	if len(s) == 0 {
		return s
	}
	all := []rune(field)
	from := len([]rune(field[:strings.Index(field, string(s))]))
	to := from + len(s)
	if _, ok := headSigns[atOf(all, from-1)]; ok && unicode.IsDigit(s[0]) && !rule.RunRune(atOf(all, from-2)) {
		from--
	}
	if r := atOf(all, to); (r == '%' || r == '℃') && unicode.IsDigit(s[len(s)-1]) {
		to++
	}
	from = headKept(all, from)
	if atOf(all, to) == '/' && strings.ContainsRune(string(all[from:to]), '/') {
		to++
	}
	return all[from:to]
}
