//ff:func feature=token type=reader control=iteration dimension=1
//ff:what stretch — one reading's way through the stretch (from, to]: its words there, printed, with the least verdict and the summed preference
package token

import (
	"math"
	"strings"
)

func stretch(p *path, from, to int) branch {
	b := branch{verdict: 1}
	var marked, names []string
	for _, e := range wordsOf(p) {
		if e.end > from && e.end <= to {
			marked, names = append(marked, e.marked), append(names, e.names)
			b.verdict, b.pref = math.Min(b.verdict, e.verdict), b.pref+e.pref
		}
	}
	b.marked, b.names = strings.Join(marked, " "), strings.Join(names, " + ")
	return b
}
