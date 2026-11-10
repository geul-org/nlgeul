//ff:func feature=token type=reader control=iteration dimension=1
//ff:what partsPref — the preference of a stem of bound words: the sum of the words' own (a word the table gives no count costs what its tokens would: a run of Latin letters and digits is one token, as it is in a new word)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func (ix *index) partsPref(parts []string) float64 {
	sum := 0.0
	for _, part := range parts {
		sum += ix.stemPref(part, rule.TailBare, true, partTokens(part))
	}
	return sum
}
