//ff:func feature=token type=rule control=sequence
//ff:what nameGraph — the judgment graph of a sign that joins two runs into a name (a path's slash, a hyphen): the claim, and the exceptions that a side is no run of Latin letters and digits and that both sides are numbers
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func nameGraph(name string) *entry {
	g := toulmin.NewGraph(name)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: name, Concludes: "기호"})
	g.Except(rule.MarkSides).With(&rule.Bits{Of: "sides", Mask: rule.SideAlnum}).Attacks(claim)
	g.Except(rule.MarkBoth).With(&rule.Bits{Of: "sides", Mask: rule.SideDigit}).Attacks(claim)
	return &entry{name: name, g: g}
}
