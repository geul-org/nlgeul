//ff:func feature=token type=rule control=sequence
//ff:what adnExcept — on the graph of a predicate kind's stem, the exceptions against its claim that a word ending in a modifying closing stands right before the stem and that a determiner does, and against them what lifts each: after a modifying closing the listed auxiliaries (a dependent noun with 하다 or 싶다: 듯하, 만하), and after either a word that is itself a modifying or a naming form; the number of edges is returned
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// A modifying form modifies a nominal, so a predicate does not stand
// right after one in its unit (떨+ㄹ 어지+었+다); 듯하다, 만하다 and
// the like are written onto a modifying form and stand there. A word
// that is itself a modifying form modifies the same nominal (작은 예쁜
// 꽃), and a naming form stands as a nominal. A determiner modifies a
// nominal as a modifying form does (세 보+아 is no candidate of 세봐;
// 이 작은 꽃 stands).
func adnExcept(g *toulmin.Graph, claim *toulmin.Rule) int {
	after := except(g, rule.PieceIs, whyAdn, fact(rule.PieceAdn))
	after.Attacks(claim)
	except(g, rule.Listed, whyBoundAux, listed("boundAux", rule.ListForm, boundAux)).Attacks(after)
	except(g, rule.PieceIs, whyTurned, fact(rule.PieceTurned)).Attacks(after)
	det := except(g, rule.PieceIs, whyDet, fact(rule.PieceDet))
	det.Attacks(claim)
	except(g, rule.PieceIs, whyTurned, fact(rule.PieceTurned)).Attacks(det)
	return 5
}
