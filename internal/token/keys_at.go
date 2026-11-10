//ff:func feature=token type=reader control=sequence
//ff:what keysAt — the index keys of one token: the token itself, and as a fused token the final consonant under it (ㅁ also under ㄻ) or the vowel 어/아 fuses into
package token

import "github.com/geul-org/koparse/pkg/rule"

func keysAt(r rune) []rune {
	keys := []rune{r}
	jong := rule.JongOf(r)
	if k := rule.LeadText(jong); jong > 0 && k != "" {
		keys = append(keys, []rune(k)[0])
	}
	if jong == jongLM {
		keys = append(keys, 'ㅁ')
	}
	if jong == 0 {
		keys = append(keys, keyVowel)
	}
	return keys
}
