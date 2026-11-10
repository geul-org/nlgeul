//ff:func feature=token type=reader control=iteration dimension=1
//ff:what right — what stands right after place j: the unit's end (or a letter that is no Hangul syllable), a particle, or something else
package token

import (
	"strings"

	"github.com/geul-org/koparse/pkg/rule"
)

func (sc *scan) right(j int) int {
	if j >= len(sc.s) || sc.s[j] < '가' || sc.s[j] > '힣' {
		return rule.RightDone
	}
	rest := string(sc.s[j:])
	for _, f := range sc.ix.heads[sc.s[j]] {
		if strings.HasPrefix(rest, f) {
			return rule.RightParticle
		}
	}
	return rule.RightOther
}
