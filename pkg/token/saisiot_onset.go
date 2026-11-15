//ff:func feature=token type=helper control=selection
//ff:what saisiotOnset — whether a syllable begins a word that 사이시옷 may be put before (「한글 맞춤법」 제30항): its onset is ㄱ ㄷ ㅂ ㅅ ㅈ (sounded hard after it), ㄴ ㅁ (ㄴ added before them), or ㅇ with ㅣ ㅑ ㅕ ㅛ ㅠ (ㄴㄴ added: 「표준 발음법」 제29항), but the syllable is not 이 itself (제27항 [붙임 3]: 이 sounded so after ㅅ is written 니); false for any other syllable and for what is not a Hangul syllable
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// Onsets by their numbers (rule.ChoOf) and vowels by theirs
// (rule.JungOf). ㅊ ㅋ ㅌ ㅍ, the doubled consonants, ㅎ and ㄹ are not
// among them: 위턱 and 위쪽 have no 사이시옷, as their back word is
// neither sounded hard nor has ㄴ added before it (제30항).
func saisiotOnset(r rune) bool {
	switch rule.ChoOf(r) {
	case 0, 3, 7, 9, 12: // ㄱ ㄷ ㅂ ㅅ ㅈ
		return true
	case 2, 6: // ㄴ ㅁ
		return true
	case 11: // ㅇ
		j := rule.JungOf(r)
		return r != '이' && (j == 20 || j == 2 || j == 6 || j == 12 || j == 17) // ㅣ ㅑ ㅕ ㅛ ㅠ
	}
	return false
}
