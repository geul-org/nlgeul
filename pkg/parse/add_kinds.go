//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what addKinds — a graph for each kind of stem (열쇠:체언, 열쇠:동사 …): the claim "this stem is a known word", the exception that its kind does not take what follows (for an adjective's stem lifted for 있다, 없다 and 계시다 before 느냐: existExcept), the rebuttals that it was rarely seen (fewer times than seenEnough) and that it is one syllable long with what lifts the second (kindRebuttals), the rebuttal that its letters are a wrong way to write another word (misspeltRebuttal), the rebuttal that it begins right after a number with the letters of a unit and a word written after a unit (unitWordRebuttal), for a verb's or an adjective's stem the exceptions that a word ending in a modifying closing or a determiner stands right before it (adnExcept), and for a word that only stands alone the same two with nothing to lift them (aloneExcept)
package parse

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
		tail := except(g, rule.StemTail, whyStemTail, &rule.Bits{Of: "tail", Mask: kindTails[kind]})
		tail.Attacks(claim)
		if kind == kindAdj {
			ix.edges += existExcept(g, tail)
		}
		ix.kinds[kind] = g
		ix.graphs, ix.edges = ix.graphs+1, ix.edges+1+kindRebuttals(g, claim, ix.seenEnough())+misspeltRebuttal(g, claim)+unitWordRebuttal(g, claim)
		if kind&(kindVerb|kindAdj) != 0 {
			ix.edges += adnExcept(g, claim)
		}
		if kind == kindAlone {
			ix.edges += aloneExcept(g, claim)
		}
	}
}
