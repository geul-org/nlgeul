//ff:func feature=token type=rule control=sequence
//ff:what gluedBoundRule — the rule of a sentence "a dependent noun the word table names a noun, written onto a nominal with its particle, is that noun and the particle": a candidate that reads its stem as a noun of the word table (열쇠:체언) whose letters are one of words (the dependent nouns of the dependent-noun table outside the closed set that stand after a modifying closing: 자, 성, 법 …), with a particle as its first marker (particleKinds), is preferred, by bonusOverCount, when no space stands before its word and the word right before it ends in a nominal (패소자가 as 패소 자+가, not 패소 자가)
package token

import (
	"slices"
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// Written onto a nominal, the letters of such a dependent noun are the
// suffix of the same letters (the standard dictionary's -자 者, -성 性,
// -법 法), which ends the noun it makes, and a particle follows it.
// gluedShortRule gives back what the rebuttal "one syllable long" took
// from the noun of one syllable there; the longer noun of the same
// letters (자가) still stood above it by its count in the word table, and
// this rule carries the dependent noun and its particle over that count.
func gluedBoundRule(words []string) sentenceRule {
	g := toulmin.NewGraph(ruleGluedBound)
	claim := g.Rule(rule.AroundIs).With(&rule.Bits{Of: "around", Mask: rule.AroundGlued}).With(&rule.Claim{Name: ruleGluedBound, Concludes: preferLabel})
	except(g, rule.AroundIs, whyNotAfterNominal, &rule.Bits{Of: "around", Mask: rule.AroundAfterNominal}, not).Attacks(claim)
	return sentenceRule{name: ruleGluedBound, graph: g, bonus: bonusOverCount, about: func(n rule.Near) bool {
		marks := rule.MarksOf(n)
		if n.Kind != tableNoun || len(marks) == 0 || !slices.Contains(words, n.Stem) {
			return false
		}
		first, _, _ := strings.Cut(marks[0], ":")
		return slices.Contains(particleKinds, first)
	}}
}
