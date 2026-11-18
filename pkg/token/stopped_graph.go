//ff:func feature=token type=rule control=sequence
//ff:what stoppedGraph — the judgment graph of one shape of a date of two numbers joined by a full stop (stopShapes): the claim "this sign is that", the exception that the runs it joins are not of the shape, and the exception that no full stop stands right after the second number (MarkStopped)
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func stoppedGraph(s rule.Shape) *entry {
	g := toulmin.NewGraph(s.Of)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: s.Of, Concludes: "기호"})
	except(g, rule.PartsShape, whyShape, &s).Attacks(claim)
	except(g, rule.MarkIs, whyNotStopped, &rule.Bits{Of: "mark", Mask: rule.MarkStopped}, not).Attacks(claim)
	return &entry{name: s.Of, g: g}
}
