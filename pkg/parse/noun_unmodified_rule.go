//ff:func feature=parse type=rule control=sequence
//ff:what nounUnmodifiedRule — the rule of a sentence "the letters of a dependent noun that stands only after a modifying form are the noun where none stands before them": a candidate that reads its stem, one of the words given (수, 줄), as a noun of the word table (열쇠:체언) is preferred when the word right before it does not end in a modifying closing (인구 수는, 빨래 줄을), and not otherwise (알 수), nor right before 알다 or 모르다 (… 줄 알고); it waits for the word after it
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The standard dictionary gives the dependent nouns 수 and 줄 only after
// a predicate's modifying form (-을 수 있다, -ㄴ 줄 알다); with no such
// form before them the same letters are the nouns 수(數) and 줄(a line).
func nounUnmodifiedRule(words []string) sentenceRule {
	noun := keyKind + ":" + kindLabel[kindNominal]
	g, claim := sentenceClaim(ruleNounUnmodified, rule.Listed, listed("꾸밈 뒤에만 서는 의존 명사", rule.ListSelfStem, words))
	except(g, rule.KindListed, whyAdn, listed("꾸밈 맺음", rule.ListPrevLast, []string{closingLabel(rule.KindAdn)})).Attacks(claim)
	except(g, rule.Listed, whyKnowNext, listed("알다·모르다", rule.ListNextStem, knowStems)).Attacks(claim)
	is := func(n rule.Near) bool { return n.Kind == noun && slices.Contains(words, n.Stem) }
	return sentenceRule{name: ruleNounUnmodified, graph: g, bonus: bonusSameLetters, about: is, waits: is}
}
