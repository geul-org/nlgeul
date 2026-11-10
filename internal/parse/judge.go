//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what judge — weigh every tail reading of the word; the word marked with "+" before its particles, and the constituents whose verdict is above 0, highest first, joined with "|"
package parse

import (
	"sort"
	"strings"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// judge evaluates the graph once for each tail reading (once with no
// tail when none was read). A constituent stands when some reading
// gives it a verdict above 0, the engine's point of "prevailed".
// 미정 when none stands.
func judge(g *toulmin.Graph, w *word) (string, string, error) {
	r := &ruling{verdicts: map[string]float64{}}
	for _, t := range append([]tailFact{{}}, w.tails...) {
		if err := weigh(g, w, t, r); err != nil {
			return "", "", err
		}
	}
	var labels []string
	for _, label := range printed {
		if r.verdicts[label] > 0 {
			labels = append(labels, label)
		}
	}
	if labels == nil {
		return mark(w, r), labelNone, nil
	}
	sort.SliceStable(labels, func(i, j int) bool { return r.verdicts[labels[i]] > r.verdicts[labels[j]] })
	return mark(w, r), strings.Join(labels, "|"), nil
}
