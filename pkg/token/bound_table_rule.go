//ff:func feature=token type=rule control=sequence
//ff:what boundTableRule — the rule of a sentence "a dependent noun the word table names a noun stands right after a modifying form": a candidate that reads its stem as a noun of the word table (열쇠:체언) whose letters are a dependent noun of the dependent-noun table that stands after a modifying closing and is not of the closed set (words: 리, 바, 지, 터 …) is preferred, by bonusOverWhole, when the word right before it ends in a modifying closing (그럴 리가 → 리+가), and not otherwise
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// A dependent noun is used only with a modifier before it (school
// grammar: 의존 명사는 관형어의 꾸밈을 받아야 쓰인다). The dependent nouns
// outside the closed set bear the word table's name of a noun, so
// boundAfterRule does not ask them and nounModifiedRule leaves their
// letters alone; after a modifying form the longer noun of the same
// letters (리가 over 리+가) then took the noun's place. This rule carries
// the dependent noun read apart with its particle over that whole word,
// as boundAfterRule does for the closed set.
func boundTableRule(words []string) sentenceRule {
	g, claim := sentenceClaim(ruleBoundTable, rule.Listed, listed("의존 명사(낱말 표 이름)", rule.ListSelfStem, words))
	except(g, rule.KindListed, whyNoAdn, listed("꾸밈 맺음", rule.ListPrevLast, []string{closingLabel(rule.KindAdn)}), not).Attacks(claim)
	return sentenceRule{name: ruleBoundTable, graph: g, bonus: bonusOverWhole, about: func(n rule.Near) bool {
		return n.Kind == tableNoun && n.Marks != "" && slices.Contains(words, n.Stem)
	}}
}
