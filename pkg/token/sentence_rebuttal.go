//ff:func feature=token type=rule control=iteration dimension=1
//ff:what sentenceRebuttal — the graph of a rule of rebuttal of a sentence with its claim on it: the graph is named as the rule is, and the claim "the grammar does not let this candidate stand here" stands for the candidates the list names (the rule function given, rule.Listed, rule.SelfHas or rule.AroundIs, with its criteria; a trace prints the rule's name for it); the rows of the case exception table under the rule's name are put on it as one exception (sentenceCases), and the other exceptions by the rule itself, and the rule carries the force forceUngrammatical
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func sentenceRebuttal(name string, fn any, specs ...toulmin.Spec) (*toulmin.Graph, *toulmin.Rule) {
	g := toulmin.NewGraph(name)
	r := g.Rule(fn)
	for _, s := range specs {
		r.With(s)
	}
	claim := r.With(&rule.Claim{Name: name, Concludes: rebutLabel})
	sentenceCases(g, claim, name)
	return g, claim
}
