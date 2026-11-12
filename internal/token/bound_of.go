//ff:func feature=token type=rule control=iteration dimension=1
//ff:what boundOf — the candidate of a stem words bound by a slash or a hyphen: the stem with each word that is written wrong as its right form (the stem itself when typos is false or none is), a candidate for each candidate of the sign that stands, with its verdict no higher than that of the weakest word (none when the stem holds no such sign, or a word of it is not known; the last word must take what follows the stem)
package token

import (
	"math"
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// With typos a word that does not stand as it is may stand as the right
// form of a word of the word table it is a wrong way to write (partRight).
func (sc *scan) boundOf(stem string, tail int, typos bool) (string, []markRead) {
	parts := partsBound(stem)
	if len(parts) < 2 {
		return stem, nil
	}
	sign := []rune(stem)[len([]rune(parts[0]))]
	left, right := []rune(parts[0]), []rune(parts[1])
	m := rule.Mark{Left: sideOf(string(left[len(left)-1:])), Right: sideOf(string(right[:1])), Run: 1}
	marks := sc.judgeMark(sc.ix.links[sign], m, string(sign))
	least := math.Inf(1)
	rights := make([]string, len(parts))
	for i, part := range parts {
		takes := rule.TailBare
		if i == len(parts)-1 {
			takes = tail
		}
		var pv float64
		rights[i], pv = sc.ix.partRight(part, takes, typos)
		least = math.Min(least, pv)
	}
	var out []markRead
	for _, mk := range marks {
		if v := math.Min(mk.v, least); v > 0 {
			out = append(out, markRead{mk.name, v})
		}
	}
	if len(out) == 0 {
		return stem, nil
	}
	return strings.Join(rights, string(sign)), out
}
