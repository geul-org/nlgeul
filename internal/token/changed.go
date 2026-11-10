//ff:func feature=token type=reader control=sequence
//ff:what changed — the forms a predicate's stem changes into before some tails, each with the kind of change: ㄹ dropped (whatever the irregular types), and by the irregular types the table gives it ㄷ → ㄹ, ㅅ dropped, ㅂ → 우 (돕, 곱 also 오), ㅎ dropped (and with 어 fused), 르 → ㄹ르, 러, 우 → ㅡ
package token

import (
	"slices"

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
	return out
}
