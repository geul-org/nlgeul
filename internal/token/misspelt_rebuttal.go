//ff:func feature=token type=rule control=sequence
//ff:what misspeltRebuttal — on the graph of a kind of stem, the rebuttal of its claim that the stem's letters are a wrong way to write another word (no exception lifts it); the number of edges is returned
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func misspeltRebuttal(g *toulmin.Graph, claim *toulmin.Rule) int {
	g.Counter(rule.PieceIs).With(fact(rule.PieceMisspelt)).Attacks(claim, forceTypo)
	return 1
}
