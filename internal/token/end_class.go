//ff:func feature=token type=reader control=selection
//ff:what endClass — how a token ends, as the rule package names it: a vowel, ㄹ, another consonant, or not a Hangul syllable
package token

import "github.com/geul-org/koparse/pkg/rule"

func endClass(r rune) int {
	switch jongOf(r) {
	case -1:
		return rule.EndOther
	case 0:
		return rule.EndVowel
	case jongL:
		return rule.EndL
	}
	return rule.EndHard
}
