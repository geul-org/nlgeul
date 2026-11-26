//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what weakPath — where a unit's letters end as a path with no head and one slash whose name after the slash holds Hangul and a sign that joins the parts of a name (nameJoins: notes/회의-결과, notes/memo-결과, docs/계약서_양식): a name of Latin letters and digits, one slash, then Hangul syllables and a name's letters with such a sign between two of them; it ends where a path ends (joinedEnd: past the Hangul after a joining sign, else after the last letter that is no Hangul syllable); 0 when the unit is not written so
package parse

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// One slash and no head are weak grounds for a path (TV/라디오 is two
// words set side by side), so this reading is a candidate beside the
// others, lowered by its graph (addPaths), and nothing is marked inside.
func weakPath(s []rune) int {
	slash := strings.IndexRune(string(s), '/')
	at := len([]rune(string(s)[:max(slash, 0)]))
	if slash < 1 || strings.Count(string(s), "/") != 1 || !isLatinLetter(s[0]) {
		return 0
	}
	end := joinedEnd(s, lastPlain(s))
	hangul, joined := false, false
	for i, r := range s {
		syllable := rule.JongOf(r) >= 0
		if i != at && !nameRune(r) && (!syllable || i < at) {
			return 0
		}
		hangul = hangul || (syllable && i < end)
		joined = joined || (i > at+1 && i+1 < len(s) && strings.ContainsRune(nameJoins, r) && s[i-1] != '/')
	}
	if !hangul || !joined {
		return 0
	}
	return end
}
