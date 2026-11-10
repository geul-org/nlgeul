//ff:func feature=token type=reader control=iteration dimension=1
//ff:what wholeWord — are a lone Hangul numeral and the piece after it together a word of the word table, or a fused form of the closed set (이건, 세대, 사회)
package token

import "github.com/geul-org/koparse/pkg/rule"

func (ix *index) wholeWord(numeral, form string) bool {
	n := []rune(numeral)
	if len(n) != 1 || rule.JongOf(n[0]) < 0 || ix.layers == layerNone {
		return false
	}
	if _, ok := ix.lex[numeral+form]; ok {
		return true
	}
	for _, c := range contracts {
		if c.form == numeral+form {
			return true
		}
	}
	return false
}
