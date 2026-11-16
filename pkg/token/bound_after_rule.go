//ff:func feature=token type=rule control=sequence
//ff:what boundAfterRule — the rule of a sentence "a dependent noun stands right after a modifying form": a candidate that reads its stem as a dependent noun (의존명사:수, 의존명사:거) is preferred, by bonusOverWhole, when the word right before it ends in a modifying closing (할 수도 → 수+도, 하는 건 → 거+는), and not otherwise, nor when that dependent noun stands after a nominal as well (nominal: 때, 간, 외, 시 …), nor when it is an adverbial dependent noun (adverbialBound: 대로, 만큼, 듯, 채) read with the particle 이/가 of the subject or the complement (미안한 듯이: the dependent noun 듯이, not 듯+이)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The other side of boundBeforeRule: after a modifying form the letters
// that a dependent noun begins are read as that noun with its particle
// first, before a longer noun (수도) or a verb (걸다+ㄴ) of the same
// letters.
func boundAfterRule(nominal []string) sentenceRule {
	g, claim := sentenceClaim(ruleBoundAfter, rule.KindListed, listed("의존 명사", rule.ListSelfKind, []string{boundSet}))
	except(g, rule.KindListed, whyNoAdn, listed("꾸밈 맺음", rule.ListPrevLast, []string{closingLabel(rule.KindAdn)}), not).Attacks(claim)
	except(g, rule.Listed, whyNominalSelf, listed("체언 뒤 의존 명사", rule.ListSelfStem, nominal)).Attacks(claim)
	adverbial := except(g, rule.Listed, whyAdverbialBound, listed("부사성 의존 명사", rule.ListSelfStem, adverbialBound))
	adverbial.Attacks(claim)
	except(g, rule.SelfHas, whyNoSubjectMark, listed("주격·보격 이·가", rule.ListSelfMarks, subjectNames), not).Attacks(adverbial)
	return sentenceRule{name: ruleBoundAfter, graph: g, bonus: bonusOverWhole, about: func(n rule.Near) bool { return labelOf(n.Kind) == boundSet }}
}
