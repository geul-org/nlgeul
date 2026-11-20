//ff:func feature=parse type=rule control=sequence
//ff:what unitWordRebuttal — on the graph of a kind of stem, the rebuttal of its claim that the stem begins right after a number and its letters are a unit and a word written after a unit (5년간, 6월경: the number's unit is read into the word); the number of edges is returned
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func unitWordRebuttal(g *toulmin.Graph, claim *toulmin.Rule) int {
	g.Counter(rule.PieceIs).With(fact(rule.PieceUnitWord)).Attacks(claim, forceUnitWord)
	return 1
}
