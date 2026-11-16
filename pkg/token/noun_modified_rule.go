//ff:func feature=token type=rule control=sequence
//ff:what nounModifiedRule — the rule of a sentence "a noun stands right after a modifier": a candidate that reads its stem as a noun of the word table (열쇠:체언), not one whose letters are another closed-set word's (skip: the dependent nouns, the determiners), is preferred, by bonusOverCount, when the word right before it ends in a modifying closing or in the adnominal case particle, or is a determiner of the closed set (따른 보고, 새의 떼는, 한 줌), and not otherwise
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// A modifier (관형어: a predicate's modifying form, a nominal with the
// adnominal case particle 의, a determiner) is followed by the nominal it
// modifies (school grammar: 관형어는 체언을 꾸민다). The letters of a
// dependent noun are left to the rules of dependent nouns, and those of
// a determiner (이, 그) are not told apart by what stands before them.
func nounModifiedRule(skip []string) sentenceRule {
	noun := keyKind + ":" + kindLabel[kindNominal]
	g, claim := sentenceClaim(ruleNounModified, rule.Listed, listed("체언", rule.ListSelfKind, []string{noun}))
	none := except(g, rule.KindListed, whyNoModifier, listed("관형어", rule.ListPrevLast, []string{closingLabel(rule.KindAdn), adnCase}), not)
	none.Attacks(claim)
	except(g, rule.Listed, whyDet, listed("관형사", rule.ListPrevKind, []string{determinerLabel})).Attacks(none)
	return sentenceRule{name: ruleNounModified, graph: g, bonus: bonusOverCount, about: func(n rule.Near) bool {
		return n.Kind == noun && !slices.Contains(skip, n.Stem)
	}}
}
