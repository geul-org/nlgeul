//ff:func feature=parse type=rule control=sequence
//ff:what apartGraph — the judgment graph of a sign that has no name but is also read as a word of its own between letters of a unit (bareApart: the equals sign of 값=3): the claim "this sign stands apart from the letters on its two sides", and the exceptions that the sign stands in a row, that both its sides are Latin letters or digits (a=1 is one name), that another sign stands right next to it, and that it stands at an edge of its unit (there it is sent up as a sign at an edge is); the number of edges is returned
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func apartGraph(sign rune) (*toulmin.Graph, int) {
	name := apartName + string(sign)
	g := toulmin.NewGraph(name)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: name, Concludes: "기호"})
	is := func(mask int) *rule.Bits { return &rule.Bits{Of: "mark", Mask: mask} }
	except(g, rule.MarkIs, whySameSigns, is(rule.MarkRepeated)).Attacks(claim)
	except(g, rule.MarkBoth, whyAlnumSides, &rule.Bits{Of: "sides", Mask: rule.SideAlnum}).Attacks(claim)
	except(g, rule.MarkIs, whyBySign, is(rule.MarkBySign)).Attacks(claim)
	except(g, rule.MarkIs, whyAtEdge, is(rule.MarkEnds|rule.MarkLeads|rule.MarkAlone)).Attacks(claim)
	return g, 4
}
