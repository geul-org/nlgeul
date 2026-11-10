//ff:func feature=token type=reader control=sequence
//ff:what asksStem — does a piece fused into the syllable before start ask, with a word table, for a predicate's stem of the table before it: an 어/아 or 았/었 fused into the stem's vowel, whether it leaves a trace in the letters (봐, 줬) or none (가, 켜), but not 해 and 했 (하 is a verb's stem whatever stands before it: a new word with 하 is read), and a lone final consonant (ㄴ, ㄹ, ㅁ) under a syllable other than the few it is always read under (하, 되, 시, 키, 이)
package token

import (
	"slices"

	"github.com/geul-org/koparse/pkg/rule"
)

// stem is the syllable as the stem had it. Without a word table the
// graph of a lone final consonant keeps to the few syllables itself.
func (sc *scan) asksStem(e *entry, start int, stem rune) bool {
	prev := sc.s[start-1]
	if e.lead == leadVowel || e.lead == jongSS {
		return prev != '해' && prev != '했'
	}
	return e.form == "" && !slices.Contains(rule.BareBases(e.lead), string(stem))
}
