//ff:func feature=token type=rule control=iteration dimension=1
//ff:what judgeMark — run the graphs of a sign's readings for one place: the names of the readings that stand, joined with "|", and their highest verdict; each outcome is noted in the trace
package token

import (
	"fmt"
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func (sc *scan) judgeMark(graphs []*entry, m rule.Mark, text string) (string, float64) {
	var names []string
	best := -1.0
	for _, e := range graphs {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, m)
		v, why := evaluate(e.g, ctx)
		note := "통과"
		if v <= 0 {
			note = "기각: " + strings.Join(why, ", ")
		}
		if sc.trace != nil {
			*sc.trace = append(*sc.trace, fmt.Sprintf("  · %s\t%s\t%.2f\t%s", text, e.name, v, note))
		}
		if v > 0 {
			names, best = append(names, e.name), max(best, v)
		}
	}
	return strings.Join(names, "|"), best
}
