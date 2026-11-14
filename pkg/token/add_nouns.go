//ff:func feature=token type=rule control=iteration dimension=1
//ff:what addNouns — a graph for each dependent noun the copula's stem is fused into (nounCopulas: 테 = 터+이): the claim "this is the noun with the stem of 이다 in it", and the exceptions that no word ending in a modifying closing stands before it that the closing after it is one the copula does not take, and that what follows is 어/아 or ㅆ fused into its own vowel (터이+어 is not 테)
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// The count is the copula's, which is not in the total the other
// markers' shares are taken of.
func (ix *index) addNouns() {
	fused := ix.fusedNames()
	for _, c := range nounCopulas {
		name := "줄임:" + c.form
		g := toulmin.NewGraph(name)
		claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: name, Concludes: "열쇠"})
		except(g, rule.PieceIs, whyNoAdn, fact(rule.PieceAdn), not).Attacks(claim)
		except(g, rule.Listed, whyBarred, listed("barred", rule.ListAfter, ix.barred)).Attacks(claim)
		except(g, rule.Listed, whyFusedIn, listed("fused", rule.ListAfter, fused)).Attacks(claim)
		ix.nouns = append(ix.nouns, &entry{name: name, form: c.form, count: copulaCount, g: g})
		ix.graphs, ix.edges = ix.graphs+1, ix.edges+3
	}
}
