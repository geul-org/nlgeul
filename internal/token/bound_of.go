//ff:func feature=token type=rule control=iteration dimension=1
//ff:what boundOf — the reading of a stem words bound by a slash or a hyphen: the stem with each word that is written wrong as its right form (the stem itself when typos is false or none is), the name of the sign's reading and its verdict, no higher than that of the weakest word (0 when the stem holds no such sign, or a word of it is not known; the last word must take what follows the stem)
package token

import (
	"math"
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// With typos a word that does not stand as it is may stand as the right
// form of a word of the word table it is a wrong way to write (partRight).
func (sc *scan) boundOf(stem string, tail int, typos bool) (string, string, float64) {
	parts := partsBound(stem)
	if len(parts) < 2 {
		return stem, "", 0
	}
	sign := []rune(stem)[len([]rune(parts[0]))]
	left, right := []rune(parts[0]), []rune(parts[1])
	m := rule.Mark{Left: sideOf(string(left[len(left)-1:])), Right: sideOf(string(right[:1])), Run: 1}
	name, v := sc.judgeMark(sc.ix.links[sign], m, string(sign))
	rights := make([]string, len(parts))
	for i, part := range parts {
		takes := rule.TailBare
		if i == len(parts)-1 {
			takes = tail
		}
		var pv float64
		rights[i], pv = sc.ix.partRight(part, takes, typos)
		v = math.Min(v, pv)
	}
	if v <= 0 {
		return stem, "", 0
	}
	return strings.Join(rights, string(sign)), name, v
}
