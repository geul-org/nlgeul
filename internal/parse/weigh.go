//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what weigh — evaluate the graph for the word with one tail reading and keep, for each constituent, the highest verdict so far; note how the word divides under the claims that stand
package parse

import (
	"strings"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func weigh(g *toulmin.Graph, w *word, t tailFact, r *ruling) error {
	ctx := toulmin.NewContext()
	ctx.Set(keyWord, w)
	ctx.Set(keyTail, t)
	results, err := g.Evaluate(ctx)
	if err != nil {
		return err
	}
	for _, res := range results {
		name, _, _ := strings.Cut(res.Name, "#")
		label := conclusions[name]
		if old, ok := r.verdicts[label]; !ok || res.Verdict > old {
			r.verdicts[label] = res.Verdict
		}
		if res.Verdict > 0 {
			noteSplit(r, name, t)
		}
	}
	return nil
}
