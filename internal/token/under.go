//ff:func feature=token type=helper control=sequence
//ff:what under — the syllable a final consonant put under it comes off: the same syllable without it (간 → 가), and for ㅁ also the syllable left with ㄹ when the two stand together as ㄻ (삶 → 살); false when the consonant does not sit there
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func under(r rune, lead int) (rune, bool) {
	jong := rule.JongOf(r)
	if lead == jongM && jong == jongLM {
		return r - rune(jongLM-rule.JongL), true
	}
	return r - rune(lead), lead > 0 && jong == lead
}
