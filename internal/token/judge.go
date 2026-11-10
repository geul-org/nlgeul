//ff:func feature=token type=rule control=sequence
//ff:what judge — run an entry's graph for a piece at one place and note the outcome in the trace
package token

import (
	"fmt"
	"strings"

	"github.com/geul-org/koparse/pkg/rule"
)

func (sc *scan) judge(e *entry, start, j int, p rule.Piece) float64 {
	v, why := verdictOf(e.g, p, string(sc.s[:j]))
	if sc.trace == nil {
		return v
	}
	text := leadText(e.lead) + string(sc.s[start:j])
	note := "통과"
	if v <= 0 {
		note = "기각: " + strings.Join(why, ", ")
	}
	*sc.trace = append(*sc.trace, fmt.Sprintf("  · %s\t%s\t%.2f\t%s", text, e.name, v, note))
	return v
}
