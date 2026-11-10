//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what weigh — evaluate the graph for the word with one tail reading and keep, for each constituent, the highest verdict so far; note how the word divides under the claims that stand
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A result is tied to its claim and conclusion by the node's Claim
// criterion, which the engine writes into the result's name.
func weigh(g *toulmin.Graph, w *word, t tailFact, r *ruling) error {
	ctx := toulmin.NewContext()
	ctx.Set(rule.KeyWord, facts(w))
	ctx.Set(rule.KeyTail, tailFacts(t))
	results, err := g.Evaluate(ctx)
	if err != nil {
		return err
	}
	for _, res := range results {
		name, label := rule.ClaimOf(res.Name)
		if old, ok := r.verdicts[label]; !ok || res.Verdict > old {
			r.verdicts[label] = res.Verdict
		}
		if res.Verdict > 0 {
			noteSplit(r, name, t)
		}
	}
	return nil
}
