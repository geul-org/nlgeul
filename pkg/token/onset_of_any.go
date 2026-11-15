//ff:func feature=token type=helper control=iteration dimension=1
//ff:what onsetOfAny — does the syllable given begin with the initial consonant of the first syllable of one of the forms given: a stem written in one of those forms begins with that consonant whatever is fused into its last syllable (거 in 건, 걸, 겁니다 and 게), so that a place whose syllable has another is no place such a stem begins at
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func onsetOfAny(r rune, forms []string) bool {
	cho := rule.ChoOf(r)
	for _, w := range forms {
		if cho >= 0 && cho == rule.ChoOf([]rune(w)[0]) {
			return true
		}
	}
	return false
}
