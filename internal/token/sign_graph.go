//ff:func feature=token type=rule control=selection
//ff:what signGraph — the judgment graph of one reading of a sign at a unit's edge: the claim "this sign is that", and the exceptions against it by where the sign stands and how many stand in a row; the number of edges is returned
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// 쉼표 is one comma. 물음표 and 느낌표 are any number of the sign. 쌍점
// and 쌍반점 are one sign at a word's end or alone; 물결표 any number
// there. 샵 and 골뱅이 are one sign with letters right after it or
// alone; 태그 and 이름 one sign with letters right after it. 제목 is the
// sign at the very start of an input line with a space right after it.
func signGraph(name string) (*toulmin.Graph, int) {
	g := toulmin.NewGraph(name)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: name, Concludes: "기호"})
	is := func(mask int) *rule.Bits { return &rule.Bits{Of: "mark", Mask: mask} }
	repeated := func() { except(g, rule.MarkIs, whySameSigns, is(rule.MarkRepeated)).Attacks(claim) }
	switch name {
	case "쉼표":
		repeated()
		return g, 1
	case "쌍점", "쌍반점":
		repeated()
		except(g, rule.MarkIs, whyNoEnd, is(rule.MarkEnds|rule.MarkAlone), not).Attacks(claim)
		return g, 2
	case "물결표":
		except(g, rule.MarkIs, whyNoEnd, is(rule.MarkEnds|rule.MarkAlone), not).Attacks(claim)
		return g, 1
	case "샵", "골뱅이":
		repeated()
		except(g, rule.MarkIs, whyNoLead, is(rule.MarkLeads|rule.MarkAlone), not).Attacks(claim)
		return g, 2
	case "태그", "이름":
		repeated()
		except(g, rule.MarkIs, whyNoLetters, is(rule.MarkLeads), not).Attacks(claim)
		return g, 2
	case headingName:
		except(g, rule.MarkIs, whyNoLineHead, is(rule.MarkHeads), not).Attacks(claim)
		except(g, rule.MarkIs, whyNoSpace, is(rule.MarkSpaced), not).Attacks(claim)
		return g, 2
	}
	return g, 0
}
