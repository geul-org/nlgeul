//ff:func feature=parse type=reader control=sequence
//ff:what bareBase — may a final consonant with no syllable after it be read under the last syllable of s: ㅁ under 하, 되, 이; ㄴ and ㄹ under 하, 되, 시, 키, 이 or an auxiliary verb's stem
package parse

import "strings"

// Under any other syllable nothing in the tables tells the consonant
// from a plain word's.
func bareBase(s []rune, lead int) bool {
	base := s[len(s)-1] - rune(lead)
	if lead == jongM {
		return strings.ContainsRune(nominalBases, base)
	}
	return strings.ContainsRune("하되시키이", base) || helperStart(s, true) >= 1
}
