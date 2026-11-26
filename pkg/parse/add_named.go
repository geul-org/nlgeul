//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what addNamed — the graph of each particle read right after the closings its row lists (closingParticles), by the particle's name: the particle read right after a closing, with the exception that the closing before it is not one of those listed; and the listed closings of each, by the same name
package parse

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func (ix *index) addNamed() {
	ix.namedAux, ix.namedEnds = map[string]*entry{}, map[string][]string{}
	for name, befores := range closingParticles {
		label := namedAuxName + name
		g := toulmin.NewGraph(label)
		claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: label, Concludes: "표지"})
		ix.namedEnds[name] = strings.Split(befores, "|")
		except(g, rule.Listed, whyNotAfter, listed("named", rule.ListBefore, ix.namedEnds[name]), not).Attacks(claim)
		ix.namedAux[name] = &entry{name: label, g: g}
		ix.graphs, ix.edges = ix.graphs+1, ix.edges+1
	}
}
