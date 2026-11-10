//ff:func feature=token type=rule control=iteration dimension=1
//ff:what evaluate — evaluate one graph in a context: the claim's verdict, and the exceptions that applied
package token

import (
	"strings"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A graph that fails gives -1.
func evaluate(g *toulmin.Graph, ctx toulmin.Context) (float64, []string) {
	results, err := g.Evaluate(ctx, toulmin.EvalOption{Trace: true})
	if err != nil || len(results) == 0 {
		return -1, nil
	}
	var why []string
	for _, t := range results[0].Trace {
		name, _, _ := strings.Cut(t.Name, "#")
		if t.Role == "except" && t.RuleVerdict > -1 {
			why = append(why, reasons[name])
		}
	}
	return results[0].Verdict, why
}
