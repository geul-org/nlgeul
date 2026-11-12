//ff:func feature=token type=rule control=iteration dimension=1
//ff:what addFused — the graphs of the fused letters (the honorific 시 under 셨·셔·십 is judged by the slot's own graph): the link vowel 으 (after a consonant other than ㄹ, before 시 or a closing that takes it as a piece), 잖 (지+않), and each fused form whose first word is fixed, which is judged as the particle under it is in a chain (and, for 거 written onto a modifying closing, by whether such a closing stands before it)
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// The counts of the pieces are those of their own entries; nothing is
// added to the total the shares are taken of.
func (ix *index) addFused() {
	rows := []*entry{
		{name: euName, form: "으", after: rule.AfterHard},
		{name: negName, form: string(negLetter)},
	}
	for _, e := range rows {
		g := toulmin.NewGraph(e.name)
		claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: e.name, Concludes: "표지"})
		except(g, rule.PieceIs, whyNoStem, fact(rule.PieceStem), not).Attacks(claim)
		if e.after != rule.AfterAny {
			except(g, rule.PieceAfter, whyAfter, &rule.Bits{Of: "after", Mask: e.after}).Attacks(claim)
			ix.edges++
		}
		if e.name == euName {
			except(g, rule.Listed, whyNotBefore, listed("takers", rule.ListAfter, append([]string{honorName}, linkTakers...)), not).Attacks(claim)
			ix.edges++
		}
		e.g = g
		ix.graphs, ix.edges = ix.graphs+1, ix.edges+1
	}
	ix.honor, ix.eu, ix.neg, ix.ji = ix.named('시', honorName), rows[0], rows[1], ix.named('지', jiName)
	ix.contracts = contracts
	if ix.lexed {
		ix.contracts = append(append(append([]contract{}, contracts...), shortPronouns...), typoContracts...)
	}
	for _, c := range ix.contracts {
		name := "줄임:" + c.form
		g := toulmin.NewGraph(name)
		claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: name, Concludes: "열쇠"})
		except(g, rule.ChainLinks, whyChain).Attacks(claim)
		if c.adn {
			except(g, rule.PieceIs, whyNoAdn, fact(rule.PieceAdn), not).Attacks(claim)
			ix.edges++
		}
		under := ix.named(c.key, c.name)
		ix.fixed = append(ix.fixed, &entry{name: name, form: c.form, class: classParticle, slot: under.slot, kase: under.kase, count: under.count, g: g})
		ix.graphs, ix.edges = ix.graphs+1, ix.edges+1
	}
}
