//ff:func feature=token type=reader control=iteration dimension=1
//ff:what wholeWord — are a Hangul numeral (one syllable or more) and the piece after it together a word of the word table, or a fused form of the closed set (이건, 세대, 사회, 여러분)
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// A numeral of more syllables (여러, 다섯, 삼십) and the piece after it
// may be a word as well: 여러분 is the pronoun, not 여러 and the unit 분.
func (ix *index) wholeWord(numeral, form string) bool {
	n := []rune(numeral)
	if len(n) == 0 || slices.ContainsFunc(n, func(r rune) bool { return rule.JongOf(r) < 0 }) || ix.layers == layerNone {
		return false
	}
	if _, ok := ix.lex[numeral+form]; ok {
		return true
	}
	for _, c := range ix.contracts {
		if c.form == numeral+form {
			return true
		}
	}
	return false
}
