//ff:func feature=parse type=rule control=sequence
//ff:what countedShortRule — the rule of a sentence "a noun of one syllable with its particle after a numeral determiner is that noun and the particle": a candidate that is a noun of the word table of one syllable with a particle or the copula written after it (shortNominal), not one whose letters are a dependent noun's (dependentWords: left to the rules of dependent nouns), is preferred, by bonusOverCount, when the word right before it is a numeral determiner of the closed set (dets: 한, 두) read as a determiner (한 대가 as 대+가, not the noun 대가); it falls, with the reason in a trace, when the word right before it is not one of dets or is not read as a determiner
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// A numeral determiner tells how many of the nominal after it there are
// (school grammar: 수 관형사), and that nominal is mostly a noun of a
// unit of one syllable (한 대, 두 편, 한 줌). The word table counts the
// noun of two syllables whose letters are such a unit with its particle
// (대가 代價, 편의 便宜) above the unit and the particle; after a numeral
// determiner this rule carries the unit and its particle over that count.
func countedShortRule(dets []string) sentenceRule {
	g, claim := sentenceClaim(ruleCountedShort, rule.Listed, listed("체언", rule.ListSelfKind, []string{tableNoun}))
	except(g, rule.Listed, whyNoNumDet, listed("수 관형사", rule.ListPrevStem, dets), not).Attacks(claim)
	except(g, rule.Listed, whyNotCounted, listed("관형사", rule.ListPrevKind, []string{determinerLabel}), not).Attacks(claim)
	return sentenceRule{name: ruleCountedShort, graph: g, bonus: bonusOverCount, about: func(n rule.Near) bool {
		return n.Kind == tableNoun && shortNominal(n) && !slices.Contains(dependentWords, n.Stem)
	}}
}
