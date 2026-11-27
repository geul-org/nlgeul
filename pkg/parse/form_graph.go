//ff:func feature=parse type=rule control=selection
//ff:what formGraph — the judgment graph of one name of a sign of the tables of special letters: the claim "this letter is that"; a name of the hand-written table or of the table of kinds (drawn told) does not stand where a sign other than itself that is no punctuation mark of a sentence stands right next to it (=>, |---|, ├──, ★☆: the signs draw something together and stay letters with no name; a punctuation mark between them does not join them: ○○·△△); by the name's kind, 가림 (a sign written for what is kept back) is lowered by a rebuttal where nothing in its place tells it — the sign neither stands two or more in a row nor is written onto letters (one ○ alone) — and 글머리 (a shape that marks an item) is one sign at the very start of an input line with a space right after it, and nothing else; the number of edges is returned
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func formGraph(name string, drawn bool) (*toulmin.Graph, int) {
	is := func(mask int) *rule.Bits { return &rule.Bits{Of: "mark", Mask: mask} }
	g := toulmin.NewGraph(name)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: name, Concludes: "기호"})
	edges := 0
	if drawn {
		except(g, rule.MarkIs, whyTouched, is(rule.MarkTouched)).Attacks(claim)
		edges++
	}
	switch nameOf(name).Kind {
	case hideKind:
		g.Counter(rule.MarkIs).With(is(rule.MarkRepeated|rule.MarkEnds|rule.MarkLeads)).With(not).Attacks(claim, forceNoHide)
		return g, edges + 1
	case bulletKind:
		except(g, rule.MarkIs, whySameSigns, is(rule.MarkRepeated)).Attacks(claim)
		except(g, rule.MarkIs, whyNoLineHead, is(rule.MarkHeads), not).Attacks(claim)
		except(g, rule.MarkIs, whyNoSpace, is(rule.MarkSpaced), not).Attacks(claim)
		return g, edges + 3
	}
	return g, edges
}
