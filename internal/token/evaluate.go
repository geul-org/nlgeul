//ff:func feature=token type=rule control=iteration dimension=1
//ff:what evaluate — evaluate one graph in a context: the claim's verdict, and when a trace is kept the reasons of the exceptions that applied
package token

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// A graph that fails gives -1. The reasons are asked for only where a
// trace is written: the verdict is the same without them.
func evaluate(g *toulmin.Graph, ctx toulmin.Context, traced bool) (float64, []string) {
	results, err := g.Evaluate(ctx, toulmin.EvalOption{Trace: traced})
	if err != nil || len(results) == 0 {
		return -1, nil
	}
	var why []string
	for _, t := range results[0].Trace {
		if t.Role == "except" && t.RuleVerdict > -1 {
			why = append(why, whyOf(t.Specs))
		}
	}
	return results[0].Verdict, why
}
