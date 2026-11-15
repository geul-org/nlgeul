//ff:func feature=token type=rule control=sequence
//ff:what orderGraph — the judgment graph of the prefix 제 right before a number: the claim, the exception that no number stands right of it, and the exception that letters that do not end a number stand left of it
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// It has the shape of frontGraph for a sign at a number's very front;
// only the reasons differ, for 제 is a letter and not a sign. The reader
// (orderFront) leaves the left side empty when the letters before 제 end
// a number.
func orderGraph(digits *rule.Bits) *entry {
	g := toulmin.NewGraph(orderHeadName)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: orderHeadName, Concludes: "기호"})
	except(g, rule.MarkRight, whyOrderRight, digits).Attacks(claim)
	except(g, rule.MarkIs, whyOrderLeft, &rule.Bits{Of: "mark", Mask: rule.MarkAfterWord}).Attacks(claim)
	return &entry{name: orderHeadName, g: g}
}
