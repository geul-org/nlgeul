//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what claim — declare one claim node: the rule function, its criteria, and the Claim that names the node and its conclusion
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func claim(g *toulmin.Graph, c map[string]*toulmin.Rule, name, concludes string, fn any, criteria ...toulmin.Spec) {
	r := g.Rule(fn)
	for _, s := range criteria {
		r.With(s)
	}
	c[name] = r.With(&rule.Claim{Name: name, Concludes: concludes})
}
