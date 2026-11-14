//ff:func feature=token type=rule control=iteration dimension=1
//ff:what addKinds — a graph for each kind of stem (열쇠:체언, 열쇠:동사 …): the claim "this stem is a known word", the exception that its kind does not take what follows, the rebuttals that it was rarely seen and that it is one syllable long with what lifts each (kindRebuttals), the rebuttal that its letters are a wrong way to write another word (misspeltRebuttal), for a verb's or an adjective's stem the exceptions that a word ending in a modifying closing or a determiner stands right before it (adnExcept), and for a word that only stands alone the same two with nothing to lift them (aloneExcept)
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// Words are rows of a table, not graphs: a stem found in the table
// hands its kind, count and length to one of these few graphs.
func (ix *index) addKinds() {
	ix.kinds = map[int]*toulmin.Graph{}
	for _, kind := range kindOrder {
		name := "열쇠:" + kindLabel[kind]
		g := toulmin.NewGraph(name)
		claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: name, Concludes: "열쇠"})
		except(g, rule.StemTail, whyStemTail, &rule.Bits{Of: "tail", Mask: kindTails[kind]}).Attacks(claim)
		ix.kinds[kind] = g
		ix.graphs, ix.edges = ix.graphs+1, ix.edges+1+kindRebuttals(g, claim)+misspeltRebuttal(g, claim)
		if kind&(kindVerb|kindAdj) != 0 {
			ix.edges += adnExcept(g, claim)
		}
		if kind == kindAlone {
			ix.edges += aloneExcept(g, claim)
		}
	}
}
