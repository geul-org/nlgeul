//ff:func feature=parse type=rule control=selection
//ff:what dashGraph — the judgment graph of one candidate (줄표, 빼기) of a sign that is written for a dash and drawn with too (the hyphen in a row of any length, the box-drawing line ─): the claim, and the exceptions of its place: its input line holds no letter or digit (a rule, a table's line), a sign that is no punctuation mark of a sentence stands in its run (-->, |---|, ├──┼); for the hyphen also that it does not stand alone with a space on each side (defeated for 줄표 when two or more stand in a row) and that one hyphen alone begins its input line (a list's mark); 빼기 is one hyphen only; where a sum is written 줄표 is lowered by a rebuttal, elsewhere 빼기 is (forceDashSum); the number of edges is returned
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func dashGraph(sign rune, name string) (*toulmin.Graph, int) {
	g := toulmin.NewGraph(name)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: name, Concludes: "기호"})
	is := func(mask int) *rule.Bits { return &rule.Bits{Of: "mark", Mask: mask} }
	except(g, rule.MarkIs, whyBareLine, is(rule.MarkBare)).Attacks(claim)
	except(g, rule.MarkIs, whyDrawn, is(rule.MarkDrawn)).Attacks(claim)
	switch {
	case sign != '-':
		return g, 2
	case name == minusName:
		except(g, rule.MarkIs, whyNotAlone, is(rule.MarkApart), not).Attacks(claim)
		except(g, rule.MarkIs, whySameSigns, is(rule.MarkRepeated)).Attacks(claim)
		except(g, rule.MarkIs, whyBullet, is(rule.MarkHeads)).Attacks(claim)
		g.Counter(rule.MarkIs).With(is(rule.MarkSum)).With(not).Attacks(claim, forceDashSum)
		return g, 6
	}
	lone := except(g, rule.MarkIs, whyNotAlone, is(rule.MarkApart), not)
	lone.Attacks(claim)
	except(g, rule.MarkIs, whyRepeated, is(rule.MarkRepeated)).Attacks(lone)
	bullet := except(g, rule.MarkIs, whyBullet, is(rule.MarkHeads))
	bullet.Attacks(claim)
	except(g, rule.MarkIs, whyRepeated, is(rule.MarkRepeated)).Attacks(bullet)
	g.Counter(rule.MarkIs).With(is(rule.MarkSum)).Attacks(claim, forceDashSum)
	return g, 7
}
