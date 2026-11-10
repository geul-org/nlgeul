//ff:func feature=token type=helper control=iteration dimension=1
//ff:what lastPlain — the place right after the last letter of a unit that is no Hangul syllable (0 when every letter is one)
package token

import "github.com/geul-org/koparse/pkg/rule"

func lastPlain(s []rune) int {
	for i := len(s); i > 0; i-- {
		if rule.JongOf(s[i-1]) < 0 {
			return i
		}
	}
	return 0
}
