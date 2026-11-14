//ff:func feature=token type=rule control=selection
//ff:what dotGraph — the judgment graph of one candidate of the full stop: the claim "this dot is that", and the structural exceptions against it
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// 마침표 ends a sentence: not between two runs, not one of dots in a
// row. 줄임표 is dots in a row. 파일명 joins a name and a listed
// extension, 도메인 the parts of a name whose last is a listed domain,
// 소수점 two numbers, 버전 numbers with a leading v or more dots than
// one. A closing candidate (마침표, 줄임표) is named with no sign after it.
// The number of edges is returned.
func dotGraph(name string) (*toulmin.Graph, int) {
	full := name + ":."
	if slices.Contains(stopNames, name) {
		full = name
	}
	g := toulmin.NewGraph(full)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: full, Concludes: "기호"})
	alnum, digits := &rule.Bits{Of: "sides", Mask: rule.SideAlnum}, &rule.Bits{Of: "sides", Mask: rule.SideDigit}
	switch name {
	case "마침표":
		except(g, rule.MarkBoth, whyBoth, alnum).Attacks(claim)
		except(g, rule.MarkIs, whyRepeated, &rule.Bits{Of: "mark", Mask: rule.MarkRepeated}).Attacks(claim)
	case "줄임표":
		except(g, rule.MarkIs, whyAlone, &rule.Bits{Of: "mark", Mask: rule.MarkRepeated}, not).Attacks(claim)
		return g, 1
	case "파일명":
		except(g, rule.MarkBoth, whySides, alnum, not).Attacks(claim)
		except(g, rule.Listed, whyExtension, listed("extensions", rule.ListExt, extensions), not).Attacks(claim)
	case "도메인":
		except(g, rule.MarkBoth, whySides, alnum, not).Attacks(claim)
		except(g, rule.Listed, whyDomain, listed("topDomains", rule.ListLast, topDomains), not).Attacks(claim)
	case "소수점":
		except(g, rule.MarkBoth, whySides, digits, not).Attacks(claim)
		except(g, rule.PartsShape, whyShape, &rule.Shape{Of: "소수점:.", Min: 2, Max: 2}).Attacks(claim)
	case "버전":
		except(g, rule.MarkBoth, whySides, alnum, not).Attacks(claim)
		except(g, rule.MarkIs, whyVersion, &rule.Bits{Of: "mark", Mask: rule.MarkNumeric}, not).Attacks(claim)
	}
	return g, 2
}
