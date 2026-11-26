//ff:func feature=parse type=rule control=sequence
//ff:what possessiveContractRule — the rule of a sentence "내, 네 and 제 standing alone before a word a modifier may stand before are 나, 너 and 저 with 의 fused": the candidate that reads a form the contracts table puts the adnominal particle 의 under (genitiveForms: 내, 네, 제) as its pronoun with that particle (관형격:의, genitiveNames) is preferred (bonusOverCount), but not when the word right after it is not one a modifier may stand before (unless it is an opening sign), nor when the word right after it is a unit and the form is a numeral determiner as well (genitiveNumerals: 네 개; 내 before a noun that is also a unit stays 나 with 의), nor when it is written onto the word before it (unless what stands right before it is an opening sign), nor when the word right before it ends in a nominal (내 is then the noun of what is inside, 제 a noun of a system), nor, for 제 (genitiveOrdinals), when the word right after it is a number (the prefix of an ordinal: 제 3) or the word right before it ends in the modifying ㄹ or 을 (the dependent noun of time: 할 제); it waits for the word after it
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

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
//
// 네 is 너 with 의 fused in the same way (너의: 네 책, 네 의견). The
// interjection of answer of the same letters is an independent word
// and modifies nothing, so the rule that carries an interjection
// standing alone does not ask about these letters and this rule
// carries the pronoun before a word a modifier may stand before.
// Before 가 the pronoun takes the form 네 (네가), written onto its
// particle. Before a unit 네 is the numeral determiner (네 개, 네 명),
// which this rule leaves alone; 내 is no numeral, so a noun after it
// that is a unit as well (내 배, 내 번호) does not take the reading
// away. Before a predicate (네 듣다) no word
// with 의 stands.
//
// 제 is 저 with 의 fused (저의: 제 생각, and the reflexive 제 눈으로,
// both written with the one pronoun 저). The same letters are the
// prefix of an ordinal, which stands before a number (제 3 자, 제 12
// 대; written onto the number it is one word with it and its letters
// are not 제 alone), the dependent noun of time, which stands after the
// modifying -ㄹ/-을 (할 제: the dictionary's 제, 적에 shortened), and a
// noun or a suffix of a system, which stands right after its nominal
// with no particle or is written onto it. Before 가 the pronoun takes
// the form 제 (제가), written onto its particle. These two exceptions
// speak of 제 alone: 내 and 네 before a number or after -ㄹ stay the
// pronoun with 의 (내 3년, 갈 내 길).
func possessiveContractRule() sentenceRule {
	g, claim := sentenceClaim(rulePossessiveContract, rule.SelfHas, listed("관형격 의", rule.ListSelfMarks, genitiveNames))
	unmodifiable := except(g, rule.AroundIs, whyNoModifiableNext, &rule.Bits{Of: "around", Mask: rule.AroundBeforeModifiable}, not)
	unmodifiable.Attacks(claim)
	except(g, rule.Listed, whyOpenAfter, listed("여는 기호", rule.ListNextKind, []string{openName})).Attacks(unmodifiable)
	unit := except(g, rule.Listed, whyUnitNext, listed("단위", rule.ListNextStem, unitWords))
	unit.Attacks(claim)
	except(g, rule.Listed, whyNotNumDet, listed("수 관형사", rule.ListSelfText, genitiveNumerals), not).Attacks(unit)
	glued := except(g, rule.AroundIs, whyGluedWord, &rule.Bits{Of: "around", Mask: rule.AroundGlued})
	glued.Attacks(claim)
	except(g, rule.Listed, whyOpenBefore, listed("여는 기호", rule.ListPrevKind, []string{openName})).Attacks(glued)
	except(g, rule.AroundIs, whyAfterNominal, &rule.Bits{Of: "around", Mask: rule.AroundAfterNominal}).Attacks(claim)
	number := except(g, rule.Listed, whyNumberNext, listed("수", rule.ListNextKind, numberKinds))
	number.Attacks(claim)
	except(g, rule.Listed, whyNotOrdinal, listed("차례의 제", rule.ListSelfText, genitiveOrdinals), not).Attacks(number)
	time := except(g, rule.Listed, whyFutureAdnBefore, listed("꾸밈 ㄹ·을", rule.ListPrevLast, futureAdnNames))
	time.Attacks(claim)
	except(g, rule.Listed, whyNotOrdinal, listed("차례의 제", rule.ListSelfText, genitiveOrdinals), not).Attacks(time)
	return sentenceRule{name: rulePossessiveContract, graph: g, bonus: bonusOverCount, about: genitiveContract, waits: genitiveContract}
}
