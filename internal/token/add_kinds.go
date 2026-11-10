//ff:func feature=token type=rule control=iteration dimension=2
//ff:what addKinds — a graph for each kind of stem (열쇠:체언, 열쇠:용언 …): the claim "this stem is a known word", and the exception that its kind does not take what follows; their verdicts are worked out once for every kind and tail
package token

import (
	"github.com/geul-org/koparse/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// Words are rows of a table, not graphs: a stem found in the table
// hands its kind to one of these few graphs.
func (ix *index) addKinds() {
	ix.fits = map[int]map[int]float64{}
	for _, kind := range kindOrder {
		name := "열쇠:" + kindLabel[kind]
		g := toulmin.NewGraph(name)
		claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: name, Concludes: "열쇠"})
		g.Except(rule.StemTail).With(&rule.Bits{Of: "tail", Mask: kindTails[kind]}).Attacks(claim)
		ix.graphs, ix.edges = ix.graphs+1, ix.edges+1
		ix.fits[kind] = map[int]float64{}
		for _, tail := range []int{rule.TailBare, rule.TailAux, rule.TailCase, rule.TailPred} {
			ix.fits[kind][tail], _ = verdictOf(g, rule.Piece{Tail: tail}, "")
		}
	}
}
