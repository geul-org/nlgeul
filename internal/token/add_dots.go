//ff:func feature=token type=rule control=iteration dimension=1
//ff:what addDots — a graph for each candidate of the full stop, a graph for each sign that sets two words side by side in one stem, and the graph of a run of Latin letters and digits that no table lists
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A run that no table lists is a word of its own: the claim stands,
// rebutted at the force forceRun by "it is in no key table".
func (ix *index) addDots() {
	for _, name := range joinNames {
		g, edges := dotGraph(name)
		ix.joins = append(ix.joins, &entry{name: name + ":.", g: g})
		ix.graphs, ix.edges = ix.graphs+1, ix.edges+edges
	}
	for _, name := range stopNames {
		g, edges := dotGraph(name)
		ix.stops = append(ix.stops, &entry{name: name, g: g})
		ix.graphs, ix.edges = ix.graphs+1, ix.edges+edges
	}
	ix.links = map[rune][]*entry{}
	for _, sign := range linkOrder {
		l := toulmin.NewGraph(linkSigns[sign])
		c := l.Rule(rule.PieceRead).With(&rule.Claim{Name: linkSigns[sign], Concludes: "기호"})
		except(l, rule.MarkBoth, whySides, &rule.Bits{Of: "sides", Mask: rule.SideAlnum | rule.SideHangul}, not).Attacks(c)
		ix.links[sign] = append(ix.links[sign], &entry{name: linkSigns[sign], g: l})
		ix.graphs, ix.edges = ix.graphs+1, ix.edges+1
	}
	g := toulmin.NewGraph(runName)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: runName, Concludes: "열쇠"})
	g.Counter(rule.PieceRead).With(&rule.Claim{Name: noKey, Concludes: "반박"}).Attacks(claim, forceRun)
	ix.graphs, ix.edges = ix.graphs+1, ix.edges+1
	ix.run = verdict(g, rule.Piece{})
}
