//ff:func feature=parse type=rule control=sequence
//ff:what nounHeadedRule — the rule of a sentence "the letters of a dependent noun at the head of a sentence or right before a middle dot are a noun": a candidate that reads its stem, one of the dependent nouns given, as a noun of the word table (열쇠:체언) is preferred when no word stands before it in the sentence or only an opening sign does (시는 …, “시 …), or when the word right after it is a middle dot (시ㆍ군ㆍ구, 시ㆍ도지사), and not when another word or sign stands right before it with no middle dot after it, nor right before 아니다 (뿐만 아니라); it waits for the word after it
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// A dependent noun is used only with a modifier or a nominal before it;
// where nothing of the sentence stands before it, or only the sign that
// opens a quotation or a bracket, the same letters are the noun (시 of
// 詩 or 市, 때, 만 of 萬). The rules of punctuation give the middle dot
// to join words listed side by side (열거할 어구들을 일정한 기준으로
// 묶어서 나타낼 때); a dependent noun does not stand as one of them
// alone, so right before a middle dot the letters are the noun too,
// whatever stands before them (시 of 市).
func nounHeadedRule(words []string) sentenceRule {
	noun := keyKind + ":" + kindLabel[kindNominal]
	g, claim := sentenceClaim(ruleNounHeaded, rule.Listed, listed("의존 명사", rule.ListSelfStem, words))
	notHeaded := except(g, rule.Headed, whyNotHeaded, listed("여는 기호", rule.ListPrevKind, []string{openName}), not)
	notHeaded.Attacks(claim)
	except(g, rule.Listed, whyDotNext, listed("가운뎃점", rule.ListNextLast, []string{middleDot})).Attacks(notHeaded)
	except(g, rule.Listed, whyDenyNext, listed("아니다", rule.ListNextStem, denyStems)).Attacks(claim)
	is := func(n rule.Near) bool { return n.Kind == noun && slices.Contains(words, n.Stem) }
	return sentenceRule{name: ruleNounHeaded, graph: g, bonus: bonusSameLetters, about: is, waits: is}
}
