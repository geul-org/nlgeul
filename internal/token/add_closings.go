//ff:func feature=token type=reader control=iteration dimension=2
//ff:what addClosings — the entries of every closing of the table and of this layer's own list; the names of those the copula does not take are kept for its graph, and every form, noted before any entry is put in, for telling a closing from a closing and a particle and a closing that takes 으 from one that does not; and the graph of an auxiliary particle right after a linking closing (보+지+도), with the exception that the closing before it is not one of those; and the graph of a particle right after a closing that asks (맞+는지+를), with the exceptions that the closing before it is not one of those and that the particle is not one such a closing takes
package token

import (
	"github.com/geul-org/nlgeul/internal/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func (ix *index) addClosings() {
	ix.closingSet = map[string]bool{}
	g := toulmin.NewGraph(linkAuxName)
	claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: linkAuxName, Concludes: "표지"})
	except(g, rule.Listed, whyNotAfter, listed("linking", rule.ListBefore, auxClosings), not).Attacks(claim)
	ix.linkAux, ix.graphs, ix.edges = &entry{name: linkAuxName, g: g}, ix.graphs+1, ix.edges+1
	ask := toulmin.NewGraph(askAuxName)
	asked := ask.Rule(rule.PieceRead).With(&rule.Claim{Name: askAuxName, Concludes: "표지"})
	except(ask, rule.Listed, whyNotAfter, listed("asking", rule.ListBefore, askClosings), not).Attacks(asked)
	except(ask, rule.Listed, whyNotTaken, listed("taken", rule.ListForm, askParticles), not).Attacks(asked)
	ix.askAux, ix.graphs, ix.edges = &entry{name: askAuxName, g: ask}, ix.graphs+1, ix.edges+2
	rows := append(parse.ClosingMarkers(), moreClosings...)
	for _, m := range rows {
		for _, e := range closingEntries(m) {
			ix.closingSet[rule.LeadText(e.lead)+e.form] = true
		}
	}
	for _, m := range rows {
		for _, e := range closingEntries(m) {
			ix.add(e)
			if e.bars {
				ix.barred = append(ix.barred, e.name)
			}
		}
	}
}
