//ff:func feature=parse type=rule control=sequence
//ff:what possessiveContractRule — the rule of a sentence "내 standing alone before a word a modifier may stand before is 나 with 의 fused": the candidate that reads a form the contracts table puts the adnominal particle 의 under (genitiveForms: 내) as its pronoun with that particle (관형격:의, genitiveNames) is preferred (bonusOverCount), but not when the word right after it is not one a modifier may stand before (unless it is an opening sign), nor when it is written onto the word before it (unless what stands right before it is an opening sign), nor when the word right before it ends in a nominal (내 is then the noun of what is inside); it waits for the word after it
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The dictionary gives 내 as 나 with the adnominal particle 의 fused
// (나의), and a word with 의 modifies the nominal after it (내 책, 내
// 말, 내 생각에, 내 것): the form is written out as its pronoun and the
// particle, as 난 and 날 are. The same letters are other words where
// the grammar tells them apart. Before 가 the pronoun takes the form 내
// (내가), which is written onto its particle and is no 나의. The noun 내
// (what is inside) stands right after the nominal it bounds, with no
// particle between them (기간 내, 건물 내), while a modifier with 의
// begins a phrase of its own. The stem of the verb 내다 stands before
// a predicate (내 버리다), which no word with 의 modifies. A modifier
// is written apart from the word before it (clause 2 of the
// orthography).
func possessiveContractRule() sentenceRule {
	g, claim := sentenceClaim(rulePossessiveContract, rule.SelfHas, listed("관형격 의", rule.ListSelfMarks, genitiveNames))
	unmodifiable := except(g, rule.AroundIs, whyNoModifiableNext, &rule.Bits{Of: "around", Mask: rule.AroundBeforeModifiable}, not)
	unmodifiable.Attacks(claim)
	except(g, rule.Listed, whyOpenAfter, listed("여는 기호", rule.ListNextKind, []string{openName})).Attacks(unmodifiable)
	glued := except(g, rule.AroundIs, whyGluedWord, &rule.Bits{Of: "around", Mask: rule.AroundGlued})
	glued.Attacks(claim)
	except(g, rule.Listed, whyOpenBefore, listed("여는 기호", rule.ListPrevKind, []string{openName})).Attacks(glued)
	except(g, rule.AroundIs, whyAfterNominal, &rule.Bits{Of: "around", Mask: rule.AroundAfterNominal}).Attacks(claim)
	is := func(n rule.Near) bool {
		return n.Kind == tableNoun && slices.Contains(genitiveNames, n.Marks) && slices.Contains(genitiveForms, n.Text)
	}
	return sentenceRule{name: rulePossessiveContract, graph: g, bonus: bonusOverCount, about: is, waits: is}
}
