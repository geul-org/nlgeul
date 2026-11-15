//ff:func feature=token type=rule control=sequence
//ff:what sentenceRuleOf — a rule of a sentence made for the tests: its claim is that the fact the list asks for (rule.List…) is in the list, and when unless is given an exception "why" defeats the claim when the fact unless asks for is not in its list; it speaks of the candidates about names, and waits for the predicate after those waits names (nil: it reads nothing after them)
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func sentenceRuleOf(name string, bonus float64, about, waits func(rule.Near) bool, claim, unless *rule.List, why string) sentenceRule {
	g := toulmin.NewGraph(name)
	c := g.Rule(rule.Listed).With(claim).With(&rule.Claim{Name: name, Concludes: "선호"})
	if unless != nil {
		except(g, rule.Listed, why, unless, not).Attacks(c)
	}
	return sentenceRule{name: name, graph: g, bonus: bonus, about: about, waits: waits}
}
