//ff:func feature=parse type=helper control=sequence
//ff:what joinedEnd — where a path ends: past the Hangul syllables after its last letter that is no Hangul syllable when that letter is a sign that joins the parts of a name (nameJoins: a/b-가격정책), as those syllables are a part of the name; else right after that letter (a/b를)
package parse

import "strings"

func joinedEnd(s []rune, end int) int {
	if end > 0 && end < len(s) && strings.ContainsRune(nameJoins, s[end-1]) {
		return len(s)
	}
	return end
}
