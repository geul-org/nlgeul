//ff:func feature=parse type=reader control=sequence
//ff:what changed — the forms a predicate's stem changes into before some tails, each with the kind of change: ㄹ dropped (whatever the irregular types), and by the irregular types the table gives it ㄷ → ㄹ, ㅅ dropped, ㅂ → 우 (돕, 곱 also 오; a stem of three syllables or more that ends in 스럽 also without the 우: 자랑스럽 → 자랑스러, read before the modifying ㄴ), ㅎ dropped (and with 어 fused), 르 → ㄹ르, 러, 우 → ㅡ; and a stem of two syllables or more that ends in 하 without its 하 (clause 40: 원하 → 원, read before 치, 케, 타, 토록, 코자, 컨대), and once more for note 2 when the syllable before 하 ends in a final consonant that sounds [ㄱ], [ㄷ] or [ㅂ] (생각하 → 생각, read before 건대)
package parse

import (
	"slices"
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The stem field of what comes back is the changed form.
func changed(key string, irr int) []alt {
	r := []rune(key)
	last := r[len(r)-1]
	head, jong := string(r[:len(r)-1]), rule.JongOf(last)
	var out []alt
	if jong == rule.JongL {
		out = append(out, alt{head + string(last-rune(jong)), irrL})
	}
	if irr&irrD != 0 && jong == jongD {
		out = append(out, alt{head + string(last+1), irrD})
	}
	if irr&irrS != 0 && jong == jongS {
		out = append(out, alt{head + string(last-rune(jong)), irrS})
	}
	if irr&irrB != 0 && jong == rule.JongB {
		out = append(out, alt{head + string(last-rune(jong)) + "우", irrB})
	}
	if irr&irrB != 0 && len(r) >= 3 && strings.HasSuffix(key, seuEnd) {
		out = append(out, alt{head + string(last-rune(jong)), irrSeu})
	}
	if irr&irrB != 0 && slices.Contains(roundB, key) {
		out = append(out, alt{head + string(last-rune(jong)) + "오", irrBO})
	}
	if irr&irrH != 0 && jong == jongH {
		out = append(out, alt{head + string(last-rune(jong)), irrH}, alt{head + string(wet(last-rune(jong), slices.Contains(wideH, key))), irrHE})
	}
	if irr&irrReu != 0 && len(r) >= 2 && last == '르' && rule.JongOf(r[len(r)-2]) == 0 {
		out = append(out, alt{string(r[:len(r)-2]) + string(r[len(r)-2]+rune(rule.JongL)) + "르", irrReu})
	}
	if irr&irrReo != 0 && last == '르' {
		out = append(out, alt{key + "르", irrReo})
	}
	if irr&irrU != 0 && jong == 0 && rule.JungOf(last) == jungU {
		out = append(out, alt{head + string(last+rune((jungEu-jungU)*28)), irrU})
	}
	if len(r) >= 2 && last == '하' {
		out = append(out, alt{head, irrHa})
	}
	if len(r) >= 2 && last == '하' && slices.Contains(voicelessEnds, rule.JongOf(r[len(r)-2])) {
		out = append(out, alt{head, irrHaW})
	}
	return out
}
