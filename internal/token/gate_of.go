//ff:func feature=token type=rule control=sequence
//ff:what gateOf — the graph that asks the word before a piece to be a stem of the word table of some kinds: the claim, and the exception that it has none of those kinds
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func gateOf(name string, kinds int) *entry {
	g := toulmin.NewGraph(name)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: name, Concludes: "열쇠"})
	g.Except(rule.StemKind).With(&rule.Bits{Of: "kinds", Mask: kinds}).Attacks(claim)
	return &entry{name: name, g: g}
}
