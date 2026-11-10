//ff:func feature=token type=rule control=sequence
//ff:what frontGraph — the judgment graph of a sign that stands before a number: the claim, the exception that no number stands right of it, and (for a sign at a number's very front) the exception that letters stand left of it
package token

import (
	"github.com/geul-org/koparse/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func frontGraph(name string, digits *rule.Bits, front bool) *entry {
	g := toulmin.NewGraph(name)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: name, Concludes: "기호"})
	g.Except(rule.MarkRight).With(digits).Attacks(claim)
	if front {
		g.Except(rule.MarkAfterWord).Attacks(claim)
	}
	return &entry{name: name, g: g}
}
