//ff:func feature=token type=rule control=sequence
//ff:what sentenceRebuttal — the graph of a rule of rebuttal of a sentence with its claim on it: the graph is named as the rule is, and the claim "the grammar does not let this candidate stand here" stands for the candidates the list names (the rule function given, rule.Listed or rule.SelfHas, with the list; a trace prints the rule's name for it); the exceptions are put on the claim by the rule itself, and the rule carries the force forceUngrammatical
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func sentenceRebuttal(name string, fn any, of *rule.List) (*toulmin.Graph, *toulmin.Rule) {
	g := toulmin.NewGraph(name)
	return g, g.Rule(fn).With(of).With(&rule.Claim{Name: name, Concludes: rebutLabel})
}
