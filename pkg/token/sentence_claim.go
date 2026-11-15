//ff:func feature=token type=rule control=sequence
//ff:what sentenceClaim — the graph of a rule of a sentence with its claim on it: the graph is named as the rule is, and the claim "this candidate is preferred here" stands for the candidates the list names (the rule function given, rule.Listed or rule.SelfHas, with the list; a trace prints the rule's name for it); the exceptions are put on the claim by the rule itself
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func sentenceClaim(name string, fn any, of *rule.List) (*toulmin.Graph, *toulmin.Rule) {
	g := toulmin.NewGraph(name)
	return g, g.Rule(fn).With(of).With(&rule.Claim{Name: name, Concludes: preferLabel})
}
