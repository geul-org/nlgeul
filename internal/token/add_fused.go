//ff:func feature=token type=rule control=iteration dimension=1
//ff:what addFused — the graphs of the fused letters: the honorific 시 under 셨·셔·십 (after a vowel), the link vowel 으 before it (after a consonant other than ㄹ), 잖 (지+않), and each fused form whose first word is fixed, which is judged as the particle under it is in a chain (and, for 거 written onto a modifying closing, by whether such a closing stands before it)
package token

import (
	"github.com/geul-org/koparse/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// The counts of the pieces are those of their own entries; nothing is
// added to the total the shares are taken of.
func (ix *index) addFused() {
	rows := []*entry{
		{name: honorName, form: "시", after: afterVowel, count: honorCount},
		{name: euName, form: "으", after: rule.AfterHard},
		{name: negName, form: string(negLetter)},
	}
	for _, e := range rows {
		g := toulmin.NewGraph(e.name)
		claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: e.name, Concludes: "표지"})
		g.Except(rule.PieceNoStem).Attacks(claim)
		if e.after != afterAny {
			g.Except(rule.PieceAfter).With(&rule.Bits{Of: "after", Mask: e.after}).Attacks(claim)
			ix.edges++
		}
		e.g = g
		ix.graphs, ix.edges = ix.graphs+1, ix.edges+1
	}
	ix.honor, ix.eu, ix.neg, ix.ji = rows[0], rows[1], rows[2], ix.named('지', jiName)
	for _, c := range contracts {
		name := "줄임:" + c.form
		g := toulmin.NewGraph(name)
		claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: name, Concludes: "열쇠"})
		g.Except(rule.ChainLinks).Attacks(claim)
		if c.adn {
			g.Except(rule.PieceNoAdn).Attacks(claim)
			ix.edges++
		}
		under := ix.named(c.key, c.name)
		ix.fixed = append(ix.fixed, &entry{name: name, form: c.form, class: classParticle, slot: under.slot, kase: under.kase, count: under.count, g: g})
		ix.graphs, ix.edges = ix.graphs+1, ix.edges+1
	}
}
