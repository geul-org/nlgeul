//ff:func feature=token type=reader control=sequence
//ff:what restoresAt — the ways of taking a fused piece out of the syllable before start: those of restores, and with a word table one more for ㅕ, which is ㅣ+어 (마셔 = 마시) and also a stem's own ㅕ with 어 fused into it (켜 = 켜+어, 켰 = 켜+었); the word table tells the two apart; and one more for 마, which is also 말+아 (하지 마, 마라: clause 18 of the orthography)
package token

import "github.com/geul-org/koparse/pkg/rule"

func (sc *scan) restoresAt(start, lead int) []restore {
	r := sc.s[start-1]
	out := restores(r, lead)
	if !sc.ix.lexed {
		return out
	}
	if lead == leadVowel && r == '마' {
		out = append(out, restore{'말', "아"})
	}
	if lead == leadVowel && rule.JungOf(r) == jungYeo {
		out = append(out, restore{r, "어"})
	}
	if lead == jongSS && rule.JungOf(r) == jungYeo {
		out = append(out, restore{r - rune(jongSS), "었"})
	}
	return out
}
