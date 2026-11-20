//ff:func feature=parse type=rule control=sequence
//ff:what nominalNewGateOf — the graph that asks the word before a piece to be a nominal of the word table or a word no table lists (a new word): the claim, the exception that its stem has none of the nominal kinds, and the exception to that exception that the tables know no kind of the stem
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func nominalNewGateOf(name string) *entry {
	g := toulmin.NewGraph(name)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: name, Concludes: "열쇠"})
	other := except(g, rule.StemKind, whyKind, &rule.Bits{Of: "kinds", Mask: kindNominal}, not)
	other.Attacks(claim)
	except(g, rule.StemKind, whyUnlisted, &rule.Bits{Of: "kinds", Mask: kindAny}, not).Attacks(other)
	return &entry{name: name, g: g}
}
