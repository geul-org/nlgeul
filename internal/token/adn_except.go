//ff:func feature=token type=rule control=sequence
//ff:what adnExcept — on the graph of a predicate kind's stem, the exception against its claim that a word ending in a modifying closing stands right before the stem, and against that exception the one that lifts it for the listed auxiliaries (a dependent noun with 하다 or 싶다: 듯하, 만하); the number of edges is returned
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A modifying form modifies a nominal, so a predicate does not stand
// right after one in its unit (떨+ㄹ 어지+었+다); 듯하다, 만하다 and
// the like are written onto a modifying form and stand there.
func adnExcept(g *toulmin.Graph, claim *toulmin.Rule) int {
	after := g.Except(rule.PieceAdn)
	after.Attacks(claim)
	g.Except(rule.FormListed).With(toulmin.NewStringListSpec("boundAux", boundAux)).Attacks(after)
	return 2
}
