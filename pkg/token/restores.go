//ff:func feature=token type=reader control=iteration dimension=1
//ff:what restores — the ways of taking a piece with a lead out of the syllable it sits in: a final consonant comes off (마신 → 마시 + ㄴ, and ㅁ off ㄻ: 삶 → 살 + ㅁ); a fused ㅆ comes off as 았/었/였 and a fused 어/아 as 아/어/여, with the stem's vowel put back
package token

import "github.com/geul-org/nlgeul/pkg/rule"

var pastOf = map[string]string{"아": "았", "어": "었", "여": "였"}

func restores(r rune, lead int) []restore {
	if base, ok := under(r, lead); lead != rule.JongSS && ok {
		return []restore{{base, rule.LeadText(lead)}}
	}
	if lead == rule.LeadVowel {
		return vowelRestores(r)
	}
	var out []restore
	for _, o := range vowelRestores(r - rule.JongSS) {
		out = append(out, restore{o.stem, pastOf[o.piece]})
	}
	return out
}
