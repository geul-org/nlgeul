//ff:func feature=token type=rule control=sequence
//ff:what argumentAdnRule — the rule of a sentence "a predicate between an argument and the word it modifies is in its modifying form": a candidate that reads its stem as a verb, an adjective or an auxiliary verb (predicateStem) and whose last piece is a modifying closing (꾸밈: 는, 은, ㄴ, ㄹ …) is preferred, by bonusOverCount, when the word right before it ends in a case particle of an argument (argumentCases: 주격, 보격, 목적격, 부사격), or has 와/과 of the comparison with a particle more (comparedNames, rule.PrevHas: 그와는 다른), and not otherwise, nor when the word right after it is not one a modifier may stand before (not a nominal, a determiner, nor a modifying or naming form), nor when the word right after it ends in a modifying closing itself, nor when the candidate's letters are a numeral determiner's and the word right after it is a unit it counts (countedUnitExcept, countedUnits: 한 번); a word whose letters are a pronoun of the closed sets with a particle (pronounParticled: 나는, 이를) is not asked; it waits for the word after it
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// A nominal with a case particle of the subject, the complement, the
// object or an adverbial is an argument of a predicate after it (school
// grammar: 격 조사는 앞말이 서술어에 대하여 갖는 자격을 나타낸다), and a
// predicate in its modifying form modifies the nominal after it (관형절:
// 물을 파는 가게, 일을 한 사람). So a word right after such an argument and
// right before a nominal reads first as that predicate's modifying form
// (팔다+는, 하다+ㄴ) rather than as a nominal with the topic particle or a
// determiner of the same letters (파+는, 한), which the word table may
// count more often. When the word after it ends in a modifying closing,
// the argument is that word's own (이를 쓴 사람: 이를 is 쓰다's object),
// so the rule does not stand. The letters of the five demonstrative
// determiners that are also an adjective's modifying form (이런, 그런,
// 저런, 어떤, 다른) are asked too: where determinerFormRule holds the
// determiner back (다른 after the 와/과 다르다 takes: 이와 다른 전제),
// this rule is what carries the adjective's form, and elsewhere the
// determiner's bonusOverUncounted outweighs it. A particle after a case
// particle adds a sense and leaves the word what it is to the predicate
// (school grammar: 보조사), so that 와/과 of the comparison with a
// particle after it (그와는 다른) is still the comparison.
// A pronoun with a particle is an argument of its own, and an argument
// often stands right after another (피고가 이를 양도하였다, 그때에 나는
// 몰랐다): a word whose letters read so (pronounParticled) is not asked,
// lest 나다+는 or 이르다+(ㄹ) of the same letters be carried over it.
func argumentAdnRule() sentenceRule {
	adn := closingLabel(rule.KindAdn)
	g, claim := sentenceClaim(ruleArgumentAdn, rule.KindListed, listed("꾸밈 맺음", rule.ListSelfLast, []string{adn}))
	noArg := except(g, rule.KindListed, whyNoArgumentBefore, listed("격 조사", rule.ListPrevLast, argumentCases), not)
	noArg.Attacks(claim)
	except(g, rule.PrevHas, whyComparedBefore, listed("와·과", rule.ListPrevLast, comparedNames)).Attacks(noArg)
	except(g, rule.AroundIs, whyNoModifiableNext, &rule.Bits{Of: "around", Mask: rule.AroundBeforeModifiable}, not).Attacks(claim)
	except(g, rule.KindListed, whyAdnNext, listed("꾸밈 맺음", rule.ListNextLast, []string{adn})).Attacks(claim)
	countedUnitExcept(g, claim, numDeterminers, countedUnits)
	is := func(n rule.Near) bool {
		return labelOf(n.Last) == adn && predicateStem(n) && !pronounParticled(n.Text)
	}
	return sentenceRule{name: ruleArgumentAdn, graph: g, bonus: bonusOverCount, about: is, waits: is}
}
