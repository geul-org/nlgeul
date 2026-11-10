//ff:func feature=token type=rule control=iteration dimension=1
//ff:what bound — is a stem words bound by a slash or a hyphen (입력/출력, 원-달러): the name of the sign's reading and its verdict, no higher than that of the weakest word (0 when the stem holds no such sign, or a word of it is not known; the last word must take what follows the stem)
package token

import (
	"math"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (sc *scan) bound(stem string, tail int) (string, float64) {
	parts := partsBound(stem)
	if len(parts) < 2 {
		return "", 0
	}
	sign := []rune(stem)[len([]rune(parts[0]))]
	left, right := []rune(parts[0]), []rune(parts[1])
	m := rule.Mark{Left: sideOf(string(left[len(left)-1:])), Right: sideOf(string(right[:1])), Run: 1}
	name, v := sc.judgeMark(sc.ix.links[sign], m, string(sign))
	for i, part := range parts {
		takes := rule.TailBare
		if i == len(parts)-1 {
			takes = tail
		}
		v = math.Min(v, sc.ix.partVerdict(part, takes))
	}
	if v <= 0 {
		return "", 0
	}
	return name, v
}
