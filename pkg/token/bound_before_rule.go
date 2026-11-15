//ff:func feature=token type=rule control=sequence
//ff:what boundBeforeRule — the rule of a sentence "a modifying form stands right before a dependent noun": a candidate whose last piece is a modifying closing (꾸밈: ㄹ/을, ㄴ/은, 는 …) is preferred, by bonusBound, when the stem of the word right after it is one of the dependent nouns given (알 수, 줄 것, 안은 채), and not otherwise; it waits for the word after it
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// A dependent noun is used only with a modifier before it (school
// grammar: 의존 명사는 관형어의 꾸밈을 받아야 쓰인다), and a word that
// ends in a particle or stands bare as a noun is no such modifier there:
// 알 before 수 is 알다 with the modifying ㄹ, not the noun 알.
func boundBeforeRule(bound []string) sentenceRule {
	adn := closingLabel(rule.KindAdn)
	g, claim := sentenceClaim(ruleBoundBefore, rule.KindListed, listed("꾸밈 맺음", rule.ListSelfLast, []string{adn}))
	except(g, rule.Listed, whyNoBoundNext, listed("의존 명사", rule.ListNextStem, bound), not).Attacks(claim)
	is := func(n rule.Near) bool { return nameKind(n.Last) == adn }
	return sentenceRule{name: ruleBoundBefore, graph: g, bonus: bonusBound, about: is, waits: is}
}
