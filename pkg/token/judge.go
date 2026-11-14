//ff:func feature=token type=rule control=sequence
//ff:what judge — run an entry's graph for a piece at one place and note the outcome in the trace
package token

import (
	"fmt"
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (sc *scan) judge(e *entry, start, j int, p rule.Piece) float64 {
	if sc.trace == nil {
		return sc.ix.judged(e.g, p, sc.letters(0, j))
	}
	v, why := verdictOf(e.g, p, sc.letters(0, j), true)
	text := sc.letters(start, j)
	if e.lead != 0 {
		text = "(" + restores(sc.s[start-1], e.lead)[0].piece + ")" + text
	}
	note := "통과"
	if v <= 0 {
		note = "기각: " + strings.Join(why, ", ")
	}
	*sc.trace = append(*sc.trace, fmt.Sprintf("  · %s\t%s\t%.2f\t%s", text, e.name, v, note))
	return v
}
