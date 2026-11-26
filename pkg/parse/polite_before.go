//ff:func feature=parse type=helper control=sequence
//ff:what politeBefore — whether an old polite slot that ends in last stands before a closing that begins with the syllable r: one in 옵 before a consonant, one in 오 before ㄴ, ㄹ, ㅁ or a vowel
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

func politeBefore(last, r rune) bool {
	cho := rule.ChoOf(r)
	if last == politeHard {
		return cho >= 0 && cho != choNone
	}
	return cho == choN || cho == choL || cho == choM || cho == choNone
}
