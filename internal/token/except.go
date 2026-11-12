//ff:func feature=token type=rule control=iteration dimension=1
//ff:what except — put an exception on a graph: the rule function, its criteria, and why it defeats what it attacks, kept on the node itself (a Claim) so that a trace prints the reason of that node and not of the function
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// The first criterion is the one the rule function reads.
func except(g *toulmin.Graph, fn any, why string, criteria ...toulmin.Spec) *toulmin.Rule {
	r := g.Except(fn)
	for _, s := range criteria {
		r.With(s)
	}
	return r.With(&rule.Claim{Name: why, Concludes: exceptLabel})
}
