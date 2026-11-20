//ff:func feature=parse type=rule control=sequence
//ff:what numberAdnRebuttal — on the graph of the copula's stem 이, the rebuttal of its claim that the modifying closing ㄴ follows it (numberAdnNames: 이 and ㄴ written 인), with the exception that lifts it where no number stands right before the stem (1인, 300인: the dependent noun 인 that counts people comes first); the number of edges is returned
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// The standard dictionary has 인 (人) as a dependent noun that counts
// people (5인 가족); right after a number the letters 인 are that noun
// rather than the copula's modifying form, which needs a nominal before
// it that the number stands for (크기가 3인 집합). The rebuttal lowers the
// copula and keeps it as a candidate.
func numberAdnRebuttal(g *toulmin.Graph, claim *toulmin.Rule) int {
	adn := g.Counter(rule.Listed).With(listed("수 뒤 꾸밈 ㄴ", rule.ListAfter, numberAdnNames))
	adn.Attacks(claim, forceNumberAdn)
	except(g, rule.PieceIs, whyNoNumber, fact(rule.PieceNum), not).Attacks(adn)
	return 2
}
