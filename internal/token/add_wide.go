//ff:func feature=token type=rule control=iteration dimension=1
//ff:what addWide — what is read only with a word table: the graphs of 신·실·심·세 as the honorific 시 (exceptions: no stem, not after a vowel, a particle right after (not for 심 and 세), a final consonant that makes no closing of the tables with what follows), the two graphs that ask the word before to be a predicate's stem or a nominal of the table, the shortened closings, slots and particles, the naming closing 기 with the graph of a particle chain right after it, and where particles begin
package token

import (
	"sort"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func (ix *index) addWide() {
	ix.pred, ix.nominal = gateOf(predGate, kindVerb|kindAdj), gateOf(nominalGate, kindNominal)
	ix.graphs, ix.edges = ix.graphs+2, ix.edges+2
	ix.heads = map[rune][]string{}
	for _, list := range ix.byKey {
		ix.addHeads(list)
	}
	letters := make([]rune, 0, len(wideLetters))
	for r := range wideLetters {
		letters = append(letters, r)
	}
	sort.Slice(letters, func(a, b int) bool { return letters[a] < letters[b] })
	ix.wide = map[rune]*entry{}
	for _, r := range letters {
		e := &entry{name: wideLetters[r], form: string(r), after: afterVowel, wide: true}
		g := toulmin.NewGraph(e.name)
		claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: e.name, Concludes: "표지"})
		g.Except(rule.PieceNoStem).Attacks(claim)
		g.Except(rule.PieceAfter).With(&rule.Bits{Of: "after", Mask: e.after}).Attacks(claim)
		g.Except(rule.PieceRight).With(&rule.Bits{Of: "right", Mask: rule.RightOther}).Attacks(claim)
		if r != nounHonor && r != vowelHonor {
			g.Except(rule.PieceRight).With(&rule.Bits{Of: "right", Mask: rule.RightParticle}).Attacks(claim)
			ix.edges++
		}
		e.g, ix.wide[r] = g, e
		ix.graphs, ix.edges = ix.graphs+1, ix.edges+3
	}
	for _, row := range shortClosings {
		ix.addShort(row)
	}
	for _, m := range gatedClosings {
		e := closingEntries(m)[0]
		e.gate = ix.pred
		ix.add(e)
	}
	noun := toulmin.NewGraph(nounAuxName)
	naming := noun.Rule(rule.PieceRead).With(&rule.Claim{Name: nounAuxName, Concludes: "표지"})
	noun.Except(rule.FormUnlisted).With(toulmin.NewStringListSpec("naming", nounHeads)).Attacks(naming)
	ix.nounAux, ix.graphs, ix.edges = &entry{name: nounAuxName, g: noun}, ix.graphs+1, ix.edges+1
	for i := range shortSlots {
		e := shortSlots[i]
		e.class, e.wide, e.gate = classSlot, true, ix.pred
		ix.add(&e)
	}
	for i := range shortParticles {
		e := shortParticles[i]
		e.wide, e.gate = true, ix.nominal
		ix.add(&e)
	}
}
