//ff:func feature=token type=rule control=sequence
//ff:what partVerdict — how well one word of a bound stem stands with what follows it: as a known word, or (a run of Latin letters and digits no table lists) as such a run; 0 for anything else
package token

import (
	"strings"

	"github.com/geul-org/koparse/pkg/rule"
)

func (ix *index) partVerdict(part string, tail int) float64 {
	if _, v := ix.known(part, tail); v > 0 {
		return v
	}
	if strings.IndexFunc(part, func(r rune) bool { return !rule.RunRune(r) }) < 0 {
		return ix.run
	}
	return 0
}
