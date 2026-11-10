//ff:func feature=token type=reader control=iteration dimension=2
//ff:what addClosings — the entries of every closing of the table and of this layer's own list; the names of those the copula does not take are kept for its graph, and every form for telling a closing from a closing and a particle; and the graph of an auxiliary particle right after a linking closing (보+지+도)
package token

import (
	"github.com/geul-org/koparse/internal/parse"
	"github.com/geul-org/koparse/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func (ix *index) addClosings() {
	ix.closingSet = map[string]bool{}
	g := toulmin.NewGraph(linkAuxName)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: linkAuxName, Concludes: "표지"})
	g.Except(rule.FormUnlisted).With(toulmin.NewStringListSpec("linking", auxClosings)).Attacks(claim)
	g.Except(rule.NextUnlisted).With(toulmin.NewStringListSpec("particles", auxParticles)).Attacks(claim)
	ix.linkAux, ix.graphs, ix.edges = &entry{name: linkAuxName, g: g}, ix.graphs+1, ix.edges+2
	for _, m := range append(parse.ClosingMarkers(), moreClosings...) {
		for _, e := range closingEntries(m) {
			ix.add(e)
			ix.closingSet[rule.LeadText(e.lead)+e.form] = true
			if e.bars {
				ix.barred = append(ix.barred, e.name)
			}
		}
	}
}
