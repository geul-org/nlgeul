//ff:func feature=token type=helper control=iteration dimension=1
//ff:what auxHead — do the letters rest begin with stem, its last syllable as it stands, with a final consonant put under it (볼, 준), or with 어/아 fused into its vowel (줘, 봐, 둬, 버려, and with ㅆ under that: 줬, 봤)
package token

import (
	"slices"

	"github.com/geul-org/koparse/pkg/rule"
)

func auxHead(rest, stem []rune) bool {
	n := len(stem)
	if len(rest) < n || !slices.Equal(rest[:n-1], stem[:n-1]) {
		return false
	}
	r, last := rest[n-1], stem[n-1]
	if rule.JungOf(r) < 0 || rule.JongOf(last) != 0 {
		return r == last
	}
	open := r - rune(rule.JongOf(r))
	if open == last {
		return true
	}
	for _, o := range vowelRestores(open) {
		if o.stem == last && (rule.JongOf(r) == 0 || rule.JongOf(r) == jongSS) {
			return true
		}
	}
	return false
}
