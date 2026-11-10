//ff:func feature=token type=reader control=sequence
//ff:what asksStem — does a piece fused into the syllable before start ask, with a word table, for a predicate's stem of the table before it: an 어/아 that leaves no trace in the letters (under ㅏ ㅓ ㅐ ㅔ, and under ㅕ when the stem ends in ㅕ itself: 켜, 켰), and a lone final consonant (ㄴ, ㄹ, ㅁ) under a syllable other than the few it is always read under (하, 되, 시, 키, 이)
package token

import (
	"slices"

	"github.com/geul-org/koparse/pkg/rule"
)

// stem is the syllable as the stem had it. Without a word table the
// graph of a lone final consonant keeps to the few syllables itself.
func (sc *scan) asksStem(e *entry, start int, stem rune) bool {
	prev := sc.s[start-1]
	if e.lead == leadVowel {
		return !rule.FusedVowel(prev, false) || stem == prev
	}
	if e.lead == jongSS {
		return stem == prev-rune(jongSS) && rule.JungOf(stem) == jungYeo
	}
	return e.form == "" && !slices.Contains(rule.BareBases(e.lead), string(stem))
}
