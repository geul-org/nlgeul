//ff:func feature=token type=rule control=selection
//ff:what dotGraph — the judgment graph of one reading of the full stop: the claim "this dot is that", and the structural exceptions against it
package token

import (
	"github.com/geul-org/koparse/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// 마침표 ends a sentence: not between two runs, not one of dots in a
// row. 줄임표 is dots in a row. 파일명 joins a name and a listed
// extension, 도메인 the parts of a name whose last is a listed domain,
// 소수점 two numbers, 버전 numbers with a leading v or more dots than
// one. The number of edges is returned.
func dotGraph(name string) (*toulmin.Graph, int) {
	g := toulmin.NewGraph(name + ":.")
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: name + ":.", Concludes: "기호"})
	alnum, digits := &rule.Bits{Of: "sides", Mask: rule.SideAlnum}, &rule.Bits{Of: "sides", Mask: rule.SideDigit}
	switch name {
	case "마침표":
		g.Except(rule.MarkBoth).With(alnum).Attacks(claim)
		g.Except(rule.MarkRepeated).Attacks(claim)
	case "줄임표":
		g.Except(rule.MarkAlone).Attacks(claim)
		return g, 1
	case "파일명":
		g.Except(rule.MarkSides).With(alnum).Attacks(claim)
		g.Except(rule.ExtUnlisted).With(toulmin.NewStringListSpec("extensions", extensions)).Attacks(claim)
	case "도메인":
		g.Except(rule.MarkSides).With(alnum).Attacks(claim)
		g.Except(rule.LastUnlisted).With(toulmin.NewStringListSpec("topDomains", topDomains)).Attacks(claim)
	case "소수점":
		g.Except(rule.MarkSides).With(digits).Attacks(claim)
		g.Except(rule.PartsShape).With(&rule.Shape{Of: "소수점:.", Min: 2, Max: 2}).Attacks(claim)
	case "버전":
		g.Except(rule.MarkSides).With(alnum).Attacks(claim)
		g.Except(rule.VersionShape).Attacks(claim)
	}
	return g, 2
}
