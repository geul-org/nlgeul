//ff:func feature=token type=rule control=sequence
//ff:what negAdverbRule — the rule of a sentence "안 and 못 right before a verb or an adjective are the negative adverb": a candidate that reads 안 or 못 (negAdverbs) alone as an adverb of the word table is preferred when the word right after it is a verb or an adjective of the tables (predStemNames), and not otherwise; the candidate waits for the word after it
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// The dictionary gives the adverbs 안 (아니) and 못 as standing before a
// predicate to deny it or say it cannot be done; the noun 안 (inside)
// takes a particle there or stands after the noun it is the inside of.
// The word table counts 안 the noun and 안 the adverb under one key, so
// that the two are read alike but for the order of the kinds, which puts
// the noun first.
func negAdverbRule() sentenceRule {
	g, claim := sentenceClaim(ruleNegAdverb, rule.Listed, listed("부정 부사", rule.ListSelfStem, negAdverbs))
	except(g, rule.Listed, whyNoPredNext, listed("동사·형용사", rule.ListNextName, predStemNames), not).Attacks(claim)
	adverb := keyKind + ":" + kindLabel[kindAdverb]
	is := func(n rule.Near) bool {
		return n.Kind == adverb && n.Marks == "" && slices.Contains(negAdverbs, n.Stem)
	}
	return sentenceRule{name: ruleNegAdverb, graph: g, bonus: bonusSentence, about: is, waits: is}
}
