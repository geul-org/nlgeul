//ff:func feature=parse type=rule control=sequence
//ff:what dotsGraph — the judgment graph of a run of middle dots read as 줄임표: the claim "this run is that", the exception that the dot stands alone (no run), and the exception that it stands at no edge of its unit (between letters, where a run of two dots stays inside the word); the number of edges is returned
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// The orthography's appendix on punctuation writes 줄임표 as six dots
// at the middle of the line (……: the 1988 appendix gives it so, and the
// 2015 one keeps it first, with three dots at the middle or six or three
// at the foot as other ways), written onto the word before it: a run of
// middle dots at a word's end, right before the letters of its unit or
// alone is that sign, as a run of full stops is.
func dotsGraph() (*toulmin.Graph, int) {
	g := toulmin.NewGraph(ellipsisName)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: ellipsisName, Concludes: "기호"})
	is := func(mask int) *rule.Bits { return &rule.Bits{Of: "mark", Mask: mask} }
	except(g, rule.MarkIs, whyAlone, is(rule.MarkRepeated), not).Attacks(claim)
	except(g, rule.MarkIs, whyNotAtEdge, is(rule.MarkEnds|rule.MarkLeads|rule.MarkAlone), not).Attacks(claim)
	return g, 2
}
