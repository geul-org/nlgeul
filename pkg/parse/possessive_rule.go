//ff:func feature=parse type=rule control=sequence
//ff:what possessiveRule — the rule of a sentence "네 and 제 standing alone before a word a modifier may stand before are the pronoun with 의 fused": the candidate that reads one of the pronouns of the closed set (pronounWords; not 내, whose letters the contracts table reads as 나 with 의: genitiveForms, possessiveContractRule) as the nominal with no particle after it is preferred (bonusOverCount), but not when the word right after it is not one a modifier may stand before (unless it is an opening sign), nor when the word right after it is a unit (네 is then the numeral determiner), nor when it is written onto the word before it (unless what stands right before it is an opening sign); it waits for the word after it
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The dictionary gives 내, 네 and 제 as the pronouns 나, 너 and 저 with
// the adnominal particle 의 fused (나의, 너의, 저의), standing before the
// nominal they modify (내 책, 네 집, 제 생각). The word table counts 네
// far more often as the interjection of answer (네, 알겠습니다), which
// is an independent word and modifies nothing (school grammar: 감탄사는
// 독립어로 쓰인다), so 네 before a nominal fell to the interjection.
// Before a unit 네 is the numeral determiner (네 개, 네 명), which this
// rule leaves alone. A modifier is written apart from the word before it
// (clause 2 of the orthography). 내 is written out as 나 with the
// particle 의 by the contracts table, and the rule of that reading
// (possessiveContractRule) speaks of it in this one's place.
func possessiveRule() sentenceRule {
	g, claim := sentenceClaim(rulePossessive, rule.Listed, listed("대명사", rule.ListSelfStem, pronounWords))
	unmodifiable := except(g, rule.AroundIs, whyNoModifiableNext, &rule.Bits{Of: "around", Mask: rule.AroundBeforeModifiable}, not)
	unmodifiable.Attacks(claim)
	except(g, rule.Listed, whyOpenAfter, listed("여는 기호", rule.ListNextKind, []string{openName})).Attacks(unmodifiable)
	except(g, rule.Listed, whyUnitNext, listed("단위", rule.ListNextStem, unitWords)).Attacks(claim)
	glued := except(g, rule.AroundIs, whyGluedWord, &rule.Bits{Of: "around", Mask: rule.AroundGlued})
	glued.Attacks(claim)
	except(g, rule.Listed, whyOpenBefore, listed("여는 기호", rule.ListPrevKind, []string{openName})).Attacks(glued)
	is := func(n rule.Near) bool {
		return n.Kind == tableNoun && n.Marks == "" && slices.Contains(pronounWords, n.Stem) && !slices.Contains(genitiveForms, n.Stem)
	}
	return sentenceRule{name: rulePossessive, graph: g, bonus: bonusOverCount, about: is, waits: is}
}
