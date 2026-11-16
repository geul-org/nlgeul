//ff:func feature=token type=rule control=sequence
//ff:what nominalBoundRule — the rule of a sentence "a dependent noun the word table names a noun stands after a nominal with its particle": a candidate that reads its stem as a noun of the word table whose letters are one of words (the dependent nouns of the dependent-noun table outside the closed set that stand after a nominal: 측, 중, 내, 분 …), with a particle as its first marker (particledNoun), is preferred, by bonusOverCount, when the word right before it ends in a nominal (회사 측은 as 측+은, not the noun 측은), and not otherwise
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// Such a dependent noun stands after a nominal (the standard dictionary's
// 측 側: 회사 측, 노조 측) and takes its particle there. The word table
// counts the longer noun of the same letters (측은 惻隱) far above it, as
// it counted the dependent noun's letters with the particle as one word;
// after a nominal this rule carries the dependent noun and its particle
// over that count. The copula's reading is left out: 내지 after a nominal
// is the adverb, not 내+(이)+지.
func nominalBoundRule(words []string) sentenceRule {
	g, claim := sentenceClaim(ruleNominalBound, rule.Listed, listed("체언 뒤 의존 명사(낱말 표 이름)", rule.ListSelfStem, words))
	except(g, rule.AroundIs, whyNotAfterNominal, &rule.Bits{Of: "around", Mask: rule.AroundAfterNominal}, not).Attacks(claim)
	return sentenceRule{name: ruleNominalBound, graph: g, bonus: bonusOverCount, about: func(n rule.Near) bool { return particledNoun(n, words) }}
}
