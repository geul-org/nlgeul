//ff:func feature=token type=rule control=sequence
//ff:what addCopula — the graphs of the copula stem: 이 said after a noun, and 이 dropped after a vowel; the claim "this is the stem of 이다 after a noun", and the exceptions that no stem stands before it, that the closing after it is one the copula does not take, and for the dropped stem that no vowel stands before it
package token

import (
	"github.com/geul-org/koparse/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// The copula's count is not added to the total the other markers'
// shares are taken of, so that their preferences stay as they were.
func (ix *index) addCopula() {
	barred := toulmin.NewStringListSpec("barred", ix.barred)
	g := toulmin.NewGraph(copulaName)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: copulaName, Concludes: "표지"})
	g.Except(rule.PieceNoStem).Attacks(claim)
	g.Except(rule.NextListed).With(barred).Attacks(claim)
	ix.copula = &entry{name: copulaName, form: "이", count: copulaCount, g: g}

	d := toulmin.NewGraph(droppedName)
	claim = d.Rule(rule.PieceRead).With(&rule.Claim{Name: droppedName, Concludes: "표지"})
	d.Except(rule.PieceNoStem).Attacks(claim)
	d.Except(rule.NextListed).With(barred).Attacks(claim)
	d.Except(rule.PieceAfter).With(&rule.Bits{Of: "after", Mask: afterVowel}).Attacks(claim)
	ix.dropped = &entry{name: droppedName, count: copulaCount, g: d}
	ix.graphs, ix.edges = ix.graphs+2, ix.edges+5
}
