//ff:func feature=parse type=reader control=sequence
//ff:what aForm — does a stem end in its 어/아 form: a syllable 어/아/여, or a vowel 어/아 is fused into (해, 봐, 가)
package parse

import (
	"strings"

	"github.com/geul-org/koparse/pkg/rule"
)

func aForm(stem []rune) bool {
	if len(stem) == 0 {
		return false
	}
	last := stem[len(stem)-1]
	return strings.ContainsRune("어아여", last) || rule.FusedVowel(last, true)
}
