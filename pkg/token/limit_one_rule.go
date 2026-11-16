//ff:func feature=token type=rule control=sequence
//ff:what limitOneRule — the rule of a sentence "한 after the modifying 는 is the noun 한(限)": the candidate that reads 한 as a nominal of the word table (열쇠:체언) is preferred when the word before it ends in the modifying closing 는 (neunAdn), and not otherwise
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The dictionary gives the noun 한(限) "(mostly in -는 한) a condition"
// (살아 있는 한, 허락하지 않는 한). The word table's counts put the
// determiner 한 first there.
func limitOneRule() sentenceRule {
	g, claim := sentenceClaim(ruleLimitOne, rule.Listed, listed("한", rule.ListSelfStem, []string{oneStem}))
	except(g, rule.Listed, whyNoNeun, listed("꾸밈 는", rule.ListPrevLast, []string{neunAdn}), not).Attacks(claim)
	return sentenceRule{name: ruleLimitOne, graph: g, bonus: bonusOverCount, about: func(n rule.Near) bool { return n.Stem == oneStem && n.Kind == oneNoun }}
}
