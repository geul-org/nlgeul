//ff:func feature=parse type=rule control=sequence
//ff:what numberNounRebuttal — on the graph of the copula's dropped stem, the rebuttal of its claim that a closing of numberNounClosings (면, 군) follows it, with the exception that lifts it where no number stands right before the stem (3면 the page and 1군 the army are nouns written after a number); the number of edges is returned
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// The dictionary has 면 (a page, a side) and 군 (an army, a county) as
// nouns that stand after a number, so right after a number these
// letters are read as that noun before the copula's dropped stem with
// the closing (2라면, 3다 take the stem: no noun is written so). The
// rebuttal lowers the stem and keeps it as a candidate.
func numberNounRebuttal(g *toulmin.Graph, claim *toulmin.Rule) int {
	noun := g.Counter(rule.Listed).With(listed("수 뒤 이름 맺음", rule.ListAfter, numberNounClosings))
	noun.Attacks(claim, forceNumberNoun)
	except(g, rule.PieceIs, whyNoNumber, fact(rule.PieceNum), not).Attacks(noun)
	return 2
}
