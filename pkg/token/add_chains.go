//ff:func feature=token type=rule control=iteration dimension=1
//ff:what addChains — the graph of each ending read right after a final closing of a quoted sentence (closingChains), by the ending's name: the ending read right after a closing, with the exception that the closing before it is not one of those it is read after
package token

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func (ix *index) addChains() {
	ix.chainAux = map[string]*entry{}
	for name, befores := range closingChains {
		label := chainAuxName + name
		g := toulmin.NewGraph(label)
		claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: label, Concludes: "표지"})
		except(g, rule.Listed, whyNotAfter, listed("chained", rule.ListBefore, strings.Split(befores, "|")), not).Attacks(claim)
		ix.chainAux[name] = &entry{name: label, g: g}
		ix.graphs, ix.edges = ix.graphs+1, ix.edges+1
	}
}
