//ff:func feature=token type=helper control=iteration dimension=1
//ff:what hangulWords — the words whose first letter is a Hangul syllable, in their order
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func hangulWords(words []string) []string {
	var out []string
	for _, w := range words {
		if r := []rune(w); len(r) > 0 && rule.JongOf(r[0]) >= 0 {
			out = append(out, w)
		}
	}
	return out
}
