//ff:func feature=token type=rule control=iteration dimension=1
//ff:what addPicked — the graph of each particle written onto a few closings only (pickedClosings), by its form: the particle read right after a closing, with the exception that the closing before it is not one of those the particle is written onto
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func (ix *index) addPicked() {
	ix.pickedAux = map[string]*entry{}
	for form, closings := range pickedClosings {
		name := pickedAuxName + form
		g := toulmin.NewGraph(name)
		claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: name, Concludes: "표지"})
		except(g, rule.Listed, whyNotAfter, listed("picked", rule.ListBefore, closings), not).Attacks(claim)
		ix.pickedAux[form] = &entry{name: name, g: g}
		ix.graphs, ix.edges = ix.graphs+1, ix.edges+1
	}
}
