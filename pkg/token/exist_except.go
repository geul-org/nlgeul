//ff:func feature=token type=rule control=sequence
//ff:what existExcept — on the graph of an adjective's stem, against the exception that its kind does not take what follows, what lifts it: the stem is one of 있다, 없다 and 계시다 (existStems); and against that, that the closing after it is not one those stems take (existClosings: 느냐); the number of edges is returned
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// The dictionary gives 느냐 after the stems of 있다, 없다 and 계시다 as
// after a verb's (없느냐), and after no other adjective's (좋으냐).
func existExcept(g *toulmin.Graph, tail *toulmin.Rule) int {
	stem := except(g, rule.Listed, whyExistStem, listed("existStems", rule.ListForm, existStems))
	stem.Attacks(tail)
	except(g, rule.StemTail, whyNotExist, &rule.Bits{Of: "tail", Mask: rule.TailExist}).Attacks(stem)
	return 2
}
