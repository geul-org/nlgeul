//ff:func feature=token type=rule control=sequence
//ff:what argumentAdnRule — the rule of a sentence "a predicate between an argument and the word it modifies is in its modifying form": a candidate that reads its stem as a verb, an adjective or an auxiliary verb (predicateStem) and whose last piece is a modifying closing (꾸밈: 는, 은, ㄴ, ㄹ …), not one whose letters are a demonstrative determiner of formDeterminers (그런, 다른 …: determinerFormRule judges them), is preferred, by bonusOverCount, when the word right before it ends in a case particle of an argument (argumentCases: 주격, 보격, 목적격, 부사격), and not otherwise, nor when the word right after it is not one a modifier may stand before (not a nominal, a determiner, nor a modifying or naming form), nor when the word right after it ends in a modifying closing itself; it waits for the word after it
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

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
// so the rule does not stand. The five demonstrative determiners whose
// letters are also an adjective's modifying form (이런, 그런, 저런, 어떤,
// 다른: formDeterminers) are not asked: a determiner stands before a
// nominal just as the modifying form does, and determinerFormRule tells
// the two apart by what does tell them apart (the 와/과 다르다 takes).
func argumentAdnRule() sentenceRule {
	adn := closingLabel(rule.KindAdn)
	g, claim := sentenceClaim(ruleArgumentAdn, rule.KindListed, listed("꾸밈 맺음", rule.ListSelfLast, []string{adn}))
	except(g, rule.KindListed, whyNoArgumentBefore, listed("격 조사", rule.ListPrevLast, argumentCases), not).Attacks(claim)
	except(g, rule.AroundIs, whyNoModifiableNext, &rule.Bits{Of: "around", Mask: rule.AroundBeforeModifiable}, not).Attacks(claim)
	except(g, rule.KindListed, whyAdnNext, listed("꾸밈 맺음", rule.ListNextLast, []string{adn})).Attacks(claim)
	is := func(n rule.Near) bool {
		return labelOf(n.Last) == adn && predicateStem(n) && !slices.Contains(formDeterminers, n.Text)
	}
	return sentenceRule{name: ruleArgumentAdn, graph: g, bonus: bonusOverCount, about: is, waits: is}
}
