//ff:func feature=parse type=rule control=sequence
//ff:what pairGraph — the judgment graph of one candidate of a sign of a pair: the claim "this sign opens a pair" or "this sign closes a pair", and the exceptions against it: for a closing candidate that no sign of its kind may stand open, for the opening candidate of a quotation mark written the same way at both ends that one of its kind may stand open and that a mark that ends a sentence stands right before it (다." — defeated when a word is written right after it: 다."그), and what the sign's row bars (it stands inside a word, apart with a space on both sides, or next to a hyphen or an equals sign); the number of edges is returned
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func pairGraph(sign rune, row pairRow, name string) (*toulmin.Graph, int) {
	full := name + ":" + string(sign)
	g := toulmin.NewGraph(full)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: full, Concludes: "기호"})
	edges := 0
	bar := func(r *toulmin.Rule) {
		r.Attacks(claim)
		edges++
	}
	opened := &rule.Bits{Of: "mark", Mask: rule.MarkOpened}
	if name == shutName {
		bar(except(g, rule.MarkIs, whyNotOpened, opened, not))
	}
	if name == openName && row.shuts {
		bar(except(g, rule.MarkIs, whyOpened, opened))
		ended := except(g, rule.MarkIs, whyEndsSentence, &rule.Bits{Of: "mark", Mask: rule.MarkAfterStop})
		bar(ended)
		except(g, rule.MarkRight, whyWordOn, &rule.Bits{Of: "sides", Mask: rule.SideHangul | rule.SideAlnum}, not).Attacks(ended)
		edges++
	}
	if row.bars&barWord != 0 {
		bar(except(g, rule.MarkBoth, whyInWord, &rule.Bits{Of: "sides", Mask: rule.SideAlnum}))
	}
	if row.bars&barApart != 0 {
		bar(except(g, rule.MarkIs, whyApart, &rule.Bits{Of: "mark", Mask: rule.MarkApart}))
		bar(except(g, rule.MarkIs, whyArrow, &rule.Bits{Of: "mark", Mask: rule.MarkArrow}))
	}
	return g, edges
}
