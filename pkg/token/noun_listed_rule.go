//ff:func feature=token type=rule control=sequence
//ff:what nounListedRule — the rule of a sentence "the letters of a dependent noun right before a middle dot are a noun": a candidate that reads its stem, one of the dependent nouns given, as a noun of the word table (열쇠:체언) is preferred when the word right after it is a middle dot (가운뎃점: 시ㆍ군ㆍ구, 시ㆍ도지사), and not otherwise; it waits for the word after it
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The rules of punctuation give the middle dot to join words listed
// side by side (열거할 어구들을 일정한 기준으로 묶어서 나타낼 때); a
// dependent noun, which needs a modifier before it, does not stand as one
// of them alone, and the same letters are the noun (시 of 市).
func nounListedRule(words []string) sentenceRule {
	noun := keyKind + ":" + kindLabel[kindNominal]
	g, claim := sentenceClaim(ruleNounListed, rule.Listed, listed("의존 명사", rule.ListSelfStem, words))
	except(g, rule.Listed, whyNoDotNext, listed("가운뎃점", rule.ListNextLast, []string{middleDot}), not).Attacks(claim)
	is := func(n rule.Near) bool { return n.Kind == noun && slices.Contains(words, n.Stem) }
	return sentenceRule{name: ruleNounListed, graph: g, bonus: bonusSentence, about: is, waits: is}
}
