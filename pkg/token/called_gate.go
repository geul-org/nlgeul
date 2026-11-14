//ff:func feature=token type=rule control=sequence
//ff:what calledGate — the graph that asks the word before the vocative 야 to be one that can be called: the claim, and the exception that it is a dependent noun of the closed set (거야 is no call to 거)
package token

import (
	"github.com/geul-org/nlgeul/pkg/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// The vocative particle is written onto the name of a person or an
// animal that is called. A dependent noun stands only with what
// modifies it and names nothing by itself.
func calledGate() *entry {
	g := toulmin.NewGraph(callGate)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: callGate, Concludes: "열쇠"})
	except(g, rule.Listed, whyBound, listed("boundNouns", rule.ListForm, parse.ClosedWords()[boundSet])).Attacks(claim)
	return &entry{name: callGate, g: g}
}
