//ff:func feature=token type=rule control=sequence
//ff:what aloneExcept — on the graph of the kind of words that only stand alone (interjections and the like), the two exceptions against its claim: a word ending in a modifying closing stands right before it, a determiner stands right before it; the number of edges is returned
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// What modifies a nominal is followed by one, or by another word that
// modifies it. An interjection is neither: it stands apart from the
// words around it.
func aloneExcept(g *toulmin.Graph, claim *toulmin.Rule) int {
	except(g, rule.PieceIs, whyAdn, fact(rule.PieceAdn)).Attacks(claim)
	except(g, rule.PieceIs, whyDet, fact(rule.PieceDet)).Attacks(claim)
	return 2
}
