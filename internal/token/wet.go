//ff:func feature=token type=helper control=sequence
//ff:what wet — the syllable a ㅎ stem's last syllable (its ㅎ gone) makes with 어 fused into it: ㅏ → ㅐ, ㅓ → ㅔ, ㅑ → ㅒ, ㅕ → ㅖ (파래, 누레, 하얘); for the stems that go wide, ㅐ whatever the vowel (그래, 어때)
package token

import "github.com/geul-org/koparse/pkg/rule"

func wet(r rune, wide bool) rune {
	if wide {
		return r + rune((1-rule.JungOf(r))*28)
	}
	return r + 28
}
