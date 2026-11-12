//ff:func feature=token type=reader control=sequence
//ff:what trimmed — where the letters of a unit that are read begin and end (in letters), the signs at its edges left out, keeping a sign that belongs to a number there (one that stands at a number's front: -3, $40; a sign of the unit table at its end: 50%, 3℃) and the signs that belong to a name: a path's head and the slash at its end, the dashes of an option, and $ _ . before Latin letters; 0 and 0 when nothing is read
package token

import (
	"slices"
	"strings"
	"unicode"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func trimmed(field string) (int, int) {
	s := trimSigns(field)
	if len(s) == 0 {
		return 0, 0
	}
	all := []rune(field)
	from := len([]rune(field[:strings.Index(field, string(s))]))
	to := from + len(s)
	if _, ok := headSigns[atOf(all, from-1)]; ok && unicode.IsDigit(s[0]) && !rule.RunRune(atOf(all, from-2)) {
		from--
	}
	if slices.Contains(unitWords, string(atOf(all, to))) && unicode.IsDigit(s[len(s)-1]) {
		to++
	}
	from = headKept(all, from)
	if atOf(all, to) == '/' && strings.ContainsRune(string(all[from:to]), '/') {
		to++
	}
	return from, to
}
