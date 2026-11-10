//ff:func feature=token type=rule control=iteration dimension=1
//ff:what addKinds — a graph for each kind of stem (열쇠:체언, 열쇠:동사 …): the claim "this stem is a known word", the exception that its kind does not take what follows, and the rebuttals that it was rarely seen and that it is one syllable long
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// Words are rows of a table, not graphs: a stem found in the table
// hands its kind, count and length to one of these few graphs.
func (ix *index) addKinds() {
	ix.kinds, ix.seen = map[int]*toulmin.Graph{}, map[[4]int]float64{}
	for _, kind := range kindOrder {
		name := "열쇠:" + kindLabel[kind]
		g := toulmin.NewGraph(name)
		claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: name, Concludes: "열쇠"})
		g.Except(rule.StemTail).With(&rule.Bits{Of: "tail", Mask: kindTails[kind]}).Attacks(claim)
		g.Counter(rule.SeenFew).With(&rule.Bits{Of: "enough", Mask: seenEnough}).Attacks(claim, forceFew)
		g.Counter(rule.ShortKey).With(&rule.Bits{Of: "syllables", Mask: shortKey}).Attacks(claim, forceShort)
		ix.kinds[kind] = g
		ix.graphs, ix.edges = ix.graphs+1, ix.edges+3
	}
}
