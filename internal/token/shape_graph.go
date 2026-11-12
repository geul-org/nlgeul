//ff:func feature=token type=rule control=sequence
//ff:what shapeGraph — the judgment graph of one candidate of a sign between runs of digits: the claim "this sign is that", and the exception that the runs it joins are not of the candidate's shape
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func shapeGraph(s rule.Shape) *entry {
	g := toulmin.NewGraph(s.Of)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: s.Of, Concludes: "기호"})
	except(g, rule.PartsShape, whyShape, &s).Attacks(claim)
	return &entry{name: s.Of, g: g}
}
