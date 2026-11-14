//ff:func feature=parse type=reader control=sequence
//ff:what bareBase — may a final consonant with no syllable after it be read under the last syllable of s: under the syllables the rule package lists, and for ㄴ and ㄹ also under an auxiliary verb's stem
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// Under any other syllable nothing in the tables tells the consonant
// from a plain word's.
func bareBase(s []rune, lead int) bool {
	base := string(s[len(s)-1] - rune(lead))
	if slices.Contains(rule.BareBases(lead), base) {
		return true
	}
	return lead != jongM && helperStart(s, true) >= 1
}
