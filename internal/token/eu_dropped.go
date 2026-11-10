//ff:func feature=token type=rule control=sequence
//ff:what euDropped — is tail t, read from b, a wrong way of putting back a dropped ㅡ: the stem's last syllable is written with ㅏ or ㅓ and read as its ㅡ with 어/아 fused in (따라 = 따르+아), but the stem is one the word table gives as 르 or 러 irregular (모르 goes to 몰라, never 모라), or the vowel does not follow the syllable before the ㅡ (아 after ㅏ ㅑ ㅗ, else 어, and 어 for a stem of one syllable: 써, never 싸)
package token

import "github.com/geul-org/koparse/pkg/rule"

func (sc *scan) euDropped(b int, t tail) bool {
	shown := rule.JungOf(sc.s[t.start-1])
	if t.base == 0 || rule.JungOf(t.base) != jungEu || (shown != jungA && shown != jungEo) {
		return false
	}
	stem := string(sc.s[b:t.start-1]) + string(t.base)
	if sc.ix.lex[stem].irr&(irrReu|irrReo) != 0 {
		return true
	}
	if t.start-1 == b {
		return shown == jungA
	}
	first := '어'
	if shown == jungA {
		first = '아'
	}
	return !rule.Harmony(sc.s[t.start-2], first)
}
