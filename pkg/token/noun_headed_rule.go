//ff:func feature=token type=rule control=sequence
//ff:what nounHeadedRule — the rule of a sentence "the letters of a dependent noun at the head of a sentence are a noun": a candidate that reads its stem, one of the dependent nouns given, as a noun of the word table (열쇠:체언) is preferred when no word stands before it in the sentence or only an opening sign does (시는 …, “시 …), and not when another word or sign stands right before it, nor right before 아니다 (뿐만 아니라); it waits for the word after it
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// A dependent noun is used only with a modifier or a nominal before it;
// where nothing of the sentence stands before it, or only the sign that
// opens a quotation or a bracket, the same letters are the noun (시 of
// 詩 or 市, 때, 만 of 萬).
func nounHeadedRule(words []string) sentenceRule {
	noun := keyKind + ":" + kindLabel[kindNominal]
	g, claim := sentenceClaim(ruleNounHeaded, rule.Listed, listed("의존 명사", rule.ListSelfStem, words))
	except(g, rule.Headed, whyNotHeaded, listed("여는 기호", rule.ListPrevKind, []string{openName}), not).Attacks(claim)
	except(g, rule.Listed, whyDenyNext, listed("아니다", rule.ListNextStem, denyStems)).Attacks(claim)
	is := func(n rule.Near) bool { return n.Kind == noun && slices.Contains(words, n.Stem) }
	return sentenceRule{name: ruleNounHeaded, graph: g, bonus: bonusSentence, about: is, waits: is}
}
