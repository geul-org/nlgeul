//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what addWide — what is read only with a word table: the graphs of 신·실·심·세 as the honorific 시 (exceptions: no stem, not after a vowel, a particle right after (not for 심 and 세), a final consonant that makes no closing of the tables with what follows), the three graphs that ask the word before to be a predicate's stem, a nominal of the table, or a nominal of the table or a new word (nominalNew, before a dropped copula stem and a quoting closing), the old polite slots read right before a closing (politeSlots: 오, 옵, 사오, 사옵, 자오, 자옵, kept in polite and not among the keyed entries), the shortened closings, slots and particles (a shortened particle's word before must be a nominal of the table that makes no word of the table with it: joinedGateOf), the closings of the table of gated closings after a predicate's stem of the table (the naming closing 기, the 하오체 closings 소, 오, 으오 and 올시다 and the closing 마 of a promise, 으마 after a consonant), with the graph of a particle chain right after a naming closing (기, and 음 and ㅁ of the closing table: chainHeads), the modifying closings 은 and 을 after a stem's ㄹ, and where particles begin, and the suffixes read after a nominal (nounSuffixes: 률, 율, 당, 사) with the graph the word before them must pass: a nominal of the table that makes no word of the table with the suffix (suffixGateOf)
package parse

import (
	"sort"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func (ix *index) addWide() {
	ix.pred, ix.nominal, ix.nominalNew = gateOf(predGate, kindVerb|kindAdj), gateOf(nominalGate, kindNominal), nominalNewGateOf(nominalNewGate)
	ix.graphs, ix.edges = ix.graphs+3, ix.edges+4
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
		e := &entry{name: wideLetters[r], form: string(r), after: rule.AfterVowel, wide: true}
		g := toulmin.NewGraph(e.name)
		claim := g.Rule(rule.PieceRead).With(&rule.Claim{Name: e.name, Concludes: "표지"})
		except(g, rule.PieceIs, whyNoStem, fact(rule.PieceStem), not).Attacks(claim)
		except(g, rule.PieceAfter, whyAfter, &rule.Bits{Of: "after", Mask: e.after}).Attacks(claim)
		except(g, rule.PieceRight, whyRight, &rule.Bits{Of: "right", Mask: rule.RightOther}).Attacks(claim)
		if r != nounHonor && r != vowelHonor {
			except(g, rule.PieceRight, whyRight, &rule.Bits{Of: "right", Mask: rule.RightParticle}).Attacks(claim)
			ix.edges++
		}
		e.g, ix.wide[r] = g, e
		ix.graphs, ix.edges = ix.graphs+1, ix.edges+3
	}
	for _, row := range shortClosings {
		ix.addShort(row)
	}
	for _, m := range gatedClosings {
		for _, e := range closingEntries(m) {
			e.gate = ix.pred
			ix.add(e)
		}
	}
	for _, m := range looseClosings {
		e := closingEntries(m)[0]
		e.gate, e.loose = ix.pred, true
		ix.add(e)
	}
	noun := toulmin.NewGraph(nounAuxName)
	naming := noun.Rule(rule.PieceRead).With(&rule.Claim{Name: nounAuxName, Concludes: "표지"})
	except(noun, rule.Listed, whyNotAfter, listed("naming", rule.ListBefore, chainHeads), not).Attacks(naming)
	ix.nounAux, ix.graphs, ix.edges = &entry{name: nounAuxName, g: noun}, ix.graphs+1, ix.edges+1
	for i := range shortSlots {
		e := shortSlots[i]
		e.class, e.wide, e.gate = classSlot, true, ix.pred
		ix.add(&e)
	}
	for i := range politeSlots {
		e := politeSlots[i]
		ix.edges += markerGraph(&e, nil, caseExceptions[e.name], true)
		ix.polite, ix.graphs = append(ix.polite, &e), ix.graphs+1
	}
	short := joinedGateOf(nominalGate, kindNominal)
	ix.graphs, ix.edges = ix.graphs+1, ix.edges+2
	for i := range shortParticles {
		e := shortParticles[i]
		e.wide, e.gate = true, short
		ix.add(&e)
	}
	ix.suffixGate, ix.graphs, ix.edges = suffixGateOf(), ix.graphs+1, ix.edges+2
	for i := range nounSuffixes {
		e := nounSuffixes[i]
		e.gate = ix.suffixGate
		ix.add(&e)
	}
}
